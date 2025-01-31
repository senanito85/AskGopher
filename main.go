package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
)

func GetResponse(client *genai.Client, ctx context.Context, question string) {
	// Initialize Gemini Pro model
	model := client.GenerativeModel("gemini-pro")

	// Configure the model
	model.SetTemperature(0.7)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(2048)

	// Generate content
	resp, err := model.GenerateContent(ctx, genai.Text(question))
	if err != nil {
		fmt.Printf("Error generating content: %v\n", err)
		return
	}

	// Print the response
	if len(resp.Candidates) > 0 {
		if len(resp.Candidates[0].Content.Parts) > 0 {
			fmt.Println(resp.Candidates[0].Content.Parts[0])
		}
	}
	fmt.Printf("\n")
}

type NullWriter int

func (NullWriter) Write([]byte) (int, error) { return 0, nil }

func main() {
	log.SetOutput(NullWriter(0))
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading .env file: %v", err)
	}

	apiKey := viper.GetString("GOOGLE_API_KEY")
	if apiKey == "" {
		log.Fatalln("Google API key not found")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	root := &cobra.Command{
		Use:   "Gemini",
		Short: "Chat with Google Gemini in Terminal",
		Run: func(cmd *cobra.Command, args []string) {
			scanner := bufio.NewScanner(os.Stdin)
			quit := false
			fmt.Println("Ask Gopher! 🦫")
			fmt.Println("------------------")
			for !quit {
				fmt.Print("\n💭Ask something (or 'quit' to end): ")
				if !scanner.Scan() {
					break
				}
				question := scanner.Text()
				switch question {
				case "quit":
					quit = true
					fmt.Println("👋 Goodbye!")
				case "":
					continue
				default:
					fmt.Println("\n🤖 Answering:")
					GetResponse(client, ctx, question)
				}
			}
		},
	}

	root.Execute()
}
