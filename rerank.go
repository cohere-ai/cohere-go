package cohere

type RerankOptions struct {
	// denotes the reranking model to be used
	Model string `json:"model,omitempty"`
	// denotes the search query for reranking
	Query string `json:"query"`
	// denotes the list of documents as strings or map of strings to be reranked
	Documents []interface{} `json:"documents"`
	// denotes the number of most relevant documents/indices to return
	TopN int `json:"top_n,omitempty"`
	// denotes the maximum number of chunks to produce internally per document
	MaxChunksPerDoc int `json:"max_chunks_per_doc,omitempty"`
	// denotes whether to return the provided documents
	ReturnDocuments bool `json:"return_documents,omitempty"`
}

type RerankDocument struct {
	Text string `json:"text"`
}

type RerankResult struct {
	// denotes a document object as map of string with key as `text`
	Document RerankDocument `json:"document,omitempty"`
	// denotes the relevance score assigned to the ranking
	RelevanceScore float64 `json:"relevance_score"`
	// denotes the index of the input document
	Index int `json:"index"`
}

type RerankResponse struct {
	ID      string         `json:"id"`
	Results []RerankResult `json:"results"`
}
