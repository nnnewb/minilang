package compiler

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/nnnewb/minilang/pkg/ast"
	"github.com/nnnewb/minilang/pkg/bnf/lexer"
	"github.com/nnnewb/minilang/pkg/bnf/parser"
)

type Compiler struct {
	m *ir.Module
}

func NewCompiler() *Compiler {
	return &Compiler{
		m: ir.NewModule(),
	}
}

func (c *Compiler) CompileString(source string) (*ir.Module, error) {
	l := lexer.NewLexer([]byte(source))
	p := parser.NewParser()
	res, err := p.Parse(l)
	if err != nil {
		panic(err)
	}

	fn := c.m.NewFunc("main", types.I32)
	blk := fn.NewBlock("")
	err = Traversal(res.(ast.Node), func(n ast.Node) (bool, error) {
		switch node := n.(type) {
		case *ast.Define:
			// TODO:
			return false, nil
		case *ast.Combination:
			c.compileCombination(blk, node.Operator, node.Operands)
			return false, nil
		default:
			return false, nil
		}
	})
	if err != nil {
		return nil, err
	}

	blk.NewRet(constant.NewInt(types.I32, 0))
	return c.m, nil
}

func (c *Compiler) compileCombination(blk *ir.Block, operator ast.Identifier, operands []ast.Node) (value.Value, error) {
	evaluatedOperands := []value.Value{}
	for _, operand := range operands {
		evaluatedID, err := c.compileEvaluateOperand(blk, operand)
		if err != nil {
			return nil, err
		}
		evaluatedOperands = append(evaluatedOperands, evaluatedID)
	}

	switch operator {
	case "lambda":
		// TODO:
		return nil, fmt.Errorf("lambda not implemented yet")
	case "+":
		var accumulator value.Value = nil
		for _, val := range evaluatedOperands {
			if accumulator == nil {
				accumulator = val
				continue
			}
			accumulator = blk.NewFAdd(accumulator, val)
		}
		return accumulator, nil
	case "*":
		var accumulator value.Value = nil
		for _, val := range evaluatedOperands {
			if accumulator == nil {
				accumulator = val
				continue
			}
			accumulator = blk.NewFMul(accumulator, val)
		}
		return accumulator, nil
	case "/":
		var accumulator value.Value = nil
		for _, val := range evaluatedOperands {
			if accumulator == nil {
				accumulator = val
				continue
			}
			accumulator = blk.NewFDiv(accumulator, val)
		}
		return accumulator, nil
	case "-":
		if len(evaluatedOperands) == 1 {
			return blk.NewFNeg(evaluatedOperands[0]), nil
		}
		var accumulator value.Value = nil
		for _, val := range evaluatedOperands {
			if accumulator == nil {
				accumulator = val
				continue
			}
			accumulator = blk.NewFSub(accumulator, val)
		}
		return accumulator, nil
	default:
		// 其他情况下 combination 解释成函数调用
		// TODO: not implemented
		return nil, fmt.Errorf("%v is not implemented", operator)
	}
}

func (c *Compiler) compileEvaluateOperand(blk *ir.Block, operand ast.Node) (value.Value, error) {
	switch operand.NodeType() {
	case ast.NTCombination:
		comb := operand.(*ast.Combination)
		evaluated, err := c.compileCombination(blk, comb.Operator, comb.Operands)
		if err != nil {
			return nil, err
		}
		return evaluated, nil
	case ast.NTString:
		g := c.m.NewGlobalDef("", constant.NewCharArray([]byte(operand.(ast.String))))
		return g, nil
	case ast.NTBoolean, ast.NTUInt, ast.NTFloat:
		return operand.(ast.LLVMLiteral).AsLLVM(), nil
	default:
		return nil, fmt.Errorf("identifier operand not implemented")
	}
}
