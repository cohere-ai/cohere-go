// Code generated by Fern. DO NOT EDIT.

package api

import (
	finetuning "github.com/cohere-ai/cohere-go/v2/finetuning"
)

type FinetuningListEventsRequest struct {
	// Maximum number of results to be returned by the server. If 0, defaults to
	// 50.
	PageSize *int `json:"-" url:"page_size,omitempty"`
	// Request a specific page of the list results.
	PageToken *string `json:"-" url:"page_token,omitempty"`
	// Comma separated list of fields. For example: "created_at,name". The default
	// sorting order is ascending. To specify descending order for a field, append
	// " desc" to the field name. For example: "created_at desc,name".
	//
	// Supported sorting fields:
	//   - created_at (default)
	OrderBy *string `json:"-" url:"order_by,omitempty"`
}

type FinetuningListFinetunedModelsRequest struct {
	// Maximum number of results to be returned by the server. If 0, defaults to
	// 50.
	PageSize *int `json:"-" url:"page_size,omitempty"`
	// Request a specific page of the list results.
	PageToken *string `json:"-" url:"page_token,omitempty"`
	// Comma separated list of fields. For example: "created_at,name". The default
	// sorting order is ascending. To specify descending order for a field, append
	// " desc" to the field name. For example: "created_at desc,name".
	//
	// Supported sorting fields:
	//   - created_at (default)
	OrderBy *string `json:"-" url:"order_by,omitempty"`
}

type FinetuningListTrainingStepMetricsRequest struct {
	// Maximum number of results to be returned by the server. If 0, defaults to
	// 50.
	PageSize *int `json:"-" url:"page_size,omitempty"`
	// Request a specific page of the list results.
	PageToken *string `json:"-" url:"page_token,omitempty"`
}

type FinetuningUpdateFinetunedModelRequest struct {
	// FinetunedModel name (e.g. `foobar`).
	Name string `json:"name" url:"-"`
	// FinetunedModel settings such as dataset, hyperparameters...
	Settings *finetuning.Settings `json:"settings,omitempty" url:"-"`
	// Current stage in the life-cycle of the fine-tuned model.
	Status *finetuning.Status `json:"status,omitempty" url:"-"`
}
