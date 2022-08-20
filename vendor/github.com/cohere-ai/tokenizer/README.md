## Tokenizers [![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
Cohere's `tokenizers` library provides an interface to encode and decode text given a computed vocabulary, and includes pre-computed tokenizers that are used to train Cohere's models. 

We plan on eventually also open sourcing tools to create new tokenizers. 

## Example using Go
Choose a tokenizer inside of the vocab folder including both a `encoder.json` file and a `vocab.bpe` file and create an encoder as seen below. The tokenizer used in this example is named the `coheretext-50k` tokenizer.
```
import (
  ...
  "github.com/cohere-ai/tokenizer"
)

encoder := tokenizer.NewFromPrebuilt("coheretext-50k")
```
    
To encode a string of text, use the Encode method. Encode returns a slice of `int64`s.
```
encoded := encoder.Encode("this is a string to be encoded")
fmt.Printf("%v", encoded)
// [6372 329 258 3852 288 345 37754]
```
To decode a slice of `int64`s, use the Decode method. Decode returns a string.
```
fmt.Printf(encoder.Decode(encoded))
// this is a string to be encoded
```

## Speed
Using a 2.5GHz CPU, encoding 1000 tokens takes approximately 6.5 milliseconds, and decoding 1000 tokens takes approximately 0.2 milliseconds.
