package ast

const NTFormals NodeType = "formals"

type Formals struct {
	Formals []Identifier
}

func (*Formals) NodeType() NodeType {
	return NTFormals
}

func NewFormals(formals []Identifier) *Formals {
	return &Formals{
		Formals: formals,
	}
}
