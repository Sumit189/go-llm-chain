package main

import (
	"fmt"

	"llm-chain/llm-agent"
)

func main() {
	fmt.Println("What would you like to try?")
	fmt.Println("1. LLM Chat")
	fmt.Println("2. LLM Chat with Memory")
	fmt.Println("3. Agent with Memory")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		llmAgent.LLMChat()
	case 2:
		llmAgent.LLMChatWithMemory()
	case 3:
		llmAgent.AgentWithMemory()
	default:
		fmt.Println("Invalid choice")
	}

}
