package LLMAgent

import (
	"bufio"
	"context"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"llm-chain/memory"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/sashabaranov/go-openai"
)

func AgentWithMemory() {
	openAIClient := openai.NewClient("OPENAI-KEY")
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	memory := &memory.Memory{}
	c := colly.NewCollector()
	for {
		var userInput string
		fmt.Print("Ask a question or give a command: ")
		userInput, _ = reader.ReadString('\n')
		userInput = userInput[:len(userInput)-1]

		if strings.Contains(userInput, "solve") {
			equation := strings.TrimPrefix(userInput, "solve ")
			result, err := solveMath(equation)
			if err != nil {
				fmt.Printf("Error solving math: %s\n", err)
			}
			fmt.Printf("Math Result: %f\n", result)
		} else if isSearchRelated(userInput) {
			fmt.Println("Searching the web...")
			var knowledgeBase string
			getSearchResult(c, &knowledgeBase, userInput)

			// format knowledgeBase and userInput to ask the AI about the search results
			modifiedUserInput := fmt.Sprintf(`
				Use the provided knowledge base to answer the user's question.
				<knowledgebase>
					%s
				</knowledgebase>
				<user_question>
					%s
				</user_question>
			`, userInput, knowledgeBase)

			// print the modifiedUserInput
			// fmt.Printf("Modified User Input: %s\n", modifiedUserInput)
			response, err := askLLM(openAIClient, memory, modifiedUserInput)
			if err != nil {
				log.Fatalf("Error during completion: %v", err)
			}
			fmt.Printf("AI Response: %s\n", response)
		} else {
			response, err := askLLM(openAIClient, memory, userInput)
			if err != nil {
				log.Fatalf("Error during completion: %v", err)
			}
			fmt.Printf("AI Response: %s\n", response)
		}
	}
}

func isSearchRelated(input string) bool {
	keywords := []string{"search", "latest", "this year", "this month", "today", "future event", "current events", "now", "recent", "2024"}
	input = strings.ToLower(input)

	for _, keyword := range keywords {
		if strings.Contains(input, keyword) {
			return true
		}
	}
	return false
}

func getSearchResult(c *colly.Collector, knowledgeBase *string, userInput string) {
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")
	})

	c.OnHTML(".g", func(e *colly.HTMLElement) {
		content := e.ChildText("span")
		if content != "" {
			*knowledgeBase += content + "\n"
		}
	})

	query := userInput
	searchURL := "https://www.google.com/search?q=" + strings.ReplaceAll(query, " ", "+")

	if err := c.Visit(searchURL); err != nil {
		fmt.Println("Error making request:", err)
	}
	c.Wait()
}

func askLLM(client *openai.Client, memory *memory.Memory, userInput string) (string, error) {
	memory.AddToMemory(openai.ChatMessageRoleUser, userInput)
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
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}

func solveMath(equation string) (float64, error) {
	// Sanitize the input to only allow valid characters
	validInput := regexp.MustCompile(`^[0-9+\-*/().\s]+$`)
	if !validInput.MatchString(equation) {
		return 0, fmt.Errorf("invalid characters in equation")
	}

	// Remove any newlines or leading/trailing whitespace
	equation = strings.TrimSpace(equation)

	// Parse the expression
	node, err := parser.ParseExpr(equation)
	if err != nil {
		return 0, fmt.Errorf("error parsing equation: %s", err)
	}

	// Evaluate the expression
	return eval(node)
}

// eval recursively evaluates the parsed expression
func eval(node ast.Expr) (float64, error) {
	switch n := node.(type) {
	case *ast.BasicLit:
		var result float64
		_, err := fmt.Sscanf(n.Value, "%f", &result)
		if err != nil {
			return 0, fmt.Errorf("error evaluating number: %s", err)
		}
		return result, nil
	case *ast.BinaryExpr:
		left, err := eval(n.X)
		if err != nil {
			return 0, err
		}
		right, err := eval(n.Y)
		if err != nil {
			return 0, err
		}

		switch n.Op {
		case token.ADD:
			return left + right, nil
		case token.SUB:
			return left - right, nil
		case token.MUL:
			return left * right, nil
		case token.QUO:
			if right == 0 {
				return 0, fmt.Errorf("division by zero")
			}
			return left / right, nil
		default:
			return 0, fmt.Errorf("unsupported operator: %s", n.Op)
		}
	case *ast.ParenExpr:
		return eval(n.X)
	default:
		return 0, fmt.Errorf("unsupported expression type: %T", n)
	}
}
