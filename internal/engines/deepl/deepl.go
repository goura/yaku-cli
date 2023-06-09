package deepl

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goura/yaku-cli/internal/ext/gen/deepl"
	"github.com/goura/yaku-cli/pkg/config"
	"golang.org/x/text/language"
)

type DeepLAPIClient interface {
	TranslateTextWithFormdataBodyWithResponse(ctx context.Context, body deepl.TranslateTextFormdataRequestBody, reqEditors ...deepl.RequestEditorFn) (*deepl.TranslateTextResponse, error)
}

type DeepLEngine struct {
	ServerURL    string
	deepLAuthKey string
	cli          DeepLAPIClient // Mostly for testing
}

func (t DeepLEngine) Name() string {
	return "deepl"
}

func (t *DeepLEngine) LoadConfig(conf config.Config) error {
	if conf.DeeplAuthKey == "" {
		return fmt.Errorf("DeepLAuthKey is not set")
	}
	t.deepLAuthKey = conf.DeeplAuthKey
	return nil
}

func (t DeepLEngine) IsSourceSizeOK(src string) bool {
	// TODO: improve
	return len(src) <= 1024*127
}

func (t DeepLEngine) SupportedSourceLanguages() ([]language.Tag, error) {
	return supportedSourceLanguages(), nil
}

func (t DeepLEngine) SupportedTargetLanguages(srcLang language.Tag) ([]language.Tag, error) {
	return supportedTargetLanguages(), nil
}

func (t DeepLEngine) Translate(ctx context.Context, srcLang language.Tag, tgtLang language.Tag, src string) (string, error) {
	sourceLanguage, err := srcLanguageTagToDeepLSourceLanguage(srcLang)
	if err != nil {
		return "", err
	}

	targetLanguage, err := tgtLanguageTagToDeepLTargetLanguage(tgtLang)
	if err != nil {
		return "", err
	}

	s, err := t.callDeepLAPI(ctx, sourceLanguage, targetLanguage, src)
	if err != nil {
		return "", err
	}

	return s, nil
}

func (t DeepLEngine) callDeepLAPI(ctx context.Context, sourceLanguage deepl.SourceLanguage, targetLanguage deepl.TargetLanguage, textItem string) (string, error) {

	// Build the request
	formality := deepl.Formality("default")

	apiReqBody := deepl.TranslateTextFormdataBody{
		Formality:  &formality,
		SourceLang: &sourceLanguage,
		TargetLang: targetLanguage,
		Text:       textItem,
	}

	// https://www.deepl.com/docs-api
	addAuthHdrFn := deepl.WithRequestEditorFn(
		func(ctx context.Context, req *http.Request) error {
			req.Header.Add("Authorization", "DeepL-Auth-Key "+t.deepLAuthKey)
			return nil
		},
	)

	// Make a request
	serverURL := "https://api.deepl.com/v2/"
	if t.ServerURL != "" {
		serverURL = t.ServerURL
	}

	cli, err := deepl.NewClientWithResponses(serverURL, addAuthHdrFn)
	if err != nil {
		return "", err
	}

	// Use the already set cli if it's already set (for testing)
	if t.cli == nil {
		t.cli = cli
	}

	apiResp, err := t.cli.TranslateTextWithFormdataBodyWithResponse(ctx, deepl.TranslateTextFormdataRequestBody(apiReqBody))
	if err != nil {
		return "", err
	}

	if apiResp.StatusCode() != http.StatusOK {
		return "", fmt.Errorf("Error: Status code: %d Body: %v\n", apiResp.StatusCode(), string(apiResp.Body))
	}
	if apiResp.JSON200.Translations == nil {
		return "", fmt.Errorf("Error: Status code was OK but Translations was nil: %v", apiResp.JSON200)
	}

	// Build the return value
	retval := ""

	for _, translation := range *apiResp.JSON200.Translations {
		if translation.Text != nil {
			retval += *translation.Text
		}
	}

	return retval, nil
}

func (t *DeepLEngine) SetEndpoint(endpoint string) error {
	t.ServerURL = endpoint
	return nil
}
