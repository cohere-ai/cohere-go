module github.com/cohere-ai/cohere-go

go 1.17

retract ( // retract all v1 releases, use v0 instead
	v1.2.3
	v1.2.2
	v1.2.1
	v1.2.0
	v1.1.0
	v1.0.0
)

require (
	github.com/cohere-ai/tokenizer v1.1.1
	github.com/stretchr/testify v1.8.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dlclark/regexp2 v1.4.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
