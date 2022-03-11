package ast

var NTString NodeType = "string"

type String string

func (String) NodeType() NodeType { return NTString }
