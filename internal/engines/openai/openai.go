package openai

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goura/yaku-cli/internal/ext/gen/openaichatcomp"
	"github.com/goura/yaku-cli/pkg/config"
	"github.com/goura/yaku-cli/pkg/types"
	"golang.org/x/text/language"
)

type OpenAIChatCompletionEngine struct {
	ServerURL    string
	openAIAPIKey string
}

func (t OpenAIChatCompletionEngine) Name() string {
	return "openai"
}

func (t *OpenAIChatCompletionEngine) LoadConfig(conf config.Config) error {
	if conf.OpenaiApiKey == "" {
		return fmt.Errorf("OpenaiApiKey is not set")
	}
	t.openAIAPIKey = conf.OpenaiApiKey
	return nil
}

func (t OpenAIChatCompletionEngine) IsSourceSizeOK(src string) bool {
	// TODO: improve
	return len(src) <= 1024*127
}

func (t OpenAIChatCompletionEngine) SupportedSourceLanguages() ([]language.Tag, error) {
	return []language.Tag{}, types.FeatureNotSupportedError
}

func (t OpenAIChatCompletionEngine) SupportedTargetLanguages(srcLang language.Tag) ([]language.Tag, error) {
	return []language.Tag{}, types.FeatureNotSupportedError
}

func (t OpenAIChatCompletionEngine) Translate(ctx context.Context, srcLang language.Tag, tgtLang language.Tag, src string) (string, error) {
	s, err := t.callChatCompAPI(ctx, srcLang, tgtLang, src)
	if err != nil {
		return "", err
	}

	return s, nil
}

func (t OpenAIChatCompletionEngine) callChatCompAPI(ctx context.Context, srcLang language.Tag, tgtLang language.Tag, textItem string) (string, error) {

	// Build the request
	//float32_zero := float32(0)

	initialCommandStr := fmt.Sprintf("You are a super translator API from %s to %s.  You must not reply anything except the translated sentence. Translate the following text:\n%s", srcLang, tgtLang, textItem)

	messages := []openaichatcomp.ChatCompletionRequestMessage{
		{
			Role:    "system",
			Content: initialCommandStr,
		},
	}

	logit_bias := map[string]interface{}{}

	float32_zero := float32(0)

	// https://platform.openai.com/docs/api-reference/chat/create
	req := openaichatcomp.CreateChatCompletionRequest{
		Model:    "gpt-3.5-turbo",
		Messages: messages,

		// The document says LogitBias is optional but the API returns an error without it
		LogitBias: &logit_bias,

		// The document says these default to 0 but the API returns an error without it
		PresencePenalty:  &float32_zero,
		FrequencyPenalty: &float32_zero,
	}
	apiReqBody := openaichatcomp.CreateChatCompletionJSONRequestBody(req)

	// https://platform.openai.com/docs/api-reference/authentication
	addAuthHdrFn := openaichatcomp.WithRequestEditorFn(
		func(ctx context.Context, req *http.Request) error {
			req.Header.Add("Authorization", "Bearer "+t.openAIAPIKey)
			return nil
		},
	)

	// Make a request
	serverURL := "https://api.openai.com/v1/"
	if t.ServerURL != "" {
		serverURL = t.ServerURL
	}

	cli, err := openaichatcomp.NewClientWithResponses(serverURL, addAuthHdrFn)
	if err != nil {
		return "", err
	}

	apiResp, err := cli.CreateChatCompletionWithResponse(ctx, apiReqBody)
	if err != nil {
		return "", err
	}

	if apiResp.StatusCode() != http.StatusOK {
		return "", fmt.Errorf("Error: Status code: %d Body: %v\n", apiResp.StatusCode(), string(apiResp.Body))
	}
	if apiResp.JSON200 == nil {
		return "", fmt.Errorf("Error: Status code was OK but response was nil: %v", apiResp.JSON200)
	}

	// Build the return value
	retval := ""

	choices := apiResp.JSON200.Choices
	for _, choice := range choices {
		if choice.Message == nil {
			// panic
			return "", fmt.Errorf("Error: Status code was OK but response was malformed: %v", choice)
		}
		retval = choice.Message.Content
		// TODO: parse the reason and continue requesting
	}

	return retval, nil
}

func (t *OpenAIChatCompletionEngine) SetEndpoint(endpoint string) error {
	t.ServerURL = endpoint
	return nil
}
