package compiler

import (
	"fmt"

	"github.com/nnnewb/minilang/pkg/ast"
)

func Traversal(root ast.Node, callback func(ast.Node) (bool, error)) error {
	_, err := visit(root, callback)
	return err
}

func visit(node ast.Node, callback func(ast.Node) (bool, error)) (bool, error) {
	switch n := node.(type) {
	case *ast.Combination:
		if cont, err := callback(n); err != nil {
			return false, err
		} else if cont {
			_, err := visit(n.Operator, callback)
			if err != nil {
				return false, err
			}

			for _, operand := range n.Operands {
				if _, err := visit(operand, callback); err != nil {
					return false, err
				}
			}

			return cont, err
		} else {
			return cont, err
		}
	case ast.Boolean:
		return callback(n)
	case ast.Float:
		return callback(n)
	case ast.UInt:
		return callback(n)
	case ast.String:
		return callback(n)
	default:
		panic(fmt.Errorf("unexpected ast node %v", node))
	}
}
