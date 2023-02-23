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

  res, err := co.Generate(cohere.GenerateOptions{
    Model:             "large",
    Prompt:            "co:here",
    MaxTokens:         10,
    Temperature:       0.75,
  })
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println("Prediction: ", res.Generations[0].Text)
}
```

A more complete example of `Generate` can be found [here](https://github.com/cohere-ai/cohere-go/blob/main/example/main.go) and example usage of other endpoints can be found [here](https://github.com/cohere-ai/cohere-go/blob/main/client_test.go).

## Versioning

To use the SDK with a specific API version, you can modify the client with the desired version as such:

```go
package main

import (
  "github.com/cohere-ai/cohere-go"
)

func main() {
  co, err := cohere.CreateClient("YOUR_API_KEY")
  if err != nil {
    fmt.Println(err)
    return
  }

  co.Version = "2022-12-06"
}
```

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
