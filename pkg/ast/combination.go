package ast

const NTCombination NodeType = "combination"

type Combination struct {
	Operator Identifier
	Operands []Node
}

func NewCombination(operator Identifier) *Combination {
	return &Combination{
		Operator: operator,
		Operands: []Node{},
	}
}

func NewCombinationWithOperands(operator Identifier, operands []Node) *Combination {
	return &Combination{
		Operator: operator,
		Operands: operands,
	}
}

func (*Combination) NodeType() NodeType { return NTCombination }
