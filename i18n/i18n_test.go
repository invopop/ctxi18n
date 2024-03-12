package i18n_test

import (
	"testing"

	"github.com/invopop/ctxi18n/i18n"
	"github.com/stretchr/testify/assert"
)

func TestMReplace(t *testing.T) {
	m := i18n.M{
		"string": "value",
		"num":    42,
	}
	out := m.Replace("This is a %{string} and a %{num}.")
	assert.Equal(t, "This is a value and a 42.", out)
}
