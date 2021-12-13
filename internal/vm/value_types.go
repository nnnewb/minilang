package vm

import (
	"log"

	"github.com/nnnewb/minilang/pkg/ast"
)

type Object interface {
	TypeName() string
}

type UInt uint64

func (u UInt) TypeName() string {
	return "UInt"
}

type Float float64

func (f Float) TypeName() string {
	return "Float"
}

type Boolean bool

func (b Boolean) TypeName() string {
	return "boolean"
}

type String string

func (s String) TypeName() string {
	return "string"
}

type Symbol string

func (s Symbol) TypeName() string {
	return "symbol"
}

type Nil struct{}

func (n Nil) TypeName() string {
	return "nil"
}

type List struct {
	underlying []Object
}

func (l *List) TypeName() string {
	return "list"
}

type Procedure struct {
	Location  int
	Code      []Instruction
	Builtin   NativeFunc
	isBuiltin bool
}

func (p *Procedure) TypeName() string {
	return "procedure"
}

func NewProcedureFromInstruction(instructions []Instruction) *Procedure {
	return &Procedure{
		Location:  0,
		Code:      instructions,
		Builtin:   nil,
		isBuiltin: false,
	}
}

func NewProcedureFromNativeFunc(fun NativeFunc) *Procedure {
	return &Procedure{
		Location:  0,
		Code:      nil,
		Builtin:   fun,
		isBuiltin: true,
	}
}

func ObjectFromLiteral(node ast.Node) Object {
	switch node := node.(type) {
	case ast.Float:
		return Float(node)
	case ast.UInt:
		return UInt(node)
	case ast.String:
		return String(node)
	case ast.Boolean:
		return Boolean(node)
	case *ast.Quoted:
		lst := &List{}
		for _, n := range node.GetValue().([]ast.Node) {
			lst.underlying = append(lst.underlying, ObjectFromLiteral(n))
		}
		return lst
	case ast.Symbol:
		return Symbol(node)
	default:
		log.Fatalf("unexpected ast node in ToValue %v(%T)", node, node)
	}
	return nil
}
