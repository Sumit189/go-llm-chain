package main

import (
	"fmt"

	"llm-chain/llm-agent"
)

func main() {
	fmt.Println("What would you like to try?")
	fmt.Println("1. LLM Chat with Memory")
	fmt.Println("2. Agent with Memory")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		LLMAgent.LLMChatWithMemory()
	case 2:
		LLMAgent.AgentWithMemory()
	default:
		fmt.Println("Invalid choice")
	}

}
