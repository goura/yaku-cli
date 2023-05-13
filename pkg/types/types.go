package types

import (
	"context"

	"github.com/goura/yaku-cli/pkg/config"
	"golang.org/x/text/language"
)

type TranslationEngine interface {
	// Engine name for tagging/logging
	Name() string

	// Load config such as API key, authentication token, etc.
	LoadConfig(conf config.Config) error

	// Return supported source languages supported by the module
	SupportedSourceLanguages() ([]language.Tag, error)

	// Return supported target languages supported by the module,
	// for a given source language
	SupportedTargetLanguages(srcLang language.Tag) ([]language.Tag, error)

	// Check if the source string is within the size limit
	IsSourceSizeOK(src string) bool

	// Do the translation
	Translate(ctx context.Context, srcLang language.Tag, tgtLang language.Tag, src string) (string, error)

	// Set an endpoint for testing, etc
	// Expecting a URL at this moment, but allowing string
	SetEndpoint(endpoint string) error
}
