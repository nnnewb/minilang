package ast

import "github.com/llir/llvm/ir/value"

type LLVMLiteral interface {
	AsLLVM() value.Value
}
