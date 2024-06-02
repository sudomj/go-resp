package parser

import (
	"io"
	"log/slog"

	"github.com/x1bdev/go-resp/pkg/buffer"
)

const (
	NUMBER_OF_BYTES_CR = 1
	NUMBER_OF_BYTES_LF = 1
)

type AggregateDataParser struct {
	buffer *buffer.Buffer
}

func NewAggregateDataParser(r io.Reader) *AggregateDataParser {
	return &AggregateDataParser{
		buffer: buffer.New(r),
	}
}

func (a *AggregateDataParser) Read() (*Instruction, error) {

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

		length, err := a.getLength()

		if err != nil {

			if err == io.EOF {
				break
			}

			slog.Error("could not read the byte", "err", err)
			return nil, err
		}

		line, err := a.readLine()

		if err != nil {

			if err == io.EOF {
				break
			}

			slog.Error("could not read the line", "err", err)
			return nil, err
		}

		token := Token{
			Type:   string(dataType),
			Length: length,
			Data:   string(line),
		}

		command.SetKeyword(line)
		command.PushArg(string(line))
		instruction.Tokens = append(instruction.Tokens, token)
	}

	return instruction, nil
}

func (a *AggregateDataParser) readByte() (byte, error) {

	return a.buffer.ReadByte()
}

func (a *AggregateDataParser) getLength() (int, error) {

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

	a.skipLF()

	return length, nil
}

func (a *AggregateDataParser) readLine() ([]byte, error) {
	value, _, err := a.buffer.ReadLine()
	return value, err
}

func (a *AggregateDataParser) skipCRLF() {
	a.skipCR()
	a.skipLF()
}

func (a *AggregateDataParser) skipCR() {
	buf := make([]byte, NUMBER_OF_BYTES_CR)
	a.buffer.Read(buf)
}

func (a *AggregateDataParser) skipLF() {
	buf := make([]byte, NUMBER_OF_BYTES_LF)
	a.buffer.Read(buf)
}

// Todo: Maybe store the types in a map or a global variables
func (a *AggregateDataParser) IsOfType(char byte) bool {

	types := []byte{'$', '*', '!', '=', '%', '~', '>'}

	for _, t := range types {

		if t == char {
			return true
		}
	}

	return false
}
