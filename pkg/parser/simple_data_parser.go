package parser

import (
	"io"
	"log/slog"

	"github.com/x1bdev/go-resp/pkg/buffer"
)

type SimpleDataParser struct {
	buffer *buffer.Buffer
}

func NewSimpleDataParser(r io.Reader) *SimpleDataParser {
	return &SimpleDataParser{
		buffer: buffer.New(r),
	}
}

func (s *SimpleDataParser) Read() (*Instruction, error) {

	commandType, err := s.readByte()

	if err != nil {
		slog.Error("could not read the byte", "err", err)
		return nil, err
	}

	command := &Command{}
	instruction := &Instruction{
		Type:    string(commandType),
		Command: command,
	}

	line, err := s.readLine()

	if err != nil {
		slog.Error("could not read the line", "err", err)
		return nil, err
	}

	token := Token{
		Type:   string(commandType),
		Length: len(line),
		Data:   string(line),
	}
	instruction.Tokens = append(instruction.Tokens, token)
	command.PushArg(string(line))

	return instruction, nil
}

func (s *SimpleDataParser) readByte() (byte, error) {

	return s.buffer.ReadByte()
}

func (s *SimpleDataParser) readLine() ([]byte, error) {
	line, _, err := s.buffer.ReadLine()
	return line, err
}

// Todo: Maybe store the types in a map or a global variables
func (s *SimpleDataParser) IsOfType(char byte) bool {

	types := []byte{'+', '-', ':', '_', '#', ',', '('}

	for _, t := range types {
		if t == char {
			return true
		}
	}

	return false
}
