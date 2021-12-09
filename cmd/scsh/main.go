package main

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/nnnewb/minilang/internal/builtin"
	"github.com/nnnewb/minilang/internal/environment"
	"github.com/nnnewb/minilang/pkg/ast"
	"github.com/nnnewb/minilang/pkg/bnf/lexer"
	"github.com/nnnewb/minilang/pkg/bnf/parser"
)

func main() {
	for {
		input := prompt.Input(">", func(d prompt.Document) []prompt.Suggest {
			return []prompt.Suggest{}
		})

		if input == ".quit" {
			break
		}

		ee := environment.NewExecutionEnv(nil)
		builtin.RegisterArithmeticBuiltin(ee)
		ee.SetValue("display", environment.BuiltinFunc(func(ee *environment.ExecutionEnv, args []environment.Value) (environment.Value, error) {
			for _, v := range args {
				fmt.Printf("%v", v)
			}
			println()
			return nil, nil
		}))

		lexer := lexer.NewLexer([]byte(input))
		parser := parser.NewParser()
		parseResult, err := parser.Parse(lexer)
		if err != nil {
			fmt.Printf("parse error %v\n", err)
			continue
		}

		val := environment.NewValueFromASTNode(parseResult.(ast.Node))
		evaluated, err := ee.Evaluate(val)
		if err != nil {
			fmt.Printf("evaluation failed, error %v\n", err)
			continue
		}
		fmt.Printf("# (%T) %v\n", evaluated, evaluated)
	}
}
