package main

import (
	"encoding/json"
	"os"

	"github.com/x1bdev/go-resp/pkg/parser"
)

func main() {

	data := []byte("*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$7\r\nmyvalue\r\n")
	tokenizer := parser.NewTokenizer(data)
	instruction, err := tokenizer.Tokenize()

	if err != nil {
		panic(err)
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(instruction)
}
