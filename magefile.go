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

func Gen() error {
	return sh.Run("go", "generate", "./...")
}

func Build() error {
	var err error

	err = sh.Run("go", "mod", "download")
	if err != nil {
		return err
	}

	err = sh.Run("go", "mod", "tidy")
	if err != nil {
		return err
	}

	err = Gen()
	if err != nil {
		return err
	}

	err = sh.Run("go", "build", "-o", binaryOutput, "cmd/scsh/main.go")
	if err != nil {
		return err
	}

	return nil
}
