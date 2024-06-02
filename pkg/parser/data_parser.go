package parser

import "github.com/x1bdev/go-resp/pkg/types"

type DataParser interface {
	IsOfType(char byte) bool
	Read() (*types.Instruction, error)
}
