package ast

var NTProgram NodeType = "program"

type Program struct {
	Combinations []*Combination
	Defines      []*Define
}

func NewProgram() *Program {
	return &Program{
		Combinations: make([]*Combination, 0),
		Defines:      make([]*Define, 0),
	}
}

func (p *Program) NodeType() NodeType { return NTProgram }

func (p *Program) AddDefine(def *Define) *Program {
	p.Defines = append(p.Defines, def)
	return p
}

func (p *Program) AddCombination(comb *Combination) *Program {
	p.Combinations = append(p.Combinations, comb)
	return p
}
