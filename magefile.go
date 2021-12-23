//go:build mage

package main

import (
	"os"
	"runtime"
	"syscall"

	"github.com/magefile/mage/sh"
)

var (
	binaryOutput string = "build/bin/scsh"
)

func init() {
	if runtime.GOOS == "windows" {
		binaryOutput += ".exe"
	}
}

func Run() error {
	return syscall.Exec(binaryOutput, []string{}, os.Environ())
}

func Build() error {
	if err := sh.Run("gocc", "-o", "pkg/bnf", "-p", "github.com/nnnewb/minilang/pkg/bnf", "spec.bnf"); err != nil {
		return err
	}

	if err := sh.Run("go", "mod", "tidy"); err != nil {
		return err
	}

	if err := sh.Run("go", "build", "-o", binaryOutput, "cmd/scsh/main.go"); err != nil {
		return err
	}

	return nil
}
