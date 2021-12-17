package compiler

import (
	"fmt"

	"github.com/nnnewb/minilang/internal/vm"
	"github.com/nnnewb/minilang/pkg/ast"
)

type Compiler struct {
}

func NewCompiler() *Compiler {
	return &Compiler{}
}

func (c *Compiler) Compile(root ast.Node) ([]vm.Instruction, error) {
	instructions := make([]vm.Instruction, 0)

	switch node := root.(type) {
	case *ast.List:
		elements := node.GetValue().([]ast.Node)
		if len(elements) == 0 {
			return make([]vm.Instruction, 0), nil
		}

		// 参数从右到左压栈
		for i := len(elements) - 1; i > 0; i-- {
			if inst, err := c.Compile(elements[i]); err != nil {
				return nil, err
			} else {
				instructions = append(instructions, inst...)
			}
		}

		// 压入参数数量
		instructions = append(instructions, push(vm.UInt(len(elements)-1)))

		// 插入调用语句
		callee := elements[0]
		if ident, ok := callee.(ast.Identifier); ok {
			instructions = append(instructions, vm.Instruction{
				OpCode:  vm.CALL,
				Operand: []vm.Object{vm.Symbol(ident)},
			})
		}
	case ast.Boolean, ast.UInt, ast.Float, ast.Symbol, *ast.Quoted, ast.String:
		instructions = append(instructions, push(vm.ObjectFromLiteral(node)))
	case ast.Identifier:
		instructions = append(instructions, load(string(node)))
	default:
		return nil, fmt.Errorf("unexpected ast node %v(%T)", node, node)
	}
	return instructions, nil
}

func (c *Compiler) ReverseCompile(obj vm.Object) (ast.Node, error) {
	switch v := obj.(type) {
	case vm.UInt:
		return ast.UInt(v), nil
	case vm.Float:
		return ast.Float(v), nil
	case vm.Boolean:
		return ast.Boolean(v), nil
	case vm.String:
		return ast.String(v), nil
	case vm.Symbol:
		return ast.Symbol(v), nil
	case *vm.List:
		root := ast.NewList()
		for _, elem := range v.GetUnderlyingList() {
			reversed, err := c.ReverseCompile(elem)
			if err != nil {
				return nil, err
			}
			root.Append(reversed)
		}
		return root, nil
	default:
		return nil, fmt.Errorf("%v(%T) cannot reversed", obj, obj)
	}
}
