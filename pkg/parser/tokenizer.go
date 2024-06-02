package parser

import (
	"bytes"
)

type Tokenizer struct {
	buf []byte
	*bytes.Buffer
	bufIdx          int
	tokenCategories []DataTypeCategory
}

func NewTokenizer(r []byte) *Tokenizer {

	reader := bytes.NewReader(r)
	buffer := bytes.NewBuffer(r)
	return &Tokenizer{
		buf:    buffer.Bytes(),
		Buffer: buffer,
		bufIdx: 0,
		tokenCategories: []DataTypeCategory{
			NewSimpleToken(reader),
			NewAggregateToken(reader),
		},
	}
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

func (t *Token) Validate() error {

	return nil
}

type Command struct {
	Keyword string   `json:"keyword"`
	Args    []string `json:"args"`
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

	char := t.buf[t.bufIdx]
	dataTypeCategory := t.getDataTypeCategory(char)
	instruction, err := dataTypeCategory.Read()

	if err != nil {
		return nil, err
	}

	return instruction, nil
}

func (t *Tokenizer) getDataTypeCategory(char byte) DataTypeCategory {

	for _, c := range t.tokenCategories {

		if c.IsOfType(char) {
			return c
		}
	}

	return nil
}
