package parser

import (
	"io"
	"log/slog"

	"github.com/x1bdev/go-resp/pkg/buffer"
)

type AggregateDataType struct {
	buffer *buffer.Buffer
}

func NewAggregateToken(r io.Reader) *AggregateDataType {
	return &AggregateDataType{
		buffer: buffer.New(r),
	}
}

func (a *AggregateDataType) Read() (*Instruction, error) {

	commandType, err := a.readByte()

	if err != nil {

		return nil, err
	}

	numberOfElements, err := a.readByte()

	if err != nil {

		return nil, err
	}

	a.skipCRLF()

	command := &Command{Args: make([]string, 0)}
	instruction := &Instruction{
		Type:    string(commandType),
		Data:    string(numberOfElements),
		Tokens:  []Token{},
		Command: command,
	}

	for {

		dataType, err := a.readByte()

		if err != nil {
			if err == io.EOF {
				break
			}

			slog.Error("could not read the byte", "err", err)
			return nil, err
		}

		slog.Info("reading byte", "data type", string(dataType))

		length, err := a.getLength()

		if err != nil {
			slog.Error("could not read the byte", "err", err)
			return nil, err
		}

		a.readByte()

		line, err := a.readLine()

		if err != nil {
			slog.Error("could not read the line", "err", err)
			return nil, err
		}

		token := Token{
			Type:   string(dataType),
			Length: length,
			Data:   string(line),
		}

		command.SetKeyword(line)
		instruction.Tokens = append(instruction.Tokens, token)
		command.PushArg(string(line))
	}

	return instruction, nil
}

func (a *AggregateDataType) readByte() (byte, error) {

	return a.buffer.ReadByte()
}

func (a *AggregateDataType) getLength() (int, error) {

	length := 0
	index := 0

	for {
		b, err := a.readByte()

		if b == '\r' {
			break
		}

		if err != nil {
			return -1, err
		}

		asciValue := int(b)
		asciZero := int('0')
		length = (index * 10) + (asciValue - asciZero)
		index += 1
	}

	return length, nil
}

func (a *AggregateDataType) readLine() ([]byte, error) {
	value, _, err := a.buffer.ReadLine()
	return value, err
}

func (a *AggregateDataType) skipCRLF() {

	escape := make([]byte, 2)
	a.buffer.Read(escape)
}

// Todo: Maybe store the types in a map or a global variables
func (a *AggregateDataType) IsOfType(char byte) bool {

	types := []byte{'$', '*', '!', '=', '%', '~', '>'}

	for _, t := range types {

		if t == char {
			return true
		}
	}

	return false
}
