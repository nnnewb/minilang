package utils

import (
	"fmt"

	"github.com/nnnewb/minilang/internal/environment"
)

func Coverup(f environment.BuiltinFunc) environment.BuiltinFunc {
	return func(env *environment.ExecutionEnv, args []environment.Value) (ret environment.Value, err error) {
		defer func() {
			e := recover()
			if e != nil {
				ret = nil
				err = e.(error)
				return
			}
		}()

		return f(env, args)
	}
}

func MustToNumber(val environment.Value) environment.Number {
	if v, ok := val.(environment.Number); !ok {
		panic(fmt.Errorf("TypeError: %v(%T) is not a number", val, val))
	} else {
		return v
	}
}
