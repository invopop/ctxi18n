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
	assert.Equal(t, "Log In", en.Get("login.button"))
	assert.Equal(t, "Extensions", en.Get("ext.test"))

	es := ls.Get("es")
	require.NotNil(t, es)
	assert.Equal(t, "Iniciar Sesión", es.Get("login.button"))

	l := ls.Match("en-US,en;q=0.9,es;q=0.8")
	require.NotNil(t, l)
	assert.Equal(t, "en", l.Code().String())

	assert.Nil(t, ls.Match("inv"))
}

func TestLocalesUnmarshalJSON(t *testing.T) {
	in := SampleLocales()
	ls := new(i18n.Locales)
	err := json.Unmarshal(in, ls)
	require.NoError(t, err)
	l := ls.Get("en")
	assert.Equal(t, "en", l.Code().String())
	assert.Equal(t, "quux", l.Get("baz.qux"))

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
	assert.Equal(t, "quux", l.Get("baz.qux"))
	assert.Equal(t, "b", l.Get("a"))
	assert.Equal(t, "zuux", l.Get("baz.zux"))
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
		}
	}`)
}
