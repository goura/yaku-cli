package translator

import (
	"context"
	"testing"

	"github.com/goura/yaku-cli/pkg/config"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func TestBuildTranslatorOK(t *testing.T) {
	instance, err := BuildTranslator("dummyfortest")
	assert.Nil(t, err)

	ctx := context.Background()
	conf := config.Config{DummyAPIKey: "a"}
	s, err := instance.DoTranslation(ctx, conf, language.AmericanEnglish, language.Japanese, "Hello!")

	assert.Equal(t, "こんにちは！", s)
}

func TestBuildTranslatorNGNoConfig(t *testing.T) {
	instance, err := BuildTranslator("dummyfortest")
	assert.Nil(t, err)

	ctx := context.Background()
	conf := config.Config{} // Empty config should fail
	_, err = instance.DoTranslation(ctx, conf, language.AmericanEnglish, language.Japanese, "Hello!")
	assert.Error(t, err)
}

func TestBuildTranslatorNGSizeTooBig(t *testing.T) {
	instance, err := BuildTranslator("dummyfortest")
	assert.Nil(t, err)

	ctx := context.Background()
	conf := config.Config{DummyAPIKey: "a"}

	s := "" // Create a too long string to fail
	for i := 0; i < 100; i++ {
		s += "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	}
	_, err = instance.DoTranslation(ctx, conf, language.AmericanEnglish, language.Japanese, s)
	assert.Error(t, err)
}

func TestBuildTranslatorNGUnsupportedSourceLanguage(t *testing.T) {
	instance, err := BuildTranslator("dummyfortest")
	assert.Nil(t, err)

	ctx := context.Background()
	conf := config.Config{DummyAPIKey: "a"}

	s := "Selamat siang!" // Indonesian not supported
	_, err = instance.DoTranslation(ctx, conf, language.Indonesian, language.Japanese, s)
	assert.Error(t, err)
}

func TestBuildTranslatorNGUnsupportedTargetLanguage(t *testing.T) {
	instance, err := BuildTranslator("dummyfortest")
	assert.Nil(t, err)

	ctx := context.Background()
	conf := config.Config{DummyAPIKey: "a"}

	s := "Hello!" // English to Indonesian not supported
	_, err = instance.DoTranslation(ctx, conf, language.AmericanEnglish, language.Indonesian, s)
	assert.Error(t, err)
}
