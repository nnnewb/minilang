package builtin

import (
	"fmt"

	"github.com/nnnewb/minilang/internal/environment"
)

func RegisterArithmeticBuiltin(ee *environment.ExecutionEnv) {
	ee.SetValue("+", environment.BuiltinFunc(sum))
	ee.SetValue("-", environment.BuiltinFunc(sub))
}

func sum(ee *environment.ExecutionEnv, args []environment.Value) (environment.Value, error) {
	if len(args) == 0 {
		return environment.Number(0), nil
	}

	var accumulator float64
	for _, v := range args {
		if val, ok := v.(environment.Number); ok {
			accumulator += float64(val)
		} else {
			return nil, fmt.Errorf("unexpected argument %v", v)
		}
	}

	return environment.Number(accumulator), nil
}

func sub(ee *environment.ExecutionEnv, args []environment.Value) (environment.Value, error) {
	if len(args) == 0 {
		return environment.Number(0), nil
	}

	var accumulator float64
	if val, ok := args[0].(environment.Number); ok {
		accumulator += float64(val)
	} else {
		return nil, fmt.Errorf("unexpected argument %v", args[0])
	}

	for _, v := range args[1:] {
		if val, ok := v.(environment.Number); ok {
			accumulator -= float64(val)
		} else {
			return nil, fmt.Errorf("unexpected argument %v", v)
		}
	}

	return environment.Number(accumulator), nil
}
