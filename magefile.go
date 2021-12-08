//go:build mage

package main

import (
	"github.com/magefile/mage/sh"
)

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

	err = sh.Run("go", "build", "-o", "scsh", "cmd/scsh/main.go")
	if err != nil {
		return err
	}

	return nil
}
