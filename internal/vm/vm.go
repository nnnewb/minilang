package vm

import (
	"fmt"
)

type SymbolLookupTable struct {
	Table map[string]Object
}

type StackBasedVMInterpreter struct {
	Stack        *Stack        // 栈空间
	IP           UInt          // Instruction Pointer 指令指针
	Instructions []Instruction // 程序指令集合
	InterOp      map[string]func(interpreter *StackBasedVMInterpreter) error
	SymbolTable  SymbolLookupTable // 符号表，从这里查找变量和可调用的对象
}

func NewStackBasedVMInterpreter(instructions []Instruction) *StackBasedVMInterpreter {
	return &StackBasedVMInterpreter{
		Stack:        NewStack(),
		IP:           0,
		Instructions: instructions,
		InterOp:      make(map[string]func(interpreter *StackBasedVMInterpreter) error),
	}
}

func (interpreter *StackBasedVMInterpreter) ExecNextInstruction() error {
	if int(interpreter.IP) >= len(interpreter.Instructions) {
		return ErrNoMoreInstructions
	}

	inst := interpreter.Instructions[interpreter.IP]
	switch inst.OpCode {
	case CALL:
		interpreter.Stack.Push(UInt(interpreter.IP + 1))
		interpreter.IP = inst.Operand[0].(UInt)
		return nil
	case INTEROP:
		sym := interpreter.Stack.Pop().(Symbol)
		if callee, ok := interpreter.InterOp[string(sym)]; ok {
			if err := callee(interpreter); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("%s is not defined", sym)
		}
	case PUSH:
		for _, ope := range inst.Operand {
			interpreter.Stack.Push(ope)
		}
	case POP:
		var i uint32
		for i = 0; i < uint32(inst.Operand[0].(UInt)); i++ {
			interpreter.Stack.Pop()
		}
	case JUMP:
		interpreter.IP = inst.Operand[0].(UInt)
	case JZ:
		cond := interpreter.Stack.Pop()
		if b, ok := cond.(Boolean); ok {
			if bool(b) {
				interpreter.IP = inst.Operand[0].(UInt)
				return nil
			}
		} else {
			return fmt.Errorf("unexpected conditional value %v(%T)", cond, cond)
		}
	case LOAD:
		for _, sym := range inst.Operand {
			if name, ok := sym.(Symbol); ok {
				if v, ok := interpreter.SymbolTable.Table[string(name)]; ok {
					interpreter.Stack.Push(v)
				} else {
					return fmt.Errorf("undefined name %s", name)
				}
			} else {
				return fmt.Errorf("unexpected operand for instruction LOAD %v(%T)", sym, sym)
			}
		}
	default:
		return fmt.Errorf("unexpected opcode %v", inst.OpCode)
	}
	interpreter.IP++
	return nil
}
