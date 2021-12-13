package builtin

import (
	"github.com/nnnewb/minilang/internal/vm"
)

func extractArgs(m *vm.MiniVM) (int, []vm.Object) {
	stackFrameBase := m.Top
	argc := m.Stack[stackFrameBase-1].(vm.UInt)
	argv := make([]vm.Object, 0, argc)
	for idx := 1; idx <= int(argc); idx++ {
		argv = append(argv, m.Stack[stackFrameBase-1-idx])
	}
	return int(argc), argv
}

func cleanupStack(m *vm.MiniVM) vm.UInt {
	// 取返回地址
	returnLoc := m.Pop().(vm.UInt)
	// 弹出所有参数和返回地址完成平栈
	argc := m.Pop().(vm.UInt)
	for idx := 0; idx < int(argc); idx++ {
		m.Pop()
	}
	return returnLoc
}

type builtinFunc = func(m *vm.MiniVM, argc int, argv []vm.Object) (vm.Object, error)

func builtinWrapper(fn builtinFunc) func(*vm.MiniVM) error {
	return func(m *vm.MiniVM) error {
		var (
			argc, argv = extractArgs(m)
			ret, err   = fn(m, argc, argv)
		)
		defer func() {
			loc := cleanupStack(m)
			m.IP = loc
			if m != nil {
				m.Push(ret)
			} else {
				m.Push(vm.Nil{})
			}
		}()
		return err
	}
}

func RegisterBuiltinFunc(m *vm.MiniVM, ident string, fun builtinFunc) {
	m.Locals[ident] = vm.NewProcedureFromNativeFunc(builtinWrapper(fun))
}
