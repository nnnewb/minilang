package main

import (
	"log"
	"os"

	"github.com/nnnewb/minilang/internal/compiler"
)

func main() {
	if len(os.Args) <= 1 {
		log.Printf("Usage: scc <source file>")
		return
	}
	for _, arg := range os.Args[1:] {
		content, err := os.ReadFile(arg)
		if err != nil {
			panic(err)
		}

		compiler := compiler.NewCompiler()
		module, err := compiler.CompileString(string(content))
		if err != nil {
			panic(err)
		}

		os.WriteFile(arg+".ll", []byte(module.String()), 0o644)
	}
}
