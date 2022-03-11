package compiler_test

import (
	"testing"

	"github.com/nnnewb/minilang/internal/compiler"
)

func TestCompile(t *testing.T) {
	compiler := compiler.NewCompiler()
	_, err := compiler.CompileString(`(add 1 2)`)
	if err != nil {
		t.Error(err)
	}
}
