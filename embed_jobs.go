// This file was auto-generated by Fern from our API Definition.

package api

import (
	fmt "fmt"
)

type CreateEmbedJobRequest struct {
	// ID of the embedding model.
	//
	// Available models and corresponding embedding dimensions:
	//
	// - `embed-english-v3.0` : 1024
	// - `embed-multilingual-v3.0` : 1024
	// - `embed-english-light-v3.0` : 384
	// - `embed-multilingual-light-v3.0` : 384
	Model string `json:"model" url:"-"`
	// ID of a [Dataset](https://docs.cohere.com/docs/datasets). The Dataset must be of type `embed-input` and must have a validation status `Validated`
	DatasetId string         `json:"dataset_id" url:"-"`
	InputType EmbedInputType `json:"input_type" url:"-"`
	// The name of the embed job.
	Name *string `json:"name,omitempty" url:"-"`
	// Specifies the types of embeddings you want to get back. Not required and default is None, which returns the Embed Floats response type. Can be one or more of the following types.
	//
	// * `"float"`: Use this when you want to get back the default float embeddings. Valid for all models.
	// * `"int8"`: Use this when you want to get back signed int8 embeddings. Valid for only v3 models.
	// * `"uint8"`: Use this when you want to get back unsigned int8 embeddings. Valid for only v3 models.
	// * `"binary"`: Use this when you want to get back signed binary embeddings. Valid for only v3 models.
	// * `"ubinary"`: Use this when you want to get back unsigned binary embeddings. Valid for only v3 models.
	EmbeddingTypes []EmbeddingType `json:"embedding_types,omitempty" url:"-"`
	// One of `START|END` to specify how the API will handle inputs longer than the maximum token length.
	//
	// Passing `START` will discard the start of the input. `END` will discard the end of the input. In both cases, input is discarded until the remaining input is exactly the maximum input token length for the model.
	Truncate *CreateEmbedJobRequestTruncate `json:"truncate,omitempty" url:"-"`
}

// One of `START|END` to specify how the API will handle inputs longer than the maximum token length.
//
// Passing `START` will discard the start of the input. `END` will discard the end of the input. In both cases, input is discarded until the remaining input is exactly the maximum input token length for the model.
type CreateEmbedJobRequestTruncate string

const (
	CreateEmbedJobRequestTruncateStart CreateEmbedJobRequestTruncate = "START"
	CreateEmbedJobRequestTruncateEnd   CreateEmbedJobRequestTruncate = "END"
)

func NewCreateEmbedJobRequestTruncateFromString(s string) (CreateEmbedJobRequestTruncate, error) {
	switch s {
	case "START":
		return CreateEmbedJobRequestTruncateStart, nil
	case "END":
		return CreateEmbedJobRequestTruncateEnd, nil
	}
	var t CreateEmbedJobRequestTruncate
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c CreateEmbedJobRequestTruncate) Ptr() *CreateEmbedJobRequestTruncate {
	return &c
}
