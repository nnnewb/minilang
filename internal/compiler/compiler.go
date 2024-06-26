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
	currentModule *ir.Module
	currentFunc   *ir.Func
	currentBlock  *ir.Block
}

func NewCompiler() *Compiler {
	return &Compiler{
		currentModule: ir.NewModule(),
	}
}

func (c *Compiler) addDeclarations() error {
	printf := c.currentModule.NewFunc("printf", types.I32, ir.NewParam("", types.I8Ptr))
	printf.Sig.Variadic = true
	return nil
}

func (c *Compiler) CompileString(source string) (*ir.Module, error) {
	l := lexer.NewLexer([]byte(source))
	p := parser.NewParser()
	res, err := p.Parse(l)
	if err != nil {
		panic(err)
	}

	err = Traversal(res.(ast.Node), func(n ast.Node) (bool, error) {
		switch node := n.(type) {
		case *ast.Program:
			c.currentModule = ir.NewModule()
			c.currentFunc = c.currentModule.NewFunc("main", types.I32)
			c.currentBlock = c.currentFunc.NewBlock("")
			c.addDeclarations()
			return true, nil
		case *ast.Define:
			// TODO:
			return false, nil
		case *ast.Combination:
			_, err = c.compileCombination(node.Operator, node.Operands)
			if err != nil {
				return false, err
			}
			return false, nil
		default:
			return false, nil
		}
	})
	if err != nil {
		return nil, err
	}

	c.currentBlock.NewRet(constant.NewInt(types.I32, 0))
	return c.currentModule, nil
}

func (c *Compiler) compileCombination(operator ast.Identifier, operands []ast.Node) (value.Value, error) {
	evaluatedOperands := []value.Value{}
	for _, operand := range operands {
		evaluatedID, err := c.compileEvaluateOperand(operand)
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
			accumulator = c.currentBlock.NewAdd(accumulator, val)
		}
		return accumulator, nil
	case "*":
		var accumulator value.Value = nil
		for _, val := range evaluatedOperands {
			if accumulator == nil {
				accumulator = val
				continue
			}
			accumulator = c.currentBlock.NewMul(accumulator, val)
		}
		return accumulator, nil
	case "/":
		var accumulator value.Value = nil
		for _, val := range evaluatedOperands {
			if accumulator == nil {
				accumulator = val
				continue
			}
			accumulator = c.currentBlock.NewSDiv(accumulator, val)
		}
		return accumulator, nil
	case "-":
		if len(evaluatedOperands) == 1 {
			return c.currentBlock.NewFNeg(evaluatedOperands[0]), nil
		}
		var accumulator value.Value = nil
		for _, val := range evaluatedOperands {
			if accumulator == nil {
				accumulator = val
				continue
			}
			accumulator = c.currentBlock.NewSub(accumulator, val)
		}
		return accumulator, nil
	default:
		// 其他情况下 combination 解释成函数调用
		for _, f := range c.currentModule.Funcs {
			if f.Name() == string(operator) {
				return c.compileFunctionCall(f, evaluatedOperands...)
			}
		}
		return nil, fmt.Errorf("function %s not defined", operator)
	}
}

func (c *Compiler) compileEvaluateOperand(operand ast.Node) (value.Value, error) {
	switch operand.NodeType() {
	case ast.NTCombination:
		comb := operand.(*ast.Combination)
		evaluated, err := c.compileCombination(comb.Operator, comb.Operands)
		if err != nil {
			return nil, err
		}
		return evaluated, nil
	case ast.NTString:
		s := string(operand.(ast.String))
		g := c.currentModule.NewGlobalDef("", constant.NewCharArrayFromString(s[1:len(s)-1]))
		return g, nil
	case ast.NTUInt:
		return (operand.(ast.UInt).AsLLVM()), nil
	default:
		return nil, fmt.Errorf("identifier operand not implemented")
	}
}

func (c *Compiler) compileFunctionCall(f *ir.Func, args ...value.Value) (value.Value, error) {
	if len(f.Sig.Params) > len(args) {
		return nil, fmt.Errorf("function %s required %d parameter, got %d", f.Name(), len(f.Sig.Params), len(args))
	}

	converted := []value.Value{}
	for idx, paramType := range f.Sig.Params {
		if args[idx].Type().Equal(paramType) {
			converted = append(converted, args[idx])
		} else {
			ptr := c.currentBlock.NewBitCast(args[idx], types.I8Ptr)
			converted = append(converted, ptr)
			// log.Printf("arg %T %v", args[idx], args[idx])
			// return nil, fmt.Errorf("parameter mismatch, function %s require %v parameter but got %v", f.Name(), paramType, args[idx].Type())
		}
	}
	converted = append(converted, args[len(f.Sig.Params):]...)

	return c.currentBlock.NewCall(f, converted...), nil
}
