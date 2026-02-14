package aimanager

import (
	"context"
	"github.com/spf13/cast"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
)

type VolcengineClient struct {
	modelName string
	client    *arkruntime.Client
}

func NewVolcengineClient(apiKey, baseUrl, modelName string) *VolcengineClient {
	c := arkruntime.NewClientWithApiKey(
		apiKey,
		arkruntime.WithBaseUrl(baseUrl),
	)
	return &VolcengineClient{
		client:    c,
		modelName: modelName,
	}
}

func (c *VolcengineClient) Send(ctx context.Context, messages []*Message) (string, error) {
	// build req
	req := model.CreateChatCompletionRequest{
		Model:    c.modelName,
		Messages: ConvertToVolcengineMessages(messages),
	}
	resp, err := c.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", err
	}
	return cast.ToString(resp.Choices[0].Message.Content.StringValue), nil
}
