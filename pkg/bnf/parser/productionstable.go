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
		String: `S' : Value	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Value : identifier	<< ast.Identifier(string(X[0].(*token.Token).Lit)), nil >>`,
		Id:         "Value",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.Identifier(string(X[0].(*token.Token).Lit)), nil
		},
	},
	ProdTabEntry{
		String: `Value : boolean_t	<< ast.Boolean(true), nil >>`,
		Id:         "Value",
		NTType:     1,
		Index:      2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.Boolean(true), nil
		},
	},
	ProdTabEntry{
		String: `Value : boolean_f	<< ast.Boolean(false), nil >>`,
		Id:         "Value",
		NTType:     1,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.Boolean(false), nil
		},
	},
	ProdTabEntry{
		String: `Value : number	<< ast.NewNumber(string(X[0].(*token.Token).Lit)) >>`,
		Id:         "Value",
		NTType:     1,
		Index:      4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewNumber(string(X[0].(*token.Token).Lit))
		},
	},
	ProdTabEntry{
		String: `Value : string	<< ast.String(string(X[0].(*token.Token).Lit)), nil >>`,
		Id:         "Value",
		NTType:     1,
		Index:      5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.String(string(X[0].(*token.Token).Lit)), nil
		},
	},
	ProdTabEntry{
		String: `Value : List	<< X[0], nil >>`,
		Id:         "Value",
		NTType:     1,
		Index:      6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ListElements : Value	<< ast.NewListWithInitial(X[0].(ast.Node)), nil >>`,
		Id:         "ListElements",
		NTType:     2,
		Index:      7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewListWithInitial(X[0].(ast.Node)), nil
		},
	},
	ProdTabEntry{
		String: `ListElements : ListElements Value	<< X[0].(*ast.List).Append(X[1].(ast.Node)), nil >>`,
		Id:         "ListElements",
		NTType:     2,
		Index:      8,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0].(*ast.List).Append(X[1].(ast.Node)), nil
		},
	},
	ProdTabEntry{
		String: `List : "(" ListElements ")"	<< X[1], nil >>`,
		Id:         "List",
		NTType:     3,
		Index:      9,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `List : "(" ")"	<< ast.NewList(), nil >>`,
		Id:         "List",
		NTType:     3,
		Index:      10,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewList(), nil
		},
	},
}