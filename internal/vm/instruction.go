package vm

// Instruction 是虚拟机指令
type Instruction struct {
	OpCode  OpCode
	Operand []Object
}
