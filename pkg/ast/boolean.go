package ast

import (
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/value"
)

type Boolean bool

var NTBoolean NodeType = "boolean"

func (Boolean) NodeType() NodeType { return NTBoolean }

func (b Boolean) AsLLVM() value.Value {
	return constant.NewBool(bool(b))
}
