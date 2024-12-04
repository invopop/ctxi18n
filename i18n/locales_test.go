package i18n_test

import (
	"encoding/json"
	"testing"

	"github.com/invopop/ctxi18n/i18n"
	"github.com/invopop/ctxi18n/internal/examples"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLocalesLoad(t *testing.T) {
	ls := new(i18n.Locales)

	err := ls.Load(examples.Content)
	require.NoError(t, err)

	en := ls.Get("en")
	require.NotNil(t, en)
	assert.Equal(t, "Log In", en.T("login.button"))
	assert.Equal(t, "Extensions", en.T("ext.test"))

	es := ls.Get("es")
	require.NotNil(t, es)
	assert.Equal(t, "Iniciar Sesión", es.T("login.button"))

	l := ls.Match("en-US,en;q=0.9,es;q=0.8")
	require.NotNil(t, l)
	assert.Equal(t, "en", l.Code().String())
	assert.Equal(t, "Log In", l.T("login.button"))
	assert.Equal(t, "Special Label", l.T("special_label"))

	l = ls.Match("es-ES,es;q=0.9,en;q=0.8")
	require.NotNil(t, l)
	assert.Equal(t, "es", l.Code().String())
	assert.Equal(t, "Iniciar Sesión", l.T("login.button"))
	assert.Equal(t, "!(MISSING: special_label)", l.T("special_label"))

	assert.Nil(t, ls.Match("inv"))
}

func TestLoadWithDefault(t *testing.T) {
	ls := new(i18n.Locales)
	err := ls.LoadWithDefault(examples.Content, "en")
	require.NoError(t, err)

	l := ls.Match("en")
	require.NotNil(t, l)
	assert.Equal(t, "en", l.Code().String())
	assert.Equal(t, "Log In", l.T("login.button"))
	assert.Equal(t, "Special Label", l.T("special_label"))

	l = ls.Match("es")
	require.NotNil(t, l)
	assert.Equal(t, "es", l.Code().String())
	assert.Equal(t, "Iniciar Sesión", l.T("login.button"))
	assert.Equal(t, "Special Label", l.T("special_label"))

	ls = new(i18n.Locales)
	err = ls.LoadWithDefault(examples.Content, "bad")
	assert.ErrorContains(t, err, "undefined default locale: bad")

}

func TestLocalesUnmarshalJSON(t *testing.T) {
	in := SampleLocales()
	ls := new(i18n.Locales)
	err := json.Unmarshal(in, ls)
	require.NoError(t, err)
	l := ls.Get("en")
	assert.Equal(t, "en", l.Code().String())
	assert.Equal(t, "quux", l.T("baz.qux"))

	// Now try merging with another set of entries
	sub := []byte(`{
		"en": {
			"a": "b",
			"baz": {
				"zux": "zuux"
			}
		}
	}`)
	err = json.Unmarshal(sub, ls)
	require.NoError(t, err)
	assert.Equal(t, "en", l.Code().String())
	assert.Equal(t, "quux", l.T("baz.qux"))
	assert.Equal(t, "b", l.T("a"))
	assert.Equal(t, "zuux", l.T("baz.zux"))

	t.Run("empty", func(t *testing.T) {
		ls := new(i18n.Locales)
		err := ls.UnmarshalJSON([]byte{})
		require.NoError(t, err)
	})
	t.Run("invalid", func(t *testing.T) {
		ls := new(i18n.Locales)
		err := ls.UnmarshalJSON([]byte("'bad'"))
		require.ErrorContains(t, err, "invalid character")
	})
}

func TestLocalesCodes(t *testing.T) {
	in := SampleLocales()
	ls := new(i18n.Locales)
	require.NoError(t, json.Unmarshal(in, ls))
	codes := ls.Codes()
	assert.Len(t, codes, 2)
	assert.Contains(t, codes, i18n.Code("en"))
	assert.Contains(t, codes, i18n.Code("es"))
}

func SampleLocales() []byte {
	return []byte(`{
		"en": {
			"foo": "bar",
			"baz": {
				"qux": "quux",
				"plural": {
					"zero": "no mice",
					"one": "%{count} mouse",
					"other": "%{count} mice"
				}
			}
		},
		"es": {
			"foo": "bara",
			"baz": {
				"qux": "quuxa",
				"plural": {
					"zero": "no ratones",
					"one": "%{count} ratón",
					"other": "%{count} ratones"
				}
			}
		}
	}`)
}
