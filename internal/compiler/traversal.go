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
	case *ast.Program:
		if cont, err := callback(n); err != nil {
			return false, err
		} else if !cont {
			return false, nil
		}

		for _, d := range n.Defines {
			if _, err := visit(d, callback); err != nil {
				return false, err
			}
		}

		for _, c := range n.Combinations {
			if _, err := visit(c, callback); err != nil {
				return false, err
			}
		}

		return true, nil
	case *ast.Define:
		if cont, err := callback(n); err != nil {
			return false, err
		} else if !cont {
			return false, nil
		}

		if _, err := visit(n.Identifier, callback); err != nil {
			return false, err
		}

		if _, err := visit(n.Combination, callback); err != nil {
			return false, err
		}

		return true, nil
	case *ast.Combination:
		if cont, err := callback(n); err != nil {
			return false, err
		} else if !cont {
			return false, nil
		}

		_, err := visit(n.Operator, callback)
		if err != nil {
			return false, err
		}

		for _, operand := range n.Operands {
			if _, err := visit(operand, callback); err != nil {
				return false, err
			}
		}

		return true, err
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
