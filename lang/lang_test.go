package lang_test

import (
	"testing"

	"fyne.io/fyne/v2/lang"
	"github.com/stretchr/testify/assert"
)

func TestLocalize_Fallback(t *testing.T) {
	assert.Equal(t, "Missing", lang.L("Missing"))
}

func TestLocalize_Struct(t *testing.T) {
	type data struct {
		Str string
	}

	assert.Equal(t, "Hello World!", lang.L("Hello {{.Str}}!", data{Str: "World"}))
}

func TestLocalize_Map(t *testing.T) {
	assert.Equal(t, "Hello World!", lang.L("Hello {{.Str}}!", map[string]string{"Str": "World"}))
}

func TestLocalizePlural_Fallback(t *testing.T) {
	assert.Equal(t, "Missing", lang.N("Missing", 1))
	assert.Equal(t, "Apple", lang.N("Apple", 1))
}
