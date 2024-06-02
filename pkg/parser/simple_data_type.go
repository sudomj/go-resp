package parser

import (
	"io"

	"github.com/x1bdev/go-resp/pkg/buffer"
)

type SimpleDataType struct {
	buffer *buffer.Buffer
}

func NewSimpleToken(r io.Reader) *SimpleDataType {
	return &SimpleDataType{
		buffer: buffer.New(r),
	}
}

func (s *SimpleDataType) Read() (*Instruction, error) {

	s.buffer.LineRead()
	return nil, nil
}

// Todo: Maybe store the types in a map or a global variables
func (s *SimpleDataType) IsOfType(char byte) bool {

	types := []byte{'+', '-', ':', '_', '#', ',', '('}

	for _, t := range types {
		if t == char {
			return true
		}
	}

	return false
}
