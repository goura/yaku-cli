package deepl

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/goura/yaku-cli/internal/ext/gen/deepl"
	"github.com/goura/yaku-cli/pkg/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

// MockClient is a mock implementation of the deepl.Client interface
type MockClient struct {
	mock.Mock
}

// TranslateTextWithFormdataBodyWithResponse is a mock implementation of the deepl.Client.TranslateTextWithFormdataBodyWithResponse method
func (m *MockClient) TranslateTextWithFormdataBodyWithResponse(ctx context.Context, body deepl.TranslateTextFormdataRequestBody, reqEditors ...deepl.RequestEditorFn) (*deepl.TranslateTextResponse, error) {
	args := m.Called(ctx, body, reqEditors)
	return args.Get(0).(*deepl.TranslateTextResponse), args.Error(1)
}

func buildTranslateTextExpectedResponse(statusCode int, detectedLanguage string, text string) (*deepl.TranslateTextResponse, error) {
	apiResp := deepl.TranslateTextResponse{}

	bodyBytes := []byte(`{"translations":[{"detected_source_language":"` + detectedLanguage + `","text":"` + text + `"}]}`)

	var dest struct {
		Translations *[]struct {
			DetectedSourceLanguage *deepl.SourceLanguage `json:"detected_source_language,omitempty"`

			Text *string `json:"text,omitempty"`
		} `json:"translations,omitempty"`
	}

	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return &apiResp, err
	}

	apiResp.HTTPResponse = &http.Response{
		StatusCode: statusCode,
	}
	apiResp.JSON200 = &dest
	apiResp.Body = []byte{}

	return &apiResp, nil
}

func TestCallDeepLAPI(t *testing.T) {
	// Create an instance of the DeepLEngine
	engine := DeepLEngine{
		deepLAuthKey: "your-auth-key",
		ServerURL:    "https://api.deepl.com/v2/",
	}

	// Create a mock client
	mockClient := new(MockClient)

	// Create an expected response
	expectedResponse, err := buildTranslateTextExpectedResponse(200, "EN", "世界、こんにちは！")

	// Set up the expected behavior of the mock client
	mockClient.On("TranslateTextWithFormdataBodyWithResponse", mock.Anything, mock.Anything, mock.Anything).Return(expectedResponse, nil)

	// Replace the actual client with the mock client
	engine.cli = mockClient

	// Define the test case inputs
	sourceLanguage := deepl.SourceLanguage("en")
	targetLanguage := deepl.TargetLanguage("ja")
	textItem := "Hello, world!"

	// Call the function being tested
	result, err := engine.callDeepLAPI(context.Background(), sourceLanguage, targetLanguage, textItem)

	// Assert the expected behavior
	assert.NoError(t, err)

	// Assert the expected translation result
	assert.Equal(t, "世界、こんにちは！", result)
}
