package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nnnewb/minilang/internal/builtin"
	"github.com/nnnewb/minilang/internal/compiler"
	"github.com/nnnewb/minilang/internal/vm"
	"github.com/nnnewb/minilang/pkg/ast"
	"github.com/nnnewb/minilang/pkg/bnf/lexer"
	"github.com/nnnewb/minilang/pkg/bnf/parser"
)

func mustReadLine(reader *bufio.Reader) string {
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return line
}

func main() {
	m := vm.NewMiniVM([]vm.Instruction{})
	builtin.RegisterLanguageFacility(m)
	builtin.RegisterArithmetic(m)
	builtin.RegisterIO(m)

	for {
		fmt.Printf(">>> ")
		reader := bufio.NewReader(os.Stdin)
		line := strings.TrimSpace(mustReadLine(reader))
		if line == ".quit" {
			break
		}

		lexer := lexer.NewLexer([]byte(line))
		parser := parser.NewParser()
		parseResult, err := parser.Parse(lexer)
		if err != nil {
			log.Printf("parse error %v\n", err)
			continue
		}

		compiler := compiler.NewCompiler()
		instructions, err := compiler.Compile(parseResult.(ast.Node))
		if err != nil {
			log.Printf("compile error %v\n", err)
		}

		m.AddInstructions(instructions)
		for {
			err = m.ExecNextInstruction()
			if err != nil {
				if err == vm.ErrNoMoreInstructions {
					break
				} else {
					log.Printf("vm error: %v\n", err)
					break
				}
			}
		}

		evaluated := m.Pop()
		fmt.Printf("%T# %v\n", evaluated, evaluated)
	}
}
