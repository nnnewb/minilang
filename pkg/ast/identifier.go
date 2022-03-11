package ast

type Identifier string

var NTIdentifier NodeType = "ident"

func (Identifier) NodeType() NodeType { return NTIdentifier }
