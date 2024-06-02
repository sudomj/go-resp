package parser

type DataParser interface {
	IsOfType(char byte) bool
	Read() (*Instruction, error)
}
