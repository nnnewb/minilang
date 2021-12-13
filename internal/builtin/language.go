package builtin

import (
	"fmt"

	"github.com/nnnewb/minilang/internal/vm"
)

func RegisterLanguageFacility(m *vm.MiniVM) {
	RegisterBuiltinFunc(m, "let", func(m *vm.MiniVM, argc int, argv []vm.Object) (vm.Object, error) {
		if argc > 2 {
			return nil, fmt.Errorf("too many argument")
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
}
