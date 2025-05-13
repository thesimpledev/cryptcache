package main

import (
	"crypto/ed25519"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

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
