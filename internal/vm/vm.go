package vm

import (
	"fmt"
	"log"
)

type SymbolLookupTable struct {
	Table map[string]Object
}

type StackBasedVMInterpreter struct {
	Stack        []Object      // 栈空间，包括传参和本地变量都存放在这里
	Top          int           // 栈顶地址
	IP           UInt          // Instruction Pointer 指令指针
	Instructions []Instruction // 程序指令集合
	InterOp      map[string]func(interpreter *StackBasedVMInterpreter) error
	SymbolTable  SymbolLookupTable // 符号表，从这里查找变量和可调用的对象
}

func NewStackBasedVMInterpreter(instructions []Instruction) *StackBasedVMInterpreter {
	return &StackBasedVMInterpreter{
		Stack:        make([]Object, 0),
		Top:          -1,
		IP:           0,
		Instructions: instructions,
		InterOp:      make(map[string]func(interpreter *StackBasedVMInterpreter) error),
	}
}

func (i *StackBasedVMInterpreter) Push(obj Object) {
	log.Printf("#%d vm.stack.push: %v (%T)\n", i.Top, obj, obj)
	if i.Top+1 < len(i.Stack) {
		i.Stack[i.Top+1] = obj
	} else {
		i.Stack = append(i.Stack, obj)
	}
	i.Top++
}

func (i *StackBasedVMInterpreter) Pop() Object {
	if i.Top < 0 {
		log.Fatalf("stack overflow: %d", i.Top)
		return nil
	}

	ret := i.Stack[i.Top]
	log.Printf("#%d vm.stack.pop: %v (%T)\n", i.Top, ret, ret)
	i.Top--
	return ret
}

func (i *StackBasedVMInterpreter) Peek() Object {
	return i.Stack[i.Top]
}

func (i *StackBasedVMInterpreter) ExecNextInstruction() error {
	if int(i.IP) >= len(i.Instructions) {
		return ErrNoMoreInstructions
	}

	inst := i.Instructions[i.IP]
	switch inst.OpCode {
	case CALL:
		i.Push(UInt(i.IP + 1))
		if sym, ok := inst.Operand[0].(Symbol); ok {
			if err := i.InterOp[string(sym)](i); err != nil {
				return err
			}
		}
		return nil
	case RET:
		if len(inst.Operand) == 0 {
			return fmt.Errorf("invalid return instruction")
		}
		returnAddress := i.Pop().(UInt)
		i.IP = returnAddress
		i.Push(inst.Operand[0])
		return nil
	case PUSH:
		for _, ope := range inst.Operand {
			i.Push(ope)
		}
	case POP:
		var idx uint32
		for idx = 0; idx < uint32(inst.Operand[0].(UInt)); idx++ {
			i.Pop()
		}
	case JUMP:
		i.IP = inst.Operand[0].(UInt)
	case JZ:
		cond := i.Pop()
		if b, ok := cond.(Boolean); ok {
			if bool(b) {
				i.IP = inst.Operand[0].(UInt)
				return nil
			}
		} else {
			return fmt.Errorf("unexpected conditional value %v(%T)", cond, cond)
		}
	case LOAD:
		for _, sym := range inst.Operand {
			if name, ok := sym.(Symbol); ok {
				if v, ok := i.SymbolTable.Table[string(name)]; ok {
					i.Push(v)
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
	i.IP++
	return nil
}
