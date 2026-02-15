package aimanager

import (
	"testing"

	"github.com/ricejson/gotool"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
)

func Test_ConvertToVolcengineMessages(t *testing.T) {
	testCases := []struct {
		name        string
		srcMessages []*Message
	}{
		{
			name: "success",
			srcMessages: []*Message{
				{
					Role:    UserRole,
					Content: "你好",
				},
				{
					Role:    AssistantRole,
					Content: "你好，我也好",
				},
				{
					Role:    SystemRole,
					Content: "你是一个有礼貌的人",
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			volcengineMessages := ConvertToVolcengineMessages(tc.srcMessages)
			require.True(t, len(tc.srcMessages) == len(volcengineMessages))
			for i, srcMessage := range tc.srcMessages {
				targetMessage := volcengineMessages[i]
				assert.Equal(t, string(srcMessage.Role), targetMessage.Role)
				assert.Equal(t, srcMessage.Content, cast.ToString(targetMessage.Content.StringValue))
			}
		})
	}
}

func Test_ConvertToMessages(t *testing.T) {
	testCases := []struct {
		name        string
		srcMessages []*model.ChatCompletionMessage
	}{
		{
			name: "success",
			srcMessages: []*model.ChatCompletionMessage{
				{
					Role: model.ChatMessageRoleUser,
					Content: &model.ChatCompletionMessageContent{
						StringValue: gotool.ToPtr("你好"),
					},
				},
				{
					Role: model.ChatMessageRoleAssistant,
					Content: &model.ChatCompletionMessageContent{
						StringValue: gotool.ToPtr("你好，我也好"),
					},
				},
				{
					Role: model.ChatMessageRoleAssistant,
					Content: &model.ChatCompletionMessageContent{
						StringValue: gotool.ToPtr("你是一个有礼貌的人"),
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			messages := ConvertToMessages(tc.srcMessages)
			require.True(t, len(tc.srcMessages) == len(messages))
			for i, srcMessage := range tc.srcMessages {
				targetMessage := messages[i]
				assert.Equal(t, srcMessage.Role, string(targetMessage.Role))
				assert.Equal(t, cast.ToString(srcMessage.Content.StringValue), targetMessage.Content)
			}
		})
	}
}
