package anthropic

import (
	"github.com/instill-ai/component/internal/util/httpclient"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/structpb"
)

type anthropicClient struct {
	httpClient *httpclient.Client
}

func newClient(apiKey string, baseURL string, logger *zap.Logger) *anthropicClient {
	c := httpclient.New("Anthropic", baseURL,
		httpclient.WithLogger(logger),
		httpclient.WithEndUserError(new(errBody)),
	)
	// Anthropic requires an API key to be set in the header "x-api-key" rather than normal "Authorization" header.
	c.Header.Set("X-Api-Key", apiKey)
	c.Header.Set("anthropic-version", "2023-06-01")

	return &anthropicClient{httpClient: c}
}

func (cl *anthropicClient) generateTextChat(request messagesReq) (messagesResp, error) {
	resp := messagesResp{}
	req := cl.httpClient.R().SetResult(&resp).SetBody(request)
	if _, err := req.Post(messagesPath); err != nil {
		return resp, err
	}
	return resp, nil
}

type errBody struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

func (e errBody) Message() string {
	return e.Error.Message
}

// getBasePath returns Anthropic's API URL. This configuration param allows us to
// override the API the connector will point to. It isn't meant to be exposed
// to users. Rather, it can serve to test the logic against a fake server.

// TODO instead of having the API value hardcoded in the codebase, it should be
// read from a setup file or environment variable.
// 2024-06-20 summer intern An-Che: This is the same implementation as the one in openai component, keep it as is for now?
func getBasePath(setup *structpb.Struct) string {
	v, ok := setup.GetFields()["base-path"]
	if !ok {
		return host
	}
	return v.GetStringValue()
}

func getAPIKey(setup *structpb.Struct) string {
	return setup.GetFields()[cfgAPIKey].GetStringValue()
}
