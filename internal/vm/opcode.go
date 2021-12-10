package vm

type OpCode int32

const (
	RESERVED = iota
	// CALL <PTR>
	// 压栈下一条指令的地址，跳转到目标位置
	CALL
	// LOAD <STR>
	// 读取局部环境里的变量压栈
	LOAD
	// PUSH <NUM>
	// 压栈对象
	PUSH
	// POP <NUM>
	// 出栈一定数量的对象，出栈的对象直接丢弃
	POP
	// JUMP <DEST>
	// 无条件跳转到目标位置
	JUMP
	// JZ <DEST>
	// 弹出栈顶，变量非 0 则跳转
	JZ
	// INTEROP
	// 压栈下一条指令的地址，调用运行时函数
	// TODO 考虑下 JIT 怎么转译这条指令
	INTEROP
)
