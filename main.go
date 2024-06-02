package main

import (
	"encoding/json"
	"os"

	"github.com/x1bdev/go-resp/pkg/parser"
)

func main() {

	data := []byte("*4\r\n$4\r\nHSET\r\n$5\r\nusers\r\n$5\r\n12345\r\n$8\r\n{\"name\":\"John\",\"age\":30}\r\n")
	tokenizer := parser.NewTokenizer(data)
	instruction, err := tokenizer.Tokenize()

	if err != nil {
		panic(err)
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(instruction)
}
