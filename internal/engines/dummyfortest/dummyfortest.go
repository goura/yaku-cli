package dummyfortest

import (
	"context"
	"fmt"

	"github.com/goura/yaku-cli/pkg/config"
	"golang.org/x/text/language"
)

type DummyEngine struct {
}

func (t DummyEngine) Name() string {
	return "dummyfortest"
}

func (t DummyEngine) LoadConfig(conf config.Config) error {
	// This implementation doesn't need any config but we will check it for tests
	if conf.DummyAPIKey == "" {
		return fmt.Errorf("DummyAPIKey is not set")
	}
	return nil
}

// This implementation can only translate a source shorter than 100 chars
func (t DummyEngine) IsSourceSizeOK(src string) bool {
	return len(src) <= 100
}

// This implementation only supports American/British English and Japanese
func (t DummyEngine) SupportedSourceLanguages() []language.Tag {
	return []language.Tag{language.AmericanEnglish, language.BritishEnglish, language.Japanese}
}

// This implementation only supports certain combination
func (t DummyEngine) SupportedTargetLanguages(srcLang language.Tag) []language.Tag {
	if srcLang == language.AmericanEnglish || srcLang == language.BritishEnglish {
		return []language.Tag{language.Japanese}
	}
	if srcLang == language.Japanese {
		return []language.Tag{language.AmericanEnglish, language.BritishEnglish}
	}
	return []language.Tag{}
}

func (t DummyEngine) Translate(ctx context.Context, srcLang language.Tag, tgtLang language.Tag, src string) (string, error) {
	// This implementation is really dumb and can only translate few fixed sentences
	if src == "Hello!" &&
		srcLang == language.AmericanEnglish || srcLang == language.BritishEnglish && tgtLang == language.Japanese {
		return "こんにちは！", nil
	}

	if src == "こんにちは！" &&
		srcLang == language.Japanese &&
		tgtLang == language.AmericanEnglish || tgtLang == language.BritishEnglish {
		return "Hello!", nil
	}

	if src == "エレベーターはどこですか？" &&
		srcLang == language.Japanese &&
		tgtLang == language.AmericanEnglish {
		return "Where is the elevator?", nil
	}

	if src == "エレベーターはどこですか？" &&
		srcLang == language.Japanese &&
		tgtLang == language.BritishEnglish {
		return "Where is the lift?", nil
	}

	return "", fmt.Errorf("not implemented")
}

func (t DummyEngine) SetEndpoint(endpoint string) error {
	// This implementation doesn't need to set an endpoint
	return nil
}
