// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/cohere-ai/cohere-go/v2/core"
	time "time"
)

type DatasetsCreateRequest struct {
	// The name of the uploaded dataset.
	Name *string `json:"-" url:"name,omitempty"`
	// The dataset type, which is used to validate the data.
	Type *DatasetType `json:"-" url:"type,omitempty"`
	// Indicates if the original file should be stored.
	KeepOriginalFile *bool `json:"-" url:"keep_original_file,omitempty"`
	// Indicates whether rows with malformed input should be dropped (instead of failing the validation check). Dropped rows will be returned in the warnings field.
	SkipMalformedInput *bool `json:"-" url:"skip_malformed_input,omitempty"`
	// List of names of fields that will be persisted in the Dataset. By default the Dataset will retain only the required fields indicated in the [schema for the corresponding Dataset type](https://docs.cohere.com/docs/datasets#dataset-types). For example, datasets of type `embed-input` will drop all fields other than the required `text` field. If any of the fields in `keep_fields` are missing from the uploaded file, Dataset validation will fail.
	KeepFields []*string `json:"-" url:"keep_fields,omitempty"`
	// List of names of fields that will be persisted in the Dataset. By default the Dataset will retain only the required fields indicated in the [schema for the corresponding Dataset type](https://docs.cohere.com/docs/datasets#dataset-types). For example, Datasets of type `embed-input` will drop all fields other than the required `text` field. If any of the fields in `optional_fields` are missing from the uploaded file, Dataset validation will pass.
	OptionalFields []*string `json:"-" url:"optional_fields,omitempty"`
	// Raw .txt uploads will be split into entries using the text_separator value.
	TextSeparator *string `json:"-" url:"text_separator,omitempty"`
	// The delimiter used for .csv uploads.
	CsvDelimiter *string `json:"-" url:"csv_delimiter,omitempty"`
}

type DatasetsListRequest struct {
	// optional filter by dataset type
	DatasetType *string `json:"-" url:"datasetType,omitempty"`
	// optional filter before a date
	Before *time.Time `json:"-" url:"before,omitempty"`
	// optional filter after a date
	After *time.Time `json:"-" url:"after,omitempty"`
	// optional limit to number of results
	Limit *string `json:"-" url:"limit,omitempty"`
	// optional offset to start of results
	Offset *string `json:"-" url:"offset,omitempty"`
}

type DatasetsCreateResponse struct {
	// The dataset ID
	Id *string `json:"id,omitempty" url:"id,omitempty"`

	_rawJSON json.RawMessage
}

func (d *DatasetsCreateResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DatasetsCreateResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DatasetsCreateResponse(value)
	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DatasetsCreateResponse) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type DatasetsGetResponse struct {
	Dataset *Dataset `json:"dataset,omitempty" url:"dataset,omitempty"`

	_rawJSON json.RawMessage
}

func (d *DatasetsGetResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DatasetsGetResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DatasetsGetResponse(value)
	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DatasetsGetResponse) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type DatasetsGetUsageResponse struct {
	// The total number of bytes used by the organization.
	OrganizationUsage *string `json:"organization_usage,omitempty" url:"organization_usage,omitempty"`

	_rawJSON json.RawMessage
}

func (d *DatasetsGetUsageResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DatasetsGetUsageResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DatasetsGetUsageResponse(value)
	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DatasetsGetUsageResponse) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type DatasetsListResponse struct {
	Datasets []*Dataset `json:"datasets,omitempty" url:"datasets,omitempty"`

	_rawJSON json.RawMessage
}

func (d *DatasetsListResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DatasetsListResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DatasetsListResponse(value)
	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DatasetsListResponse) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}
