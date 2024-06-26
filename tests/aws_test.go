package tests

import (
	"context"
	"errors"
	"io"
	"testing"

	cohere "github.com/cohere-ai/cohere-go/v2"
	client "github.com/cohere-ai/cohere-go/v2/client"
	"github.com/cohere-ai/cohere-go/v2/core"
	"github.com/stretchr/testify/require"
)

func Generate(t *testing.T, model *string, client client.Client) {
	prediction, err := client.Generate(
		context.TODO(),
		&cohere.GenerateRequest{
			Model:  model,
			Prompt: "count with me!",
		},
	)

	require.NoError(t, err)
	print(prediction)
}

func GenerateStream(t *testing.T, model *string, client client.Client) {
	t.Skip("Skip until auth is set up")
	stream, err := client.GenerateStream(
		context.TODO(),
		&cohere.GenerateStreamRequest{
			Model:  model,
			Prompt: "Cohere is",
		},
	)

	require.NoError(t, err)

	// Make sure to close the stream when you're done reading.
	// This is easily handled with defer.
	defer stream.Close()

	for {
		message, err := stream.Recv()

		if errors.Is(err, io.EOF) {
			// An io.EOF error means the server is done sending messages
			// and should be treated as a success.
			break
		}

		if message.TextGeneration != nil {
			print(message.TextGeneration.Text)
		}
	}
}

// Test Chat
func Chat(t *testing.T, model *string, client client.Client) {
	t.Skip("Skip until auth is set up")
	chat, err := client.Chat(
		context.TODO(),
		&cohere.ChatRequest{
			Model:   model,
			Message: "2",
		},
	)

	require.NotEmpty(t, chat.Text)
	require.NotEmpty(t, chat.GenerationId)
	require.NotEmpty(t, chat.FinishReason)

	require.NoError(t, err)
	print(chat)
}

// Test ChatStream
func ChatStream(t *testing.T, model *string, client client.Client) {
	t.Skip("Skip until auth is set up")
	stream, err := client.ChatStream(
		context.TODO(),
		&cohere.ChatStreamRequest{
			Model:   model,
			Message: "Cohere is",
		},
	)

	typeSet := make(map[string]bool)

	require.NoError(t, err)

	// Make sure to close the stream when you're done reading.
	// This is easily handled with defer.
	defer stream.Close()

	for {
		message, err := stream.Recv()

		if errors.Is(err, io.EOF) {
			// An io.EOF error means the server is done sending messages
			// and should be treated as a success.
			break
		}

		typeSet[message.EventType] = true

		if message.TextGeneration != nil {
			// print(message.TextGeneration.Text)
			require.NotEmpty(t, message.TextGeneration.Text)
		} else if message.StreamEnd != nil {
			require.NotEmpty(t, message.StreamEnd.FinishReason)
			require.NotEmpty(t, message.StreamEnd.Response.Text)
		}
	}

	require.Contains(t, typeSet, "text-generation")
	require.Contains(t, typeSet, "stream-end")
}

func Rerank(t *testing.T, model *string, client client.Client) {
	t.Skip("Skip until auth is set up")
	rerank, err := client.Rerank(
		context.TODO(),
		&cohere.RerankRequest{
			Model: model,
			Query: "What is the capital of the United States?",
			Documents: []*cohere.RerankRequestDocumentsItem{
				{String: "Carson City is the capital city of the American state of Nevada."},
				{String: "The Commonwealth of the Northern Mariana Islands is a group of islands in the Pacific Ocean. Its capital is Saipan."},
				{String: "Washington, D.C. (also known as simply Washington or D.C., and officially as the District of Columbia) is the capital of the United States. It is a federal district."},
				{String: "Capital punishment (the death penalty) has existed in the United States since beforethe United States was a country. As of 2017, capital punishment is legal in 30 of the 50 states."},
			},
		})

	require.NoError(t, err)
	print(rerank)
}

func Embed(t *testing.T, model *string, client client.Client) {
	t.Skip("Skip until auth is set up")
	embed, err := client.Embed(
		context.TODO(),
		&cohere.EmbedRequest{
			Model:     model,
			Texts:     []string{"hello", "goodbye"},
			InputType: cohere.EmbedInputTypeSearchDocument.Ptr(),
		})

	require.NoError(t, err)
	print(embed)
}

