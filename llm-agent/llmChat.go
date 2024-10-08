package llmAgent

import (
	"context"
	"fmt"
	"log"

	"github.com/sashabaranov/go-openai"
)

func LLMChat() {
	// Initialize OpenAI client with your API key
	client := openai.NewClient("OPENAI-KEY")

	// Sample document to analyze
	document := "Artificial intelligence is transforming industries from healthcare to autonomous vehicles. It enables machines to perform tasks that require human intelligence."

	// Step 1: Analyze the document
	summary, err := analyzeDocument(client, document)
	if err != nil {
		log.Fatalf("Error analyzing document: %v", err)
	}
	fmt.Println("Document Summary:", summary)

	// Step 2: Generate follow-up questions based on the summary
	questions, err := generateQuestions(client, summary)
	if err != nil {
		log.Fatalf("Error generating questions: %v", err)
	}
	fmt.Println("Follow-up Questions:", questions)
}

// Function to analyze a document using OpenAI GPT
func analyzeDocument(client *openai.Client, document string) (string, error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("Analyze the following document and summarize it:\n\n%s", document),
				},
			},
			MaxTokens:   100,
			Temperature: 0.7,
		},
	)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}

// Function to generate follow-up questions based on the summary
func generateQuestions(client *openai.Client, summary string) (string, error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("Generate follow-up questions based on the following summary:\n\n%s", summary),
				},
			},
			MaxTokens:   100,
			Temperature: 0.7,
		},
	)

	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, err
}
