package ast

import (
	"strconv"

	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

var NTFloat NodeType = "float"

type Float float64

func NewFloat(text string) (Float, error) {
	if f, err := strconv.ParseFloat(text, 64); err != nil {
		return 0, err
	} else {
		return Float(f), nil
	}
}

func (Float) NodeType() NodeType { return NTFloat }

func (f Float) AsLLVM() value.Value {
	return constant.NewFloat(types.Float, float64(f))
}
