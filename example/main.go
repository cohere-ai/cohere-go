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
	res, err := co.Generate("medium", cohere.GenerateOptions{
		Prompt:            prompt,
		MaxTokens:         20,
		Temperature:       1,
		K:                 5,
		P:                 0,
		StopSequences:     []string{"?"},
		ReturnLikelihoods: cohere.ReturnAll,
	})
	if err != nil {
		fmt.Println("An error occurred: ", err.Error())
		return
	}

	fmt.Println("Prompt: ", prompt)
	fmt.Println("Result: ", res.Generations[0].Text)
	fmt.Println("Likelihoods: ", res.Generations[0].TokenLikelihoods)
}
