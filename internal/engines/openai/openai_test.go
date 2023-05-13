package openai

import (
	"testing"

	"github.com/goura/yaku-cli/pkg/config"
	"github.com/goura/yaku-cli/pkg/types"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func TestLoadConfig(t *testing.T) {
	engine := OpenAIChatCompletionEngine{}

	// Empty config should fail
	empty_conf := config.Config{}
	err := engine.LoadConfig(empty_conf)
	assert.Error(t, err)

	// OK case
	conf := config.Config{OpenaiApiKey: "a"}
	err = engine.LoadConfig(conf)
	assert.NoError(t, err)
}

func TestIsSizeOK(t *testing.T) {
	engine := OpenAIChatCompletionEngine{}

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
	engine := OpenAIChatCompletionEngine{}

	_, err := engine.SupportedSourceLanguages()
	assert.ErrorIsf(t, err, types.FeatureNotSupportedError, "Returns FeatureNotSupportedError")
}

func TestSupportedTargetLanguages(t *testing.T) {
	engine := OpenAIChatCompletionEngine{}

	_, err := engine.SupportedTargetLanguages(language.AmericanEnglish)
	assert.ErrorIsf(t, err, types.FeatureNotSupportedError, "Returns FeatureNotSupportedError")
}

func TestSetEndpoint(t *testing.T) {
	engine := OpenAIChatCompletionEngine{}

	// SetEndpoint works
	url := "https://example.com/v2/translate"
	if err := engine.SetEndpoint(url); err != nil {
		t.Error(err)
	}
	assert.Equal(t, url, engine.ServerURL)
}
