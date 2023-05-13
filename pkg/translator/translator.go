package translator

import (
	"context"
	"errors"
	"fmt"

	"github.com/goura/yaku-cli/pkg/config"
	"github.com/goura/yaku-cli/pkg/types"
	"github.com/thoas/go-funk"
	"golang.org/x/text/language"
)

type TranslatorInstance struct {
	engine types.TranslationEngine
}

func (t TranslatorInstance) DoTranslation(ctx context.Context, conf config.Config, srcLang language.Tag, tgtLang language.Tag, src string) (string, error) {
	name := t.engine.Name()

	if err := t.engine.LoadConfig(conf); err != nil {
		return "", fmt.Errorf("engine:%s doesn't have enough config. err:%v", name, err)
	}

	if !t.engine.IsSourceSizeOK(src) {
		return "", fmt.Errorf("engine:%s source string is too long", name)
	}

	srcLangTags, err := t.engine.SupportedSourceLanguages()
	if errors.Is(err, types.FeatureNotSupportedError) {
		// pass
	} else if err != nil {
		return "", fmt.Errorf("engine:%s error while checking source language", err)
	} else if !funk.Contains(srcLangTags, srcLang) {
		return "", fmt.Errorf("engine:%s doesn't support %v as source language", name, srcLang)
	}

	tgtLangTags, err := t.engine.SupportedTargetLanguages(srcLang)
	if errors.Is(err, types.FeatureNotSupportedError) {
		// pass
	} else if err != nil {
		return "", fmt.Errorf("engine:%s error while checking target language", err)
	} else if !funk.Contains(tgtLangTags, tgtLang) {
		return "", fmt.Errorf("translating source:%v to target:%v is not supported by %s", srcLang, tgtLang, name)
	}

	return t.engine.Translate(ctx, srcLang, tgtLang, src)
}

func NewTranslator(engine types.TranslationEngine) TranslatorInstance {
	t := TranslatorInstance{}
	t.engine = engine
	return t
}
