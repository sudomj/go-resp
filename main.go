package main

import (
	"encoding/json"
	"os"

	"github.com/x1bdev/go-resp/pkg/parser"
)

func main() {

	data := []byte("*6\r\n$4\r\nHSET\r\n$8\r\nuser:123\r\n$4\r\nname\r\n$4\r\nJohn\r\n$3\r\nage\r\n$2\r\n30\r\n$5\r\nemail\r\n$16\r\njohn@example.com\r\n")
	tokenizer := parser.NewTokenizer(data)
	instruction, err := tokenizer.Tokenize()

	if err != nil {
		panic(err)
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(instruction)
}
