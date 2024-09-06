package chroma

import (
	"github.com/instill-ai/component/internal/util/httpclient"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/structpb"
)

func newClient(setup *structpb.Struct, logger *zap.Logger) *httpclient.Client {
	c := httpclient.New("Chroma", getURL(setup),
		httpclient.WithLogger(logger),
	)

	c.SetHeader("Authorization", "Bearer "+getAPIKey(setup))
	c.SetHeader("Content-Type", "application/json")

	return c
}

func getURL(setup *structpb.Struct) string {
	return setup.GetFields()["url"].GetStringValue()
}

func getAPIKey(setup *structpb.Struct) string {
	return setup.GetFields()["api-key"].GetStringValue()
}
