package i18n

import (
	"context"
	"encoding/json"
	"fmt"
)

// Locale holds the internationalization entries for a specific locale.
type Locale struct {
	code Code
	dict *Dict
	rule PluralRule
}

const (
	localeKey Code = "locale"
)

// PluralRule defines a simple method that expects a dictionary and number and
// will find a matching string value.
type PluralRule func(d *Dict, num int) string

// NewLocale creates a new locale with the provided key and dictionary.
func NewLocale(code Code, dict *Dict) *Locale {
	l := &Locale{
		code: code,
		dict: dict,
	}
	l.rule = mapPluralRule(code)
	return l

}

// Code returns the language code of the locale.
func (l *Locale) Code() Code {
	return l.code
}

// Get provides the value from the dictionary stored by the locale.
func (l *Locale) Get(key string, args ...any) string {
	return interpolate(l.dict.get(key), args...)
}

// GetN uses the locale pluralization rules to determine which
// string value to provide based on the provided number.
func (l *Locale) GetN(key string, n int, args ...any) string {
	entry := l.dict.getEntry(key)
	if entry == nil {
		return MissingDictKey
	}
	return interpolate(l.rule(entry, n), args...)
}

// PluralRule provides the pluralization rule for the locale.
func (l *Locale) PluralRule() PluralRule {
	return l.rule
}

// UnmarshalJSON attempts to load the locale from a JSON byte slice.
func (l *Locale) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	l.dict = new(Dict)
	if err := json.Unmarshal(data, l.dict); err != nil {
		return err
	}
	return nil
}

// WithContext inserts the locale into the context so that it can be
// loaded later with `GetLocale`.
func (l *Locale) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, localeKey, l)
}

// GetLocale retrieves the locale from the context.
func GetLocale(ctx context.Context) *Locale {
	if l, ok := ctx.Value(localeKey).(*Locale); ok {
		return l
	}
	return nil
}

func interpolate(s string, args ...any) string {
	if len(args) > 0 {
		switch arg := args[0].(type) {
		case M:
			return arg.Replace(s)
		default:
			return fmt.Sprintf(s, args...)
		}
	}
	return s
}
