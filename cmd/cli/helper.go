package main

import (
	"bytes"
	"crypto/ed25519"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

func cryptcacheExists() bool {
	cwd, err := os.Getwd()
	if err != nil {
		return false
	}

	path := filepath.Join(cwd, "cryptcache.toml")
	_, err = os.Stat(path)
	return err == nil
}

type kv struct {
	key string
	val any
}

func sortTOML(data map[string]any, excludeKeys ...string) ([]byte, error) {
	exclude := make(map[string]struct{})
	for _, k := range excludeKeys {
		exclude[k] = struct{}{}
	}

	ordered := make([]kv, 0, len(data))

	title, hasTitle := data["title"]
	profilesRaw, hasProfiles := data["profiles"]
	var profiles []string
	if hasProfiles {
		profiles, _ = profilesRaw.([]string)
	}

	if hasTitle {
		ordered = append(ordered, kv{"title", title})
	}
	if hasProfiles {
		ordered = append(ordered, kv{"profiles", profiles})
	}

	delete(data, "title")
	delete(data, "profiles")

	for _, profile := range profiles {
		prefix := profile + "."
		matching := map[string]any{}
		for k, v := range data {
			if strings.HasPrefix(k, prefix) {
				matching[k] = v
				delete(data, k)
			}
		}

		sortedKeys := make([]string, 0, len(matching))
		for k := range matching {
			sortedKeys = append(sortedKeys, k)
		}
		sort.Strings(sortedKeys)

		for _, k := range sortedKeys {
			if _, skip := exclude[k]; !skip {
				ordered = append(ordered, kv{k, matching[k]})
			}
		}
	}

	leftoverKeys := make([]string, 0, len(data))
	for k := range data {
		leftoverKeys = append(leftoverKeys, k)
	}
	sort.Strings(leftoverKeys)
	for _, k := range leftoverKeys {
		if _, skip := exclude[k]; !skip {
			ordered = append(ordered, kv{k, data[k]})
		}
	}

	var buf bytes.Buffer
	enc := toml.NewEncoder(&buf)
	for i, entry := range ordered {
		if i > 0 {
			buf.WriteString("\n")
		}
		if err := enc.Encode(map[string]any{entry.key: entry.val}); err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

func generateSignature(data []byte, privateKeyPath string) ([]byte, error) {

	privateKey, err := resolveKey(privateKeyPath)
	if err != nil {
		return nil, err
	}

	return ed25519.Sign(privateKey, data), nil
}

func verifySignature(data []byte, sig []byte, publicKeyPath string) (bool, error) {
	publicKey, err := resolveKey(publicKeyPath)
	if err != nil {
		return false, err
	}

	if ed25519.Verify(publicKey, data, sig) {
		return true, nil
	}
	return false, nil
}

func resolveKey(input string) ([]byte, error) {
	switch {
	case strings.HasPrefix(input, "http://") || strings.HasPrefix(input, "https://"):
		resp, err := http.Get(input)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch public key from URL: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to fetch public key: HTTP %d", resp.StatusCode)
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read public key response body: %w", err)
		}
		return data, nil

	case strings.HasPrefix(input, "file:"):
		path := strings.TrimPrefix(input, "file:")
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("failed to read public key from file %s: %w", path, err)
		}
		return data, nil

	case strings.HasPrefix(input, "inline:"):
		return []byte(strings.TrimPrefix(input, "inline:")), nil

	default:
		// fallback to file
		data, err := os.ReadFile(input)
		if err != nil {
			return nil, fmt.Errorf("failed to read public key from file %s: %w", input, err)
		}
		return data, nil
	}
}