func Tool(t *testing.T, model *string, client client.Client) {
	t.Skip("Skip until auth is set up")
	tools := []*cohere.Tool{
		{
			Name:        "sales_database",
			Description: "Connects to a database about sales volumes",
			ParameterDefinitions: map[string]*cohere.ToolParameterDefinitionsValue{
				"day": {
					Description: cohere.String("Retrieves sales data from this day, formatted as YYYY-MM-DD."),
					Type:        "str",
					Required:    cohere.Bool(true),
				},
			},
		},
	}

	toolsResponse, err := client.Chat(
		context.TODO(),
		&cohere.ChatRequest{
			Model:   model,
			Message: "How good were the sales on September 29?",
			Tools:   tools,
			Preamble: cohere.String(`
				## Task Description
				You help people answer their questions and other requests interactively. You will be asked a very wide array of requests on all kinds of topics. You will be equipped with a wide range of search engines or similar tools to help you, which you use to research your answer. You should focus on serving the user's needs as best you can, which will be wide-ranging.

				## Style Guide
				Unless the user asks for a different style of answer, you should answer in full sentences, using proper grammar and spelling.
			`),
			ForceSingleStep: cohere.Bool(true),
		})

	require.NoError(t, err)
	require.NotNil(t, toolsResponse.ToolCalls)
	require.Len(t, toolsResponse.ToolCalls, 1)
	require.Equal(t, toolsResponse.ToolCalls[0].Name, "sales_database")
	require.Equal(t, toolsResponse.ToolCalls[0].Parameters["day"], "2023-09-29")

	print(toolsResponse)

	localTools := map[string]func(string) *[]map[string]interface{}{
		"sales_database": func(day string) *[]map[string]interface{} {
			return &[]map[string]interface{}{
				{
					"numberOfSales":    120,
					"totalRevenue":     48500,
					"averageSaleValue": 404.17,
					"date":             "2023-09-29",
				},
			}
		},
	}

	toolResults := make([]*cohere.ToolResult, 0)

	for _, toolCall := range toolsResponse.ToolCalls {
		result := localTools[toolCall.Name](toolCall.Parameters["day"].(string))
		toolResult := &cohere.ToolResult{
			Call:    toolCall,
			Outputs: *result,
		}
		toolResults = append(toolResults, toolResult)
	}

	citedResponse, err := client.Chat(
		context.TODO(),
		&cohere.ChatRequest{
			Model:           model,
			Message:         "How good were the sales on September 29?",
			Tools:           tools,
			ToolResults:     toolResults,
			ForceSingleStep: cohere.Bool(true),
		})

	require.NoError(t, err)

	require.Equal(t, citedResponse.Documents[0]["averageSaleValue"], "404.17")
	require.Equal(t, citedResponse.Documents[0]["date"], "2023-09-29")
	require.Equal(t, citedResponse.Documents[0]["numberOfSales"], "120")
	require.Equal(t, citedResponse.Documents[0]["totalRevenue"], "48500")
}

func TestNewAwsClient(t *testing.T) {
	t.Skip("Skip until auth is set up")
	bedrockClient := client.NewBedrockClient([]core.RequestOption{}, []client.AwsRequestOption{
		client.WithAwsRegion("us-east-1"),
		client.WithAwsAccessKey(""),
		client.WithAwsSecretKey(""),
		client.WithAwsSessionToken(""),
	})

	bedrockModels := map[string]*string{
		"generate": cohere.String("cohere.command-text-v14"),
		"embed":    cohere.String("cohere.embed-multilingual-v3"),
		"chat":     cohere.String("cohere.command-r-plus-v1:0"),
	}

	sagemakerClent := client.NewSagemakerClient([]core.RequestOption{}, []client.AwsRequestOption{
		client.WithAwsRegion("us-east-1"),
		client.WithAwsAccessKey(""),
		client.WithAwsSecretKey(""),
		client.WithAwsSessionToken(""),
	})

	sagemakerModels := map[string]*string{
		"generate": cohere.String("cohere-command-light"),
		"embed":    cohere.String("cohere-embed-multilingual-v3"),
		"chat":     cohere.String("cohere-command-plus"),
		"rerank":   cohere.String("cohere-rerank"),
	}

	t.Run("Bedrock tests", func(t *testing.T) {
		t.Run("Generate", func(t *testing.T) { Generate(t, bedrockModels["generate"], *bedrockClient) })
		t.Run("GenerateStream", func(t *testing.T) { GenerateStream(t, bedrockModels["generate"], *bedrockClient) })
		t.Run("Chat", func(t *testing.T) { Chat(t, bedrockModels["chat"], *bedrockClient) })
		t.Run("ChatStream", func(t *testing.T) { ChatStream(t, bedrockModels["chat"], *bedrockClient) })
		t.Run("Embed", func(t *testing.T) { Embed(t, bedrockModels["embed"], *bedrockClient) })
		t.Run("Tool", func(t *testing.T) { Tool(t, bedrockModels["chat"], *bedrockClient) })
	})

	t.Run("Sagemaker tests", func(t *testing.T) {
		t.Run("Generate", func(t *testing.T) { Generate(t, sagemakerModels["generate"], *sagemakerClent) })
		t.Run("GenerateStream", func(t *testing.T) { GenerateStream(t, sagemakerModels["generate"], *sagemakerClent) })
		t.Run("Chat", func(t *testing.T) { Chat(t, sagemakerModels["chat"], *sagemakerClent) })
		t.Run("ChatStream", func(t *testing.T) { ChatStream(t, sagemakerModels["chat"], *sagemakerClent) })
		t.Run("Rerank", func(t *testing.T) { Rerank(t, sagemakerModels["rerank"], *sagemakerClent) })
		t.Run("Embed", func(t *testing.T) { Embed(t, sagemakerModels["embed"], *sagemakerClent) })
		t.Run("Tool", func(t *testing.T) { Tool(t, sagemakerModels["chat"], *sagemakerClent) })
	})
}
