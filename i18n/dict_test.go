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
	assert.Equal(t, "bar", dict.get("foo"))
	assert.Equal(t, "quux", dict.get("baz.qux"))
	assert.Equal(t, "", dict.get("baz.plural"))
	assert.Equal(t, "!(MISSING)", dict.get("random"))
}
