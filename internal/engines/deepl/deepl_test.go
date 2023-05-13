package deepl

import (
	"testing"

	"github.com/goura/yaku-cli/pkg/config"
	"github.com/stretchr/testify/assert"
	"github.com/thoas/go-funk"
	"golang.org/x/text/language"
)

func TestLoadConfig(t *testing.T) {
	engine := DeepLEngine{}

	// Empty config should fail
	empty_conf := config.Config{}
	err := engine.LoadConfig(empty_conf)
	assert.Error(t, err)

	// OK case
	conf := config.Config{DeeplAuthKey: "a"}
	err = engine.LoadConfig(conf)
	assert.NoError(t, err)
}

func TestIsSizeOK(t *testing.T) {
	engine := DeepLEngine{}

	var isOkay bool

	// Small string should pass
	isOkay = engine.IsSourceSizeOK("Hello!")
	assert.True(t, isOkay)

	// Big string should fail
	s := ""
	for i := 0; i < (16 * 128); i++ {
		s += "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" // 64 letters
	}

	isOkay = engine.IsSourceSizeOK(s)
	assert.False(t, isOkay)
}

func TestSupportedSourceLanguages(t *testing.T) {
	engine := DeepLEngine{}

	langs, err := engine.SupportedSourceLanguages()
	assert.NoError(t, err)

	// Test some languages with fallback
	assert.True(t, funk.Contains(langs, language.English))
	assert.True(t, funk.Contains(langs, language.Portuguese))
	assert.True(t, funk.Contains(langs, language.AmericanEnglish))
	assert.True(t, funk.Contains(langs, language.BritishEnglish))
}

func TestSupportedTargetLanguages(t *testing.T) {
	engine := DeepLEngine{}

	langs, err := engine.SupportedTargetLanguages(language.AmericanEnglish)
	assert.NoError(t, err)

	// Test some languages with fallback
	assert.True(t, funk.Contains(langs, language.English))
	assert.True(t, funk.Contains(langs, language.Portuguese))
	assert.True(t, funk.Contains(langs, language.EuropeanPortuguese))
}

func TestSetEndpoint(t *testing.T) {
	engine := DeepLEngine{}

	// SetEndpoint works
	url := "https://example.com/v2/translate"
	if err := engine.SetEndpoint(url); err != nil {
		t.Error(err)
	}
	assert.Equal(t, url, engine.ServerURL)
}
