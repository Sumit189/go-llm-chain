# LLM-Powered Agents with Memory in Go

This repository demonstrates how to build LLM-powered agents with memory using Go. The agents are designed to handle various tasks such as chat interactions, solving math equations, and web searches while maintaining conversational context.

## Features
- **LLM Chat with Memory**: Chatbot that remembers previous conversations, providing more contextual responses.
- **Agent with Memory**: Autonomous agent capable of solving problems and automating tasks, such as web searches and math problem-solving, with memory retention.

## Project Structure
<pre>
root/
│   main.go             # Main entry point to select the desired functionality.
└───llm-agent/
    │   llmChatWithMemory.go # Chat interaction with memory.
    │   agentWithMemory.go   # Automated agent with memory to solve problems.
</pre>

## Getting Started

### Prerequisites
- Go version 1.23 or higher
- OpenAI API key for LLM interactions

### Installation
1. Clone the repository:
   ```bash
   git clone git@github.com:Sumit189/go-llm-chain.git
   cd go-llm-chain
   ```
2. Install the required dependencies
    ```bash
    go mod tidy
    ```
### Usage
1. Run the application:
    ```bash
    go run main.go
    ```
2. Select the desired functionality:
    - To interact with the LLM chatbot with memory, enter `1`.
    - To utilize the Agent (with Web Search capabilities and Equation Evaluation) that remembers past interactions, enter `2`.

### Contributing
Feel free to fork the repository and submit pull requests for any improvements or new features!
