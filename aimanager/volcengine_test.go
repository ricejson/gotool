package aimanager

import (
	"context"

	"github.com/bytedance/mockey"
	"github.com/go-playground/assert/v2"
	"github.com/ricejson/gotool"
	"github.com/stretchr/testify/require"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"

	"testing"
)

func Test_VolcengineClient(t *testing.T) {
	testCases := []struct {
		name     string
		client   *VolcengineClient
		messages []*Message
		mockFunc func()
	}{
		{
			name:   "success",
			client: NewVolcengineClient("", "", ""),
			messages: []*Message{
				{
					Role:    UserRole,
					Content: "你好",
				},
			},
			mockFunc: func() {
				mockey.Mock((*arkruntime.Client).CreateChatCompletion).Return(model.ChatCompletionResponse{
					Choices: []*model.ChatCompletionChoice{
						{
							Message: model.ChatCompletionMessage{
								Content: &model.ChatCompletionMessageContent{
									StringValue: gotool.ToPtr("你好"),
								},
							},
						},
					},
				}, nil).Build()
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFunc()
			resp, err := tc.client.Send(context.Background(), tc.messages)
			require.NoError(t, err)
			assert.Equal(t, resp, "你好")
		})
	}
}
