package vm

import (
	"fmt"
	"log"
)

type NativeFunc func(m *MiniVM) error

type MiniVM struct {
	Stack        []Object          // 栈空间，包括传参和本地变量都存放在这里
	Top          int               // 栈顶地址
	Locals       map[string]Object // 本地变量，从这里查找变量和可调用的对象
	IP           UInt              // Instruction Pointer 指令指针
	Instructions []Instruction     // 程序指令集合
}

func NewMiniVM(instructions []Instruction) *MiniVM {
	return &MiniVM{
		Stack:        make([]Object, 0),
		Top:          -1,
		IP:           0,
		Instructions: instructions,
		Locals:       make(map[string]Object),
	}
}

func (m *MiniVM) LookupProc(name string) (*Procedure, error) {
	if result, ok := m.Locals[name]; ok {
		if proc, ok := result.(*Procedure); ok {
			return proc, nil
		} else {
			return nil, fmt.Errorf("%v(%T) is not procedure", result, result)
		}
	} else {
		return nil, fmt.Errorf("%s is not defined", name)
	}
}

func (m *MiniVM) Push(obj Object) {
	if m.Top+1 < len(m.Stack) {
		m.Stack[m.Top+1] = obj
	} else {
		m.Stack = append(m.Stack, obj)
	}
	m.Top++
}

func (m *MiniVM) Pop() Object {
	if m.Top < 0 {
		log.Fatalf("stack overflow: %d", m.Top)
		return nil
	}

	ret := m.Stack[m.Top]
	m.Top--
	return ret
}

func (m *MiniVM) Peek() Object {
	return m.Stack[m.Top]
}

func (m *MiniVM) AddInstructions(instructions []Instruction) int {
	loc := len(m.Instructions) - 1
	m.Instructions = append(m.Instructions, instructions...)
	return loc
}

func (m *MiniVM) instCall(inst Instruction) error {
	if len(inst.Operand) == 0 {
		return fmt.Errorf("invalid CALL instruction: %s", inst)
	}

	if sym, ok := inst.Operand[0].(Symbol); ok {
		if proc, err := m.LookupProc(string(sym)); err != nil {
			return err
		} else {
			if proc.isBuiltin {
				m.Push(UInt(m.IP + 1))
				return proc.Builtin(m)
			} else if proc.Location != 0 {
				m.Push(UInt(m.IP + 1))
				m.IP = UInt(proc.Location)
			} else {
				log.Fatalf("invalid callable object")
				return nil
			}
		}

		return nil
	}

	log.Fatalf("invalid CALL instruction operand %v", inst.Operand)
	return nil
}

func (m *MiniVM) instRet(inst Instruction) error {
	if len(inst.Operand) == 0 {
		return fmt.Errorf("invalid RET instruction: %s", inst)
	}

	returnAddress := m.Pop().(UInt)
	m.IP = returnAddress
	m.Push(inst.Operand[0])
	return nil
}

func (m *MiniVM) instPush(inst Instruction) error {
	if len(inst.Operand) == 0 || len(inst.Operand) > 1 {
		return fmt.Errorf("invalid PUSH instruction: %s", inst)
	}

	m.Push(inst.Operand[0])
	return nil
}

func (m *MiniVM) instPop(inst Instruction) error {
	if len(inst.Operand) == 0 || len(inst.Operand) > 1 {
		return fmt.Errorf("invalid POP instruction: %s", inst)
	}

	m.Pop()
	return nil
}

func (m *MiniVM) instJump(inst Instruction) error {
	if len(inst.Operand) == 0 || len(inst.Operand) > 1 {
		return fmt.Errorf("invalid JUMP instruction: %s", inst)
	}

	m.IP = inst.Operand[0].(UInt)
	return nil
}

func (m *MiniVM) instJz(inst Instruction) error {
	if len(inst.Operand) == 0 || len(inst.Operand) > 1 {
		return fmt.Errorf("invalid JZ instruction: %s", inst)
	}

	cond := m.Pop()
	if b, ok := cond.(Boolean); ok {
		if bool(b) {
			m.IP = inst.Operand[0].(UInt)
			return nil
		}
	} else {
		return fmt.Errorf("unexpected conditional value %v(%T)", cond, cond)
	}

	return nil
}

func (m *MiniVM) instLoad(inst Instruction) error {
	if len(inst.Operand) == 0 || len(inst.Operand) > 1 {
		return fmt.Errorf("invalid LOAD instruction: %s", inst)
	}

	if name, ok := inst.Operand[0].(Symbol); ok {
		if v, ok := m.Locals[string(name)]; ok {
			m.Push(v)
		} else {
			return fmt.Errorf("undefined name %s", name)
		}
	} else {
		return fmt.Errorf("unexpected operand for instruction LOAD %v(%T)", inst.Operand[0], inst.Operand[0])
	}
	return nil
}

func (m *MiniVM) ExecNextInstruction() error {
	if int(m.IP) >= len(m.Instructions) {
		return ErrNoMoreInstructions
	}

	inst := m.Instructions[m.IP]
	switch inst.OpCode {
	case CALL:
		return m.instCall(inst)
	case RET:
		m.instRet(inst)
	case PUSH:
		m.instPush(inst)
	case POP:
		m.instPop(inst)
	case JUMP:
		return m.instJump(inst)
	case JZ:
		return m.instJz(inst)
	case LOAD:
		m.instLoad(inst)
	default:
		return fmt.Errorf("unexpected opcode %v", inst.OpCode)
	}
	m.IP++
	return nil
}
