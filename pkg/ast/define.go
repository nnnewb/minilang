package ast

const NTDefine NodeType = "define"

type Define struct {
	Identifier  Identifier
	Combination *Combination
}

func (*Define) NodeType() NodeType { return NTDefine }

func NewDefine(identifier Identifier, combination *Combination) *Define {
	return &Define{
		Identifier:  identifier,
		Combination: combination,
	}
}
