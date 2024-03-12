package i18n

import (
	"encoding/json"
	"strings"
)

const (
	// MissingDictKey is provided when a specific entry in the dictionary
	// cannot be found.
	MissingDictKey = "!(MISSING)"
)

// Dict holds the internationalization entries for a specific locale.
type Dict struct {
	value   string
	entries map[string]*Dict
}

// Get provides the value from the dictionary at the provided key location.
func (d *Dict) get(key string) string {
	entry := d.getEntry(key)
	if entry == nil {
		return MissingDictKey
	}
	return entry.value
}

// GetEntry recursively retrieves the dictionary at the provided key location.
func (d *Dict) getEntry(key string) *Dict {
	if key == "" {
		return nil
	}
	n := strings.SplitN(key, ".", 2)
	entry, ok := d.entries[n[0]]
	if !ok {
		return nil
	}
	if len(n) == 1 {
		return entry
	}
	return entry.getEntry(n[1])
}

// merge combines the entries of the second dictionary into this one.
func (d *Dict) merge(d2 *Dict) {
	if d2 == nil {
		return
	}
	if d.entries == nil {
		d.entries = make(map[string]*Dict)
	}
	for k, v := range d2.entries {
		if d.entries[k] == nil {
			d.entries[k] = v
			continue
		}
		d.entries[k].merge(v)
	}
}

// UnmarshalJSON attempts to load the dictionary data from a JSON byte slice.
func (d *Dict) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	if data[0] == '"' {
		d.value = string(data[1 : len(data)-1])
		return nil
	}
	d.entries = make(map[string]*Dict)
	return json.Unmarshal(data, &d.entries)
}
