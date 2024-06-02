package parser

type DataTypeCategory interface {
	IsOfType(char byte) bool
	Read() (*Instruction, error)
}
