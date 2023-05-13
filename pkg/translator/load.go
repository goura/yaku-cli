package translator

import (
	"fmt"

	"github.com/goura/yaku-cli/internal/engines/deepl"
	"github.com/goura/yaku-cli/internal/engines/dummyfortest"
	"github.com/goura/yaku-cli/internal/engines/openai"
)

func NewDummyTranslator() TranslatorInstance {
	engine := dummyfortest.DummyEngine{}
	return NewTranslator(engine)
}

func NewDeepLTranslator() TranslatorInstance {
	engine := deepl.DeepLEngine{}
	return NewTranslator(&engine)
}

func NewOpenAITranslator() TranslatorInstance {
	engine := openai.OpenAIChatCompletionEngine{}
	return NewTranslator(&engine)
}

func BuildTranslator(engineTag string) (instance TranslatorInstance, err error) {
	// Instanciate a translator with the specified engine
	switch engineTag {
	case "dummyfortest":
		instance = NewDummyTranslator()
	case "deepl":
		instance = NewDeepLTranslator()
	case "openai":
		instance = NewOpenAITranslator()
	default:
		return instance, fmt.Errorf("engine:%s is not supported", engineTag)
	}
	return instance, nil
}
