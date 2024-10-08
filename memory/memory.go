package memory

import "github.com/sashabaranov/go-openai" // Importing the OpenAI package

// Message structure to store individual messages
type Message struct {
	Role    string // Role can be "user" or "assistant"
	Content string // The content of the message
}

// Memory structure to store conversation history
type Memory struct {
	conversationHistory []Message
}

// AddToMemory method to add a new message to memory
func (m *Memory) AddToMemory(role, content string) {
	m.conversationHistory = append(m.conversationHistory, Message{
		Role:    role,
		Content: content,
	})
}

// GetMemory method to retrieve the conversation history as a slice of ChatCompletionMessage
func (m *Memory) GetMemory() []openai.ChatCompletionMessage {
	var messages []openai.ChatCompletionMessage
	for _, msg := range m.conversationHistory {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}
	return messages
}
