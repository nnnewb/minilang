package vm

type OpCode int32

const (
	RESERVED = iota
	// CALL <STR>
	// 压栈下一条指令的地址，跳转到指定位置
	CALL
	// RET <OBJ>
	// 取栈顶的对象作为跳转地址，压栈返回值
	RET
	// LOAD <STR>
	// 读取局部环境里的变量压栈
	LOAD
	// PUSH <NUM>
	// 压栈对象
	PUSH
	// POP <NUM>
	// 出栈一定数量的对象，出栈的对象直接丢弃
	POP
)

func (o OpCode) String() string {
	return []string{
		"RESERVED",
		"CALL",
		"RET",
		"LOAD",
		"PUSH",
		"POP",
	}[o]
}
