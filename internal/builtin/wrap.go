package builtin

import (
	"github.com/nnnewb/minilang/internal/vm"
)

func extractArgs(i *vm.StackBasedVMInterpreter) (int, []vm.Object) {
	stackFrameBase := i.Top
	argc := i.Stack[stackFrameBase-1].(vm.UInt)
	argv := make([]vm.Object, 0, argc)
	for idx := 1; idx <= int(argc); idx++ {
		argv = append(argv, i.Stack[stackFrameBase-1-idx])
	}
	return int(argc), argv
}

func cleanupStack(i *vm.StackBasedVMInterpreter) vm.UInt {
	// 取返回地址
	returnLoc := i.Pop().(vm.UInt)
	// 弹出所有参数和返回地址完成平栈
	argc := i.Pop().(vm.UInt)
	for idx := 0; idx < int(argc); idx++ {
		i.Pop()
	}
	return returnLoc
}

type BuiltinFunc = func(i *vm.StackBasedVMInterpreter, argc int, argv []vm.Object) (vm.Object, error)

func builtinWrapper(fn BuiltinFunc) func(*vm.StackBasedVMInterpreter) error {
	return func(i *vm.StackBasedVMInterpreter) error {
		argc, argv := extractArgs(i)
		if ret, err := fn(i, argc, argv); err != nil {
			return err
		} else {
			loc := cleanupStack(i)
			i.Push(ret)
			i.IP = loc
			return nil
		}
	}
}

func RegisterBuiltinFunc(i *vm.StackBasedVMInterpreter, ident string, fun BuiltinFunc) {
	i.InterOp[ident] = builtinWrapper(fun)
}
