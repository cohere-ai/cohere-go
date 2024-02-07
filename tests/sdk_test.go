package tests

import (
	"context"
	"errors"
	"io"
	"os"
	"strings"
	"testing"

	cohere "github.com/cohere-ai/cohere-go/v2"
	client "github.com/cohere-ai/cohere-go/v2/client"
	"github.com/stretchr/testify/require"
)

type MyReader struct {
	io.Reader
	name string
}

func (m *MyReader) Name() string {
	return m.name
}

func strPointer(s string) *string {
	return &s
}

func TestNewClient(t *testing.T) {
	client := client.NewClient(client.WithToken(os.Getenv("COHERE_API_KEY")))

	t.Run("TestGenerate", func(t *testing.T) {
		prediction, err := client.Generate(
			context.TODO(),
			&cohere.GenerateRequest{
				Prompt: "count with me!",
			},
		)

		require.NoError(t, err)
		print(prediction)
	})

	t.Run("TestGenerateStream", func(t *testing.T) {
		stream, err := client.GenerateStream(
			context.TODO(),
			&cohere.GenerateStreamRequest{
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
	})

	// Test Chat
	t.Run("TestChat", func(t *testing.T) {
		chat, err := client.Chat(
			context.TODO(),
			&cohere.ChatRequest{
				Message: "2",
			},
		)

		require.NoError(t, err)
		print(chat)
	})

	// Test ChatStream
	t.Run("TestChatStream", func(t *testing.T) {
		stream, err := client.ChatStream(
			context.TODO(),
			&cohere.ChatStreamRequest{
				Message: "Cohere is",
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
	})

	t.Run("TestClassify", func(t *testing.T) {
		classify, err := client.Classify(
			context.TODO(),
			&cohere.ClassifyRequest{
				Examples: []*cohere.ClassifyExample{
					{
						Text:  strPointer("orange"),
						Label: strPointer("fruit"),
					},
					{
						Text:  strPointer("pear"),
						Label: strPointer("fruit"),
					},
					{
						Text:  strPointer("lettuce"),
						Label: strPointer("vegetable"),
					},
					{
						Text:  strPointer("cauliflower"),
						Label: strPointer("vegetable"),
					},
				},
				Inputs: []string{"Abiu"},
			},
		)

		require.NoError(t, err)
		print(classify)
	})

	t.Run("TestTokenizeDetokenize", func(t *testing.T) {
		str := "token mctoken face"

		tokenise, err := client.Tokenize(
			context.TODO(),
			&cohere.TokenizeRequest{
				Text:  str,
				Model: strPointer("base"),
			},
		)

		require.NoError(t, err)
		print(tokenise)

		detokenise, err := client.Detokenize(
			context.TODO(),
			&cohere.DetokenizeRequest{
				Tokens: tokenise.Tokens,
			})

		require.NoError(t, err)
		print(detokenise)

		require.Equal(t, str, detokenise.Text)
	})

	t.Run("TestSummarize", func(t *testing.T) {
		summarise, err := client.Summarize(
			context.TODO(),
			&cohere.SummarizeRequest{
				Text: "the quick brown fox jumped over the lazy dog and then the dog jumped over the fox the quick brown fox jumped over the lazy dog the quick brown fox jumped over the lazy dog the quick brown fox jumped over the lazy dog the quick brown fox jumped over the lazy dog",
			})

		require.NoError(t, err)
		print(summarise)
	})

	t.Run("TestRerank", func(t *testing.T) {
		rerank, err := client.Rerank(
			context.TODO(),
			&cohere.RerankRequest{
				Query: "What is the capital of the United States?",
				Documents: []*cohere.RerankRequestDocumentsItem{
					cohere.NewRerankRequestDocumentsItemFromString("Carson City is the capital city of the American state of Nevada."),
					cohere.NewRerankRequestDocumentsItemFromString("The Commonwealth of the Northern Mariana Islands is a group of islands in the Pacific Ocean. Its capital is Saipan."),
					cohere.NewRerankRequestDocumentsItemFromString("Washington, D.C. (also known as simply Washington or D.C., and officially as the District of Columbia) is the capital of the United States. It is a federal district."),
					cohere.NewRerankRequestDocumentsItemFromString("Capital punishment (the death penalty) has existed in the United States since beforethe United States was a country. As of 2017, capital punishment is legal in 30 of the 50 states."),
				},
			})

		require.NoError(t, err)
		print(rerank)
	})

	t.Run("TestEmbed", func(t *testing.T) {
		embed, err := client.Embed(
			context.TODO(),
			&cohere.EmbedRequest{
				Texts:     []string{"hello", "goodbye"},
				Model:     strPointer("embed-english-v3.0"),
				InputType: cohere.EmbedInputTypeSearchDocument.Ptr(),
			})

		require.NoError(t, err)
		print(embed)
	})

	t.Run("TestCreateDataset", func(t *testing.T) {
		t.Skip("While we have issues with dataset upload")

		dataset, err := client.Datasets.Create(
			context.TODO(),
			&MyReader{Reader: strings.NewReader(`{"text": "The quick brown fox jumps over the lazy dog"}`), name: "test.jsonl"},
			&MyReader{Reader: strings.NewReader(""), name: "a.jsonl"},
			&cohere.DatasetsCreateRequest{
				Name: strPointer("prompt-completion-dataset"),
				Type: cohere.DatasetTypeEmbedResult.Ptr(),
			})

		require.NoError(t, err)
		print(dataset)
	})

	t.Run("TestListDatasets", func(t *testing.T) {
		datasets, err := client.Datasets.List(
			context.TODO(),
			&cohere.DatasetsListRequest{})

		require.NoError(t, err)
		print(datasets)
	})

	t.Run("TestGetDatasetUsage", func(t *testing.T) {
		t.Skip("While we have issues with dataset upload")
		dataset_usage, err := client.Datasets.GetUsage(context.TODO())

		require.NoError(t, err)
		print(dataset_usage)
	})

	t.Run("TestGetDataset", func(t *testing.T) {
		t.Skip("While we have issues with dataset upload")
		dataset, err := client.Datasets.Get(context.TODO(), "id")

		require.NoError(t, err)
		print(dataset)
	})

	t.Run("TestUpdateDataset", func(t *testing.T) {
		t.Skip("While we have issues with dataset upload")
		_, err := client.Datasets.Delete(context.TODO(), "id")
		require.NoError(t, err)
	})

	t.Run("TestCreateEmbedJob", func(t *testing.T) {
		t.Skip("While we have issues with dataset upload")
		job, err := client.EmbedJobs.Create(
			context.TODO(),
			&cohere.CreateEmbedJobRequest{
				DatasetId: "id",
				InputType: cohere.EmbedInputTypeSearchDocument,
			})

		require.NoError(t, err)
		print(job)
	})

	t.Run("TestListEmbedJobs", func(t *testing.T) {
		embed_jobs, err := client.EmbedJobs.List(context.TODO())

		require.NoError(t, err)
		print(embed_jobs)
	})

	t.Run("TestGetEmbedJob", func(t *testing.T) {
		t.Skip("While we have issues with dataset upload")
		embed_job, err := client.EmbedJobs.Get(context.TODO(), "id")

		require.NoError(t, err)
		print(embed_job)
	})

	t.Run("TestCancelEmbedJob", func(t *testing.T) {
		t.Skip("While we have issues with dataset upload")
		err := client.EmbedJobs.Cancel(context.TODO(), "id")

		require.NoError(t, err)
	})

	t.Run("TestConnectorCRUD", func(t *testing.T) {
		connector, err := client.Connectors.Create(
			context.TODO(),
			&cohere.CreateConnectorRequest{
				Name: "Example connector",
				Url:  "https://dummy-connector-o5btz7ucgq-uc.a.run.app/search",
				ServiceAuth: &cohere.CreateConnectorServiceAuth{
					Token: "dummy-connector-token",
					Type:  "bearer",
				},
			})

		require.NoError(t, err)
		print(connector)

		updated_connector, err := client.Connectors.Update(
			context.TODO(),
			connector.Connector.Id,
			&cohere.UpdateConnectorRequest{
				Name: strPointer("Example connector renamed"),
			})

		require.NoError(t, err)
		print(updated_connector)

		my_connector, err := client.Connectors.Get(context.TODO(), connector.Connector.Id)

		require.NoError(t, err)
		print(my_connector)

		connectors, err := client.Connectors.List(
			context.TODO(),
			&cohere.ConnectorsListRequest{})

		require.NoError(t, err)
		print(connectors)

		oauth, err := client.Connectors.OAuthAuthorize(
			context.TODO(),
			connector.Connector.Id,
			&cohere.ConnectorsOAuthAuthorizeRequest{
				AfterTokenRedirect: strPointer("https://test.com"),
			})

		// find a way to test this
		require.Error(t, err)
		print(oauth)

		delete, err := client.Connectors.Delete(context.TODO(), connector.Connector.Id)

		require.NoError(t, err)
		print(delete)
	})
}
