package main

import (
	"fmt"
	"os"

	"github.com/cohere-ai/cohere-go"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "API_KEY not specified")
		os.Exit(1)
	}

	co := cohere.CreateClient(apiKey)

	prompt := "What is your"
	fmt.Println("Prompt: ", prompt)
	// text, err := co.Generate(cohere.Orca, prompt, 10, 0.75)
	res, err := co.GenerateAdvanced(cohere.Seal, cohere.GenerateOptions{
		Prompt:            prompt,
		MaxTokens:         20,
		Temperature:       1,
		K:                 5,
		P:                 0,
		StopSequences:     []string{"?"},
		ReturnLikelihoods: cohere.ALL,
	})
	if err != nil {
		fmt.Println("An error occurred: ", err.Error())
		return
	}

	fmt.Println("Result: ", prompt+res.Text)
	fmt.Println("Likelihoods: ", res.TokenLikelihoods)
}
