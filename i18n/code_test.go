package i18n_test

import (
	"testing"

	"github.com/invopop/ctxi18n/i18n"
	"github.com/stretchr/testify/assert"
)

func TestCode(t *testing.T) {
	c := i18n.Code("en-US")
	assert.Equal(t, "en-US", c.String())
	assert.Equal(t, i18n.Code("en"), c.Base())

	c = i18n.Code("en")
	assert.Equal(t, i18n.Code("en"), c.Base())

	c = i18n.Code("x")
	assert.Equal(t, i18n.Code("x"), c.Base())

	c = i18n.Code("")
	assert.Equal(t, i18n.Code(""), c.Base())
}

func TestParseAcceptLanguage(t *testing.T) {
	list := i18n.ParseAcceptLanguage("en")
	assert.Equal(t, []i18n.Code{"en"}, list)

	list = i18n.ParseAcceptLanguage("en-US, en;q=0.5")
	assert.Equal(t, []i18n.Code{"en-US", "en"}, list)

	list = i18n.ParseAcceptLanguage("en-US, en;q=0.5, es-419;q=0.8")
	assert.Equal(t, []i18n.Code{"en-US", "en", "es-419"}, list)
}
