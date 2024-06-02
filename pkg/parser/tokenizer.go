package parser

import (
	"bytes"
)

type Tokenizer struct {
	Type            byte
	tokenCategories []DataParser
}

type Instruction struct {
	Type    string   `json:"type"`
	Data    string   `json:"data"`
	Tokens  []Token  `json:"tokens"`
	Command *Command `json:"command"`
}

type Token struct {
	Type   string `json:"type"`
	Length int    `json:"length"`
	Data   string `json:"data"`
}

type Command struct {
	Keyword string   `json:"keyword"`
	Args    []string `json:"args"`
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

func (t *Token) Validate() error {

	return nil
}

func (c *Command) SetKeyword(keyword []byte) {

	if c.Keyword == "" {
		c.Keyword = string(keyword)
	}
}

func (c *Command) PushArg(arg string) {

	if arg == c.Keyword {
		return
	}

	c.Args = append(c.Args, arg)
}

func (t *Tokenizer) Tokenize() (*Instruction, error) {

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
