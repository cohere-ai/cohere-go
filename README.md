# Cohere Go SDK

This package provides functionality developed to simplify interfacing with the [cohere.ai](https://cohere.ai/) natural language API in Go.

## Documentation

See the [API's documentation](https://docs.cohere.ai/).

Also see some code examples for the SDK [here](https://github.com/cohere-ai/cohere-go/blob/main/example/main.go).

## Installation

```
go get github.com/cohere-ai/cohere-go
```

### Requirements

- Go 1.17+

## Usage

To use this library, you must have an API key and specify it as a string when creating the `cohere.Client` struct. API keys can be created through the Cohere CLI or Playground. This is a basic example of the creating the client and using the `Generate` endpoint.

```go
package main

import (
	"fmt"

	"github.com/cohere-ai/cohere-go"
)

func main() {
	co, err := cohere.CreateClient("YOUR_API_KEY")
	if err != nil {
		fmt.Println(err)
		return
	}

	prompt := "Tell me a joke"
	res, err := co.Generate(cohere.GenerateOptions{
		Model:       "medium",
		Prompt:      prompt,
		MaxTokens:   cohere.Uint(100),
		Temperature: cohere.Float64(0.75),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Prediction:\n%s%s\n", prompt, res.Generations[0].Text)
}
```

or similarly you can stream results instead of blocking on complete generation:

```go
package main

import (
	"fmt"
	"os"

	"github.com/cohere-ai/cohere-go"
)

func main() {
	co, err := cohere.CreateClient("YOUR_API_KEY")
	if err != nil {
		fmt.Println(err)
		return
	}

	prompt := `Write 5 titles for a blog ideas for the keywords "large language model" or "text generation"`
	ch := co.Stream(cohere.GenerateOptions{
		Model:       "xlarge",
		Prompt:      prompt,
		MaxTokens:   cohere.Uint(200),
		Temperature: cohere.Float64(0.75),
	})

	fmt.Printf("Completion:\n%s", prompt)
	for res := range ch {
		if res.Err != nil {
			fmt.Println(res.Err)
			break
		}
		fmt.Printf(res.Token.Token)
	}
  	fmt.Println()
}
```

A more complete example of `Generate` can be found [here](https://github.com/cohere-ai/cohere-go/blob/main/example/main.go) and example usage of other endpoints can be found [here](https://github.com/cohere-ai/cohere-go/blob/main/client_test.go).

## Versioning

This SDK supports the latest API version. For more information, please refer to the [Versioning Docs](https://docs.cohere.ai/reference/versioning).

## Endpoints

For a full breakdown of endpoints and arguments, please consult the [Cohere Docs](https://docs.cohere.ai/).

| Cohere Endpoint  | Function            |
| ---------------- | ------------------- |
| /generate        | co.Generate()       |
| /embed           | co.Embed()          |
| /classify        | co.Classify()       |
| /summarize       | co.Summarize()      |
| /tokenize        | co.Tokenize()       |
| /detokenize      | co.Detokenize()     |
| /detect-language | co.DetectLanguage() |

## Models

To view an up-to-date list of available models please consult the [Cohere CLI](https://docs.cohere.ai/command/). To get started try out `large`.

## Responses

All of the endpoint functions will return a Cohere object corresponding to the endpoint (e.g. for generation, it would be `GenerateResponse`). The responses can be found as fields on the struct (e.g. generations would be `GenerateResponse.Generations`). The names of these fields and a detailed breakdown of the response body can be found in the [Cohere Docs](https://docs.cohere.ai/).

## Errors

Unsuccessful API calls from the SDK will return an error. Please see the documentation's page on [errors](https://docs.cohere.ai/errors-reference) for more information about what the errors mean.
