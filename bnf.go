//go:build ignore
// +build ignore

//go:generate gocc -o pkg/bnf -p github.com/nnnewb/minilang/pkg/bnf spec.bnf
package minilang
