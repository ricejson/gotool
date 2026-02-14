package aimanager

import (
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
)

type Role string

const (
	UserRole      Role = "user"
	AssistantRole Role = "assistant"
	SystemRole    Role = "system"
)

type Message struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

func ConvertToVolcengineMessages(messages []*Message) []*model.ChatCompletionMessage {
	volcengineMessages := make([]*model.ChatCompletionMessage, 0, len(messages))
	for _, message := range messages {
		volcengineMessages = append(volcengineMessages, &model.ChatCompletionMessage{
			Role: string(message.Role),
			Content: &model.ChatCompletionMessageContent{
				StringValue: volcengine.String(message.Content),
			},
		})
	}
	return volcengineMessages
}

func ConvertToMessages(chatMessages []*model.ChatCompletionMessage) []*Message {
	messages := make([]*Message, 0, len(chatMessages))
	for _, chatMessage := range chatMessages {
		messages = append(messages, &Message{
			Role:    Role(chatMessage.Role),
			Content: volcengine.StringValue(chatMessage.Content.StringValue),
		})
	}
	return messages
}
