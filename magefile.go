//go:build mage

package main

import (
	"runtime"

	"github.com/magefile/mage/sh"
)

var (
	sccBinaryOutput string = "build/bin/scc"
)

func init() {
	if runtime.GOOS == "windows" {
		sccBinaryOutput += ".exe"
	}
}

func Build() error {
	if err := sh.Run("gocc", "-o", "pkg/bnf", "-p", "github.com/nnnewb/minilang/pkg/bnf", "spec.bnf"); err != nil {
		return err
	}

	if err := sh.Run("go", "mod", "tidy"); err != nil {
		return err
	}

	if err := sh.Run("go", "build", "-o", sccBinaryOutput, "cmd/scc/main.go"); err != nil {
		return err
	}

	return nil
}
