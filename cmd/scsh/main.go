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
		inst, err := compiler.Compile(parseResult.(ast.Node))
		if err != nil {
			fmt.Printf("compile error %v\n", err)
		}

		executor := vm.NewStackBasedVMInterpreter(inst)
		builtin.RegisterArithmetic(executor)
		for {
			if int(executor.IP) < len(executor.Instructions) {
				fmt.Printf("IP[%d] -> %s\n", executor.IP, executor.Instructions[executor.IP])
			}

			err = executor.ExecNextInstruction()
			if err != nil {
				if err == vm.ErrNoMoreInstructions {
					break
				} else {
					fmt.Printf("vm error: %v\n", err)
					break
				}
			}
		}

		evaluated := executor.Pop()
		fmt.Printf("%T# %v\n", evaluated, evaluated)
	}
}
