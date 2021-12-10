package main

import (
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/nnnewb/minilang/internal/ilcompiler"
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

		compiler := ilcompiler.NewCompiler()
		inst, err := compiler.Compile(parseResult.(ast.Node))
		if err != nil {
			fmt.Printf("compile error %v\n", err)
		}

		if len(inst) > 0 {
			executor := vm.NewStackBasedVMInterpreter(inst)
			executor.InterOp["display"] = func(interpreter *vm.StackBasedVMInterpreter) error {
				sb := strings.Builder{}
				argc := uint32(interpreter.Stack.Pop().(vm.UInt))
				var i uint32
				for i = 0; i < argc; i++ {
					arg := interpreter.Stack.Pop()
					sb.WriteString(fmt.Sprintf("%v", arg))
					if i+1 != argc {
						sb.WriteString(" ")
					}
				}
				println(sb.String())
				interpreter.Stack.Push(vm.UInt(0))
				return nil
			}
			for {
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
		}
	}
}
