package builtin

import (
	"github.com/nnnewb/minilang/internal/environment"
	"github.com/nnnewb/minilang/internal/utils"
)

func RegisterArithmeticBuiltin(ee *environment.ExecutionEnv) {
	ee.SetValue("+", utils.Coverup(environment.BuiltinFunc(sum)))
	ee.SetValue("-", utils.Coverup(environment.BuiltinFunc(sub)))
}

func sum(ee *environment.ExecutionEnv, args []environment.Value) (environment.Value, error) {
	if len(args) == 0 {
		return environment.Number(0), nil
	}

	var accumulator float64
	for _, arg := range args {
		accumulator += float64(utils.MustToNumber(arg))
	}

	return environment.Number(accumulator), nil
}

func sub(ee *environment.ExecutionEnv, args []environment.Value) (environment.Value, error) {
	if len(args) == 0 {
		return environment.Number(0), nil
	}

	accumulator := float64(utils.MustToNumber(args[0]))
	for _, arg := range args[1:] {
		accumulator -= float64(utils.MustToNumber(arg))
	}

	return environment.Number(accumulator), nil
}
