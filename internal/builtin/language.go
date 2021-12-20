package builtin

import (
	"fmt"

	"github.com/nnnewb/minilang/internal/compiler"
	"github.com/nnnewb/minilang/internal/vm"
)

func RegisterLanguageFacility(m *vm.MiniVM) {
	RegisterBuiltinFunc(m, "let", func(m *vm.MiniVM, argc int, argv []vm.Object) (vm.Object, error) {
		if argc != 2 {
			return nil, fmt.Errorf("let requires exact two argument")
		}

		if sym, ok := argv[0].(vm.Symbol); ok {
			if _, ok := m.Locals[string(sym)]; ok {
				return nil, fmt.Errorf("%s has already defined", sym)
			}
			m.Locals[string(sym)] = argv[1]
		} else {
			return nil, fmt.Errorf("%v(%T) is not a symbol", argv[0], argv[0])
		}

		return argv[1], nil
	})
	RegisterBuiltinFunc(m, "proc", func(m *vm.MiniVM, argc int, argv []vm.Object) (vm.Object, error) {
		if argc != 2 {
			return nil, fmt.Errorf("proc requires exact two argument")
		}

		if sym, ok := argv[0].(vm.Symbol); ok {
			if _, ok := m.Locals[string(sym)]; ok {
				return nil, fmt.Errorf("%s has already defined", sym)
			}

			c := compiler.NewCompiler()

			// 反向翻译成 AST
			reversed, err := c.ReverseCompile(argv[1])
			if err != nil {
				return nil, err
			}

			// 然后正向编译到指令合集
			instructions, err := c.Compile(m, reversed)
			if err != nil {
				return nil, err
			}

			procedure := vm.NewProcedureFromInstruction(instructions)
			m.Locals[string(sym)] = procedure
			return procedure, nil
		} else {
			return nil, fmt.Errorf("%v(%T) is not a symbol", argv[0], argv[0])
		}
	})
}
