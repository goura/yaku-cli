package deepl

import (
	"fmt"

	"github.com/goura/yaku-cli/internal/ext/gen/deepl"
	"golang.org/x/text/language"
)

// srcLanguageTagToDeepLSourceLanguage converts golang.org/x/text/language.Tag to deepl.SourceLanguage
func srcLanguageTagToDeepLSourceLanguage(tag language.Tag) (deepl.SourceLanguage, error) {
	switch tag {
	case language.Bulgarian:
		return deepl.SourceLanguageBG, nil
	case language.Czech:
		return deepl.SourceLanguageCS, nil
	case language.Danish:
		return deepl.SourceLanguageDA, nil
	case language.German:
		return deepl.SourceLanguageDE, nil
	case language.Greek:
		return deepl.SourceLanguageEL, nil

	// English variants all EN
	case language.English:
		return deepl.SourceLanguageEN, nil
	case language.AmericanEnglish:
		return deepl.SourceLanguageEN, nil
	case language.BritishEnglish:
		return deepl.SourceLanguageEN, nil

	case language.Spanish:
		return deepl.SourceLanguageES, nil
	case language.Estonian:
		return deepl.SourceLanguageET, nil
	case language.Finnish:
		return deepl.SourceLanguageFI, nil
	case language.French:
		return deepl.SourceLanguageFR, nil
	case language.Hungarian:
		return deepl.SourceLanguageHU, nil
	case language.Indonesian:
		return deepl.SourceLanguageID, nil
	case language.Italian:
		return deepl.SourceLanguageIT, nil
	case language.Japanese:
		return deepl.SourceLanguageJA, nil
	case language.Korean:
		return deepl.SourceLanguageKO, nil
	case language.Lithuanian:
		return deepl.SourceLanguageLT, nil
	case language.Latvian:
		return deepl.SourceLanguageLV, nil
	case language.Norwegian:
		return deepl.SourceLanguageNB, nil
	case language.Dutch:
		return deepl.SourceLanguageNL, nil
	case language.Polish:
		return deepl.SourceLanguagePL, nil

	// Portuguse-variant all PT
	case language.Portuguese:
		return deepl.SourceLanguagePT, nil
	case language.BrazilianPortuguese:
		return deepl.SourceLanguagePT, nil
	case language.EuropeanPortuguese:
		return deepl.SourceLanguagePT, nil

	case language.Romanian:
		return deepl.SourceLanguageRO, nil
	case language.Russian:
		return deepl.SourceLanguageRU, nil
	case language.Slovak:
		return deepl.SourceLanguageSK, nil
	case language.Slovenian:
		return deepl.SourceLanguageSL, nil
	case language.Swedish:
		return deepl.SourceLanguageSV, nil
	case language.Turkish:
		return deepl.SourceLanguageTR, nil
	case language.Ukrainian:
		return deepl.SourceLanguageUK, nil
	case language.Chinese:
		return deepl.SourceLanguageZH, nil
	}
	return deepl.SourceLanguageEN, fmt.Errorf("source language:%v is not supported by DeepL", tag)
}

// tgtLanguageTagToDeepLTargetLanguage converts golang.org/x/text/language.Tag to deepl.TargetLanguage
func tgtLanguageTagToDeepLTargetLanguage(tag language.Tag) (deepl.TargetLanguage, error) {
	switch tag {
	case language.Bulgarian:
		return deepl.TargetLanguageBG, nil
	case language.Czech:
		return deepl.TargetLanguageCS, nil
	case language.Danish:
		return deepl.TargetLanguageDA, nil
	case language.German:
		return deepl.TargetLanguageDE, nil
	case language.Greek:
		return deepl.TargetLanguageEL, nil

	// English falls back to ENUS
	case language.English:
		return deepl.TargetLanguageENUS, nil

	case language.AmericanEnglish:
		return deepl.TargetLanguageENUS, nil
	case language.BritishEnglish:
		return deepl.TargetLanguageENGB, nil
	case language.Spanish:
		return deepl.TargetLanguageES, nil
	case language.Estonian:
		return deepl.TargetLanguageET, nil
	case language.Finnish:
		return deepl.TargetLanguageFI, nil
	case language.French:
		return deepl.TargetLanguageFR, nil
	case language.Hungarian:
		return deepl.TargetLanguageHU, nil
	case language.Indonesian:
		return deepl.TargetLanguageID, nil
	case language.Italian:
		return deepl.TargetLanguageIT, nil
	case language.Japanese:
		return deepl.TargetLanguageJA, nil
	case language.Korean:
		return deepl.TargetLanguageKO, nil
	case language.Lithuanian:
		return deepl.TargetLanguageLT, nil
	case language.Latvian:
		return deepl.TargetLanguageLV, nil
	case language.Norwegian:
		return deepl.TargetLanguageNB, nil
	case language.Dutch:
		return deepl.TargetLanguageNL, nil
	case language.Polish:
		return deepl.TargetLanguagePL, nil
	case language.BrazilianPortuguese:
		return deepl.TargetLanguagePTBR, nil

	// Portuguese falls back to PTPT
	case language.Portuguese:
		return deepl.TargetLanguagePTPT, nil

	case language.EuropeanPortuguese:
		return deepl.TargetLanguagePTPT, nil
	case language.Romanian:
		return deepl.TargetLanguageRO, nil
	case language.Russian:
		return deepl.TargetLanguageRU, nil
	case language.Slovak:
		return deepl.TargetLanguageSK, nil
	case language.Slovenian:
		return deepl.TargetLanguageSL, nil
	case language.Swedish:
		return deepl.TargetLanguageSV, nil
	case language.Turkish:
		return deepl.TargetLanguageTR, nil
	case language.Ukrainian:
		return deepl.TargetLanguageUK, nil
	case language.Chinese:
		return deepl.TargetLanguageZH, nil
	}
	return deepl.TargetLanguageENUS, fmt.Errorf("source language:%v is not supported by DeepL", tag)
}

func supportedSourceLanguages() []language.Tag {
	return []language.Tag{
		language.Bulgarian,
		language.Czech,
		language.Danish,
		language.German,
		language.Greek,
		language.English,
		language.AmericanEnglish,
		language.BritishEnglish,
		language.Spanish,
		language.Estonian,
		language.Finnish,
		language.French,
		language.Hungarian,
		language.Indonesian,
		language.Italian,
		language.Japanese,
		language.Korean,
		language.Lithuanian,
		language.Latvian,
		language.Norwegian,
		language.Dutch,
		language.Polish,
		language.Portuguese,
		language.BrazilianPortuguese,
		language.EuropeanPortuguese,
		language.Romanian,
		language.Russian,
		language.Slovak,
		language.Slovenian,
		language.Swedish,
		language.Turkish,
		language.Ukrainian,
		language.Chinese,
	}
}

func supportedTargetLanguages() []language.Tag {
	return []language.Tag{
		language.Bulgarian,
		language.Czech,
		language.Danish,
		language.German,
		language.Greek,
		language.English,
		language.AmericanEnglish,
		language.BritishEnglish,
		language.Spanish,
		language.Estonian,
		language.Finnish,
		language.French,
		language.Hungarian,
		language.Indonesian,
		language.Italian,
		language.Japanese,
		language.Korean,
		language.Lithuanian,
		language.Latvian,
		language.Norwegian,
		language.Dutch,
		language.Polish,
		language.BrazilianPortuguese,
		language.Portuguese,
		language.EuropeanPortuguese,
		language.Romanian,
		language.Russian,
		language.Slovak,
		language.Slovenian,
		language.Swedish,
		language.Turkish,
		language.Ukrainian,
		language.Chinese,
	}
}
