package LLMAgent

import (
	"bufio"
	"context"
	"fmt"
	"llm-chain/memory"
	"log"
	"os"

	"github.com/sashabaranov/go-openai"
)

func LLMChatWithMemory() {
	// Initialize OpenAI client and memory
	client := openai.NewClient("OPENAI-KEY")
	memory := &memory.Memory{}
	// Create a buffered reader for user input
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	for {
		// Simulate a conversation
		var userInput string
		fmt.Print("You: ")

		// Use reader to capture the user input safely
		userInput, _ = reader.ReadString('\n')
		userInput = userInput[:len(userInput)-1]

		// Add the user's message to memory
		memory.AddToMemory(openai.ChatMessageRoleUser, userInput)

		// Get AI response
		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:       openai.GPT4,
				Messages:    memory.GetMemory(),
				MaxTokens:   100,
				Temperature: 0.7,
			},
		)
		if err != nil {
			log.Fatalf("Error during completion: %v", err)
		}

		// Output and store the response in memory
		aiResponse := resp.Choices[0].Message.Content
		fmt.Printf("GenAI: %s\n", aiResponse)
		memory.AddToMemory(openai.ChatMessageRoleSystem, aiResponse)
	}
}
