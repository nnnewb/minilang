// Code generated by gocc; DO NOT EDIT.

package parser

import (
	"github.com/nnnewb/minilang/pkg/ast"
	"github.com/nnnewb/minilang/pkg/bnf/token"
)

type (
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib, interface{}) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String:     `S' : Program	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Program : Program Define	<< X[0].(*ast.Program).AddDefine(X[1].(*ast.Define)), nil >>`,
		Id:         "Program",
		NTType:     1,
		Index:      1,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0].(*ast.Program).AddDefine(X[1].(*ast.Define)), nil
		},
	},
	ProdTabEntry{
		String:     `Program : Program Combination	<< X[0].(*ast.Program).AddCombination(X[1].(*ast.Combination)), nil >>`,
		Id:         "Program",
		NTType:     1,
		Index:      2,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0].(*ast.Program).AddCombination(X[1].(*ast.Combination)), nil
		},
	},
	ProdTabEntry{
		String:     `Program : Define	<< ast.NewProgram().AddDefine(X[0].(*ast.Define)), nil >>`,
		Id:         "Program",
		NTType:     1,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewProgram().AddDefine(X[0].(*ast.Define)), nil
		},
	},
	ProdTabEntry{
		String:     `Program : Combination	<< ast.NewProgram().AddCombination(X[0].(*ast.Combination)), nil >>`,
		Id:         "Program",
		NTType:     1,
		Index:      4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewProgram().AddCombination(X[0].(*ast.Combination)), nil
		},
	},
	ProdTabEntry{
		String:     `Define : "(" "define" Identifier Combination ")"	<< ast.NewDefine(X[2].(ast.Identifier), X[3].(*ast.Combination)), nil >>`,
		Id:         "Define",
		NTType:     2,
		Index:      5,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewDefine(X[2].(ast.Identifier), X[3].(*ast.Combination)), nil
		},
	},
	ProdTabEntry{
		String:     `Combination : "(" "lambda" "(" Formals ")" Combination ")"	<< ast.NewCombinationWithOperands(ast.Identifier("lambda"), append([]ast.Node{}, ast.NewFormals(X[3].([]ast.Identifier)), X[5].(*ast.Combination))), nil >>`,
		Id:         "Combination",
		NTType:     3,
		Index:      6,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewCombinationWithOperands(ast.Identifier("lambda"), append([]ast.Node{}, ast.NewFormals(X[3].([]ast.Identifier)), X[5].(*ast.Combination))), nil
		},
	},
	ProdTabEntry{
		String:     `Combination : "(" Identifier Operand ")"	<< ast.NewCombinationWithOperands(X[1].(ast.Identifier), X[2].([]ast.Node)), nil >>`,
		Id:         "Combination",
		NTType:     3,
		Index:      7,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewCombinationWithOperands(X[1].(ast.Identifier), X[2].([]ast.Node)), nil
		},
	},
	ProdTabEntry{
		String:     `Combination : "(" Identifier ")"	<< ast.NewCombination(X[1].(ast.Identifier)), nil >>`,
		Id:         "Combination",
		NTType:     3,
		Index:      8,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewCombination(X[1].(ast.Identifier)), nil
		},
	},
	ProdTabEntry{
		String:     `Formals : Formals Identifier	<< append(X[0].([]ast.Identifier), X[1].(ast.Identifier)), nil >>`,
		Id:         "Formals",
		NTType:     4,
		Index:      9,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return append(X[0].([]ast.Identifier), X[1].(ast.Identifier)), nil
		},
	},
	ProdTabEntry{
		String:     `Formals : Identifier	<< []ast.Identifier{X[0].(ast.Identifier)}, nil >>`,
		Id:         "Formals",
		NTType:     4,
		Index:      10,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return []ast.Identifier{X[0].(ast.Identifier)}, nil
		},
	},
	ProdTabEntry{
		String:     `Operand : Value	<< []ast.Node{X[0].(ast.Node)}, nil >>`,
		Id:         "Operand",
		NTType:     5,
		Index:      11,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return []ast.Node{X[0].(ast.Node)}, nil
		},
	},
	ProdTabEntry{
		String:     `Operand : Operand Value	<< append(X[0].([]ast.Node), X[1].(ast.Node)), nil >>`,
		Id:         "Operand",
		NTType:     5,
		Index:      12,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return append(X[0].([]ast.Node), X[1].(ast.Node)), nil
		},
	},
	ProdTabEntry{
		String:     `BooleanLit : boolean_t	<< ast.Boolean(true), nil >>`,
		Id:         "BooleanLit",
		NTType:     6,
		Index:      13,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.Boolean(true), nil
		},
	},
	ProdTabEntry{
		String:     `BooleanLit : boolean_f	<< ast.Boolean(false), nil >>`,
		Id:         "BooleanLit",
		NTType:     6,
		Index:      14,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.Boolean(false), nil
		},
	},
	ProdTabEntry{
		String:     `Identifier : identifier	<< ast.Identifier(string(X[0].(*token.Token).Lit)), nil >>`,
		Id:         "Identifier",
		NTType:     7,
		Index:      15,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.Identifier(string(X[0].(*token.Token).Lit)), nil
		},
	},
	ProdTabEntry{
		String:     `Value : Identifier	<< X[0].(ast.Identifier), nil >>`,
		Id:         "Value",
		NTType:     8,
		Index:      16,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0].(ast.Identifier), nil
		},
	},
	ProdTabEntry{
		String:     `Value : BooleanLit	<< X[0].(ast.Boolean), nil >>`,
		Id:         "Value",
		NTType:     8,
		Index:      17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0].(ast.Boolean), nil
		},
	},
	ProdTabEntry{
		String:     `Value : uint	<< ast.NewUInt(string(X[0].(*token.Token).Lit)) >>`,
		Id:         "Value",
		NTType:     8,
		Index:      18,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewUInt(string(X[0].(*token.Token).Lit))
		},
	},
	ProdTabEntry{
		String:     `Value : string	<< ast.String(string(X[0].(*token.Token).Lit)), nil >>`,
		Id:         "Value",
		NTType:     8,
		Index:      19,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.String(string(X[0].(*token.Token).Lit)), nil
		},
	},
	ProdTabEntry{
		String:     `Value : Combination	<< X[0], nil >>`,
		Id:         "Value",
		NTType:     8,
		Index:      20,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
}
