package compiler

import "github.com/nnnewb/minilang/internal/vm"

func push(val vm.Object) vm.Instruction {
	return vm.Instruction{
		OpCode:  vm.PUSH,
		Operand: []vm.Object{val},
	}
}

func load(name string) vm.Instruction {
	return vm.Instruction{
		OpCode:  vm.LOAD,
		Operand: []vm.Object{vm.Symbol(name)},
	}
}
