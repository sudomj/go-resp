package parser

import (
	"bytes"

	"github.com/x1bdev/go-resp/pkg/types"
)

type Tokenizer struct {
	Type            byte
	tokenCategories []DataParser
}

func NewTokenizer(r []byte) *Tokenizer {

	reader := bytes.NewReader(r)

	return &Tokenizer{
		Type: r[0],
		tokenCategories: []DataParser{
			NewSimpleDataParser(reader),
			NewAggregateDataParser(reader),
		},
	}
}

func (t *Tokenizer) Tokenize() (*types.Instruction, error) {

	dataTypeCategory := t.getDataTypeCategory(t.Type)
	instruction, err := dataTypeCategory.Read()

	if err != nil {
		return nil, err
	}

	return instruction, nil
}

func (t *Tokenizer) getDataTypeCategory(char byte) DataParser {

	for _, c := range t.tokenCategories {

		if c.IsOfType(char) {
			return c
		}
	}

	return nil
}
