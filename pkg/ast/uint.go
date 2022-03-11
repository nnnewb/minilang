package ast

import (
	"strconv"

	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type UInt float64

var NTUInt NodeType = "uint"

func NewUInt(text string) (UInt, error) {
	if f, err := strconv.ParseUint(text, 10, 64); err != nil {
		return 0, err
	} else {
		return UInt(f), nil
	}
}

func (UInt) NodeType() NodeType { return NTUInt }

func (u UInt) AsLLVM() value.Value {
	return constant.NewInt(types.I32, int64(u))
}
