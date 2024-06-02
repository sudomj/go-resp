package buffer

import (
	"bufio"
	"io"
)

type Buffer struct {
	*bufio.Reader
}

func New(r io.Reader) *Buffer {

	return &Buffer{
		Reader: bufio.NewReader(r),
	}
}

func (b *Buffer) BulkRead() ([]byte, error) {

	buf := []byte{}
	for {

		b, err := b.ReadBytes('\n')

		if err != nil {

			if err == io.EOF {
				break
			}

			return nil, err
		}

		buf = append(buf, b...)
	}

	return buf, nil
}

func (b *Buffer) LineRead() ([]byte, error) {

	buf, _, err := b.ReadLine()

	if err != nil {

		return nil, err
	}

	return buf, nil
}
