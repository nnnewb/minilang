package vm

import "fmt"

// Instruction 是虚拟机指令
type Instruction struct {
	OpCode  OpCode
	Operand []Object
}

func (i *Instruction) String() string {
	return fmt.Sprintf("%s %v", string(i.OpCode), i.Operand)
}
