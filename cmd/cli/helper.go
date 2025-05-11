package main

import (
	"bytes"
	"sort"

	"github.com/pelletier/go-toml/v2"
)

func sortTOML(data map[string]any) ([]byte, error) {
	ordered := make([]kv, 0, len(data))

	// Extract title and profiles
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

	// Remove title and profiles from original map
	delete(data, "title")
	delete(data, "profiles")

	// Add profile sections in order from profiles list
	for _, profile := range profiles {
		prefix := profile + "."
		matching := map[string]any{}
		for k, v := range data {
			if len(k) > len(prefix) && k[:len(prefix)] == prefix {
				matching[k] = v
				delete(data, k)
			}
		}

		// Sort keys within the profile group
		sortedKeys := make([]string, 0, len(matching))
		for k := range matching {
			sortedKeys = append(sortedKeys, k)
		}
		sort.Strings(sortedKeys)

		for _, k := range sortedKeys {
			ordered = append(ordered, kv{k, matching[k]})
		}
	}

	// Append any remaining keys sorted
	leftoverKeys := make([]string, 0, len(data))
	for k := range data {
		leftoverKeys = append(leftoverKeys, k)
	}
	sort.Strings(leftoverKeys)
	for _, k := range leftoverKeys {
		ordered = append(ordered, kv{k, data[k]})
	}

	// Marshal the ordered map with spacing between top-level keys
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

type kv struct {
	key string
	val any
}
