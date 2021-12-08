package main

import (
	"github.com/c-bata/go-prompt"
	"github.com/nnnewb/scsh/pkg/bnf/lexer"
)

func main() {
	for {
		input := prompt.Input(">", func(d prompt.Document) []prompt.Suggest {
			return []prompt.Suggest{}
		})

		if input == ".quit" {
			break
		}

		lexer := lexer.NewLexer([]byte(input))

		tok := lexer.Scan()
		for tok != nil && tok.IDValue() != "" {
			println(tok.IDValue())
			tok = lexer.Scan()
		}
	}
}
