package translator

import (
	"github.com/goura/yaku-cli/internal/engines/deepl"
	"github.com/goura/yaku-cli/internal/engines/dummyfortest"
)

func NewDummyTranslator() TranslatorInstance {
	engine := dummyfortest.DummyEngine{}
	return NewTranslator(engine)
}

func NewDeepLTranslator() TranslatorInstance {
	engine := deepl.DeepLEngine{}
	return NewTranslator(&engine)
}
