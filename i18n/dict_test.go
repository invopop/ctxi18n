package i18n

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDictUnmarshalJSON(t *testing.T) {
	ex := `{
		"foo": "bar",
		"baz": {
			"qux": "quux",
			"plural": {
				"zero": "no mice",
				"one": "%s mouse",
				"other": "%s mice"
			}
		}
	}`
	dict := new(Dict)
	err := json.Unmarshal([]byte(ex), dict)
	require.NoError(t, err)
	assert.Equal(t, "bar", dict.Get("foo"))
	assert.Equal(t, "quux", dict.Get("baz.qux"))
	assert.Equal(t, "", dict.Get("baz.plural"))
	assert.Equal(t, "!(MISSING)", dict.Get("random"))
}

func TestDictAdd(t *testing.T) {
	d := NewDict()
	d.Add("foo", "bar")
	assert.Equal(t, "bar", d.Get("foo"))

	d.Add("plural", map[string]any{
		"zero":  "no mice",
		"one":   "%s mouse",
		"other": "%s mice",
	})
	assert.Equal(t, "no mice", d.Get("plural.zero"))
	assert.Equal(t, "%s mice", d.Get("plural.other"))

	d.Add("bad", 10) // ignore
	assert.Equal(t, MissingDictKey, d.Get("bad"))

	d.Add("self", d)
	assert.Equal(t, "bar", d.Get("self.foo"))
}

func TestDictMerge(t *testing.T) {
	ex := `{
		"foo": "bar",
		"baz": {
			"qux": "quux",
			"plural": {
				"zero": "no mice",
				"one": "%s mouse",
				"other": "%s mice"
			}
		}
	}`
	d1 := new(Dict)
	require.NoError(t, json.Unmarshal([]byte(ex), d1))

	ex2 := `{
		"extra": "value"
	}`
	d2 := new(Dict)
	require.NoError(t, json.Unmarshal([]byte(ex2), d2))

	d1.Merge(nil) // does nothing

	d3 := new(Dict)
	d3.Merge(d2)
	assert.Equal(t, "value", d3.Get("extra"))

	d1.Merge(d2)
	assert.Equal(t, "bar", d1.Get("foo"))
	assert.Equal(t, "value", d1.Get("extra"))
}
