// Package ctxi18n is used to internationalize applications using the context
// for the locale.
package ctxi18n

import (
	"context"
	"errors"
	"io/fs"

	"github.com/invopop/ctxi18n/i18n"
)

var (
	// DefaultLocale defines the default or fallback locale code to use
	// if no other match inside the packages list was found.
	DefaultLocale i18n.Code = "en"
)

var (
	locales *i18n.Locales
)

var (
	// ErrMissingLocale implies that the requested locale was not found
	// in the current index.
	ErrMissingLocale = errors.New("locale not defined")
)

// Load walks through all the files in provided File System and prepares
// an internal global list of locales ready to use.
func Load(fs fs.FS) error {
	locales = new(i18n.Locales)
	return locales.Load(fs)
}

// WithLocale tries to match the provided code with a locale and ensures
// it is available inside the context.
func WithLocale(ctx context.Context, locale string) (context.Context, error) {
	l := locales.Match(locale)
	if l == nil {
		l = locales.Get(DefaultLocale)
		if l == nil {
			return nil, ErrMissingLocale
		}
	}
	return l.WithContext(ctx), nil
}

// Locale provides the locale object currently stored in the context.
func Locale(ctx context.Context) *i18n.Locale {
	return i18n.GetLocale(ctx)
}
