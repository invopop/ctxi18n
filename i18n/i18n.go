// Package i18n is responsible for keeping the key internationalization in one
// place.
package i18n

import (
	"context"
	"fmt"
	"strings"
)

// M stands for map and is a simple helper to make it easier to work with
// internationalization maps.
type M map[string]any

// T is responsible for translating a key into a string by extracting
// the local from the context.
func T(ctx context.Context, key string, args ...any) string {
	l := GetLocale(ctx)
	if l == nil {
		return MissingDictKey
	}
	return l.Get(key, args...)
}

// N returns the pluralized translation of the provided key using n
// as the count.
func N(ctx context.Context, key string, n int, args ...any) string {
	l := GetLocale(ctx)
	if l == nil {
		return MissingDictKey
	}
	return l.GetN(key, n, args...)
}

// Replace is used to interpolate the matched keys in the provided
// string with their values in the map.
//
// Interpolation is performed using the `%{key}` pattern.
func (m M) Replace(in string) string {
	for k, v := range m {
		in = strings.Replace(in, fmt.Sprintf("%%{%s}", k), fmt.Sprint(v), -1)
	}
	return in
}
