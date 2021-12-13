package builtin

import (
	"fmt"

	"github.com/nnnewb/minilang/internal/vm"
)

func RegisterArithmetic(i *vm.StackBasedVMInterpreter) {
	RegisterBuiltinFunc(i, "+", sum)
}

func sum(i *vm.StackBasedVMInterpreter, argc int, argv []vm.Object) (vm.Object, error) {
	accumulator := 0.0
	for _, arg := range argv {
		if ui, ok := arg.(vm.UInt); ok {
			accumulator += float64(ui)
		} else if fp, ok := arg.(vm.Float); ok {
			accumulator += float64(fp)
		} else {
			return nil, fmt.Errorf("TypeError: %v(%T) is not a number", arg, arg)
		}
	}
	return vm.Float(accumulator), nil
}
