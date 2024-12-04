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
	assert.Equal(t, "bar", dict.Get("foo").Value())
	assert.Equal(t, "quux", dict.Get("baz.qux").Value())
	assert.Empty(t, dict.Get("baz.plural").Value())
	assert.Empty(t, dict.Get("random").Value())

	t.Run("empty", func(t *testing.T) {
		d := new(Dict)
		err := d.UnmarshalJSON([]byte{})
		require.NoError(t, err)
	})
}

func TestDictAdd(t *testing.T) {
	d := NewDict()
	assert.Nil(t, d.Get(""))
	d.Add("foo", "bar")
	assert.Equal(t, "bar", d.Get("foo").Value())

	d.Add("plural", map[string]any{
		"zero":  "no mice",
		"one":   "%s mouse",
		"other": "%s mice",
	})
	assert.Equal(t, "no mice", d.Get("plural.zero").Value())
	assert.Equal(t, "%s mice", d.Get("plural.other").Value())

	d.Add("bad", 10) // ignore
	assert.Nil(t, d.Get("bad"))

	d.Add("self", d)
	assert.Equal(t, "bar", d.Get("self.foo").Value())
}

func TestDictHas(t *testing.T) {
	t.Run("simple case", func(t *testing.T) {
		d := NewDict()
		assert.False(t, d.Has("foo"))
		d.Add("foo", "bar")
		assert.True(t, d.Has("foo"))
		assert.False(t, d.Has("baz"))
	})
	t.Run("nested", func(t *testing.T) {
		ex := `{
			"foo": "bar",
			"baz": {
				"qux": "quux"
			}
		}`
		d := new(Dict)
		require.NoError(t, json.Unmarshal([]byte(ex), d))
		assert.True(t, d.Has("foo"))
		assert.True(t, d.Has("baz.qux"))
		assert.False(t, d.Has("baz.quux"))
	})
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
		"foo": "baz",
		"extra": "value"
	}`
	d2 := new(Dict)
	require.NoError(t, json.Unmarshal([]byte(ex2), d2))

	d1.Merge(nil) // does nothing

	d3 := new(Dict)
	d3.Merge(d2)
	assert.Equal(t, "value", d3.Get("extra").Value())

	d1.Merge(d2)
	assert.Equal(t, "bar", d1.Get("foo").Value(), "should not overwrite")
	assert.Equal(t, "value", d1.Get("extra").Value())
}
