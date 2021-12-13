package main

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/nnnewb/minilang/internal/builtin"
	"github.com/nnnewb/minilang/internal/compiler"
	"github.com/nnnewb/minilang/internal/vm"
	"github.com/nnnewb/minilang/pkg/ast"
	"github.com/nnnewb/minilang/pkg/bnf/lexer"
	"github.com/nnnewb/minilang/pkg/bnf/parser"
)

func main() {
	m := vm.NewMiniVM([]vm.Instruction{})
	builtin.RegisterArithmetic(m)
	builtin.RegisterLanguageFacility(m)

	for {
		input := prompt.Input(">>> ", func(d prompt.Document) []prompt.Suggest {
			return []prompt.Suggest{}
		})

		if input == ".quit" {
			break
		}

		lexer := lexer.NewLexer([]byte(input))
		parser := parser.NewParser()
		parseResult, err := parser.Parse(lexer)
		if err != nil {
			fmt.Printf("parse error %v\n", err)
			continue
		}

		compiler := compiler.NewCompiler()
		instructions, err := compiler.Compile(parseResult.(ast.Node))
		if err != nil {
			fmt.Printf("compile error %v\n", err)
		}

		m.AddInstructions(instructions)
		for {
			err = m.ExecNextInstruction()
			if err != nil {
				if err == vm.ErrNoMoreInstructions {
					break
				} else {
					fmt.Printf("vm error: %v\n", err)
					break
				}
			}
		}

		evaluated := m.Pop()
		fmt.Printf("%T# %v\n", evaluated, evaluated)
	}
}
