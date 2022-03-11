package ast

type NodeType string

type Node interface {
	NodeType() NodeType
}
