package environment

import (
	"fmt"
)

type ExecutionEnvironment interface {
	LookupName(name string) Value
	LookupLocalName(name string) Value
	// SetValue set new value and return old one
	SetValue(name string, val Value) Value
}

type ExecutionEnv struct {
	symbols map[string]Value
	parent  ExecutionEnvironment
}

func NewExecutionEnv(parent ExecutionEnvironment) *ExecutionEnv {
	return &ExecutionEnv{
		symbols: make(map[string]Value),
		parent:  parent,
	}
}

func (ee *ExecutionEnv) SetValue(name string, val Value) Value {
	old, ok := ee.symbols[name]
	ee.symbols[name] = val
	if ok {
		return old
	}
	return nil
}

func (ee *ExecutionEnv) LookupName(name string) Value {
	if val := ee.LookupLocalName(name); val != nil {
		return val
	}
	return ee.parent.LookupName(name)
}

func (ee *ExecutionEnv) LookupLocalName(name string) Value {
	if val, ok := ee.symbols[name]; ok {
		return val
	}
	return nil
}

func (ee *ExecutionEnv) EvaluateListElements(list List) (List, error) {
	evaluated := make(List, 0, len(list))
	for _, val := range list {
		if val.GetType() == VTList {
			if result, err := ee.EvaluateFunctionCall(val.GetValue().(List)); err != nil {
				return nil, err
			} else {
				evaluated = append(evaluated, result)
			}
		} else {
			evaluated = append(evaluated, val)
		}
	}
	return evaluated, nil
}

func (ee *ExecutionEnv) EvaluateFunctionCall(list List) (Value, error) {
	if len(list) > 0 {
		first := list[0]
		if first.GetType() != VTIdentifier {
			return nil, fmt.Errorf("first element in list has type %s, which is not callable", first.GetType())
		}

		name, ok := first.GetValue().(Identifier)
		if !ok {
			return nil, fmt.Errorf("internal error, evaluate identifier failed %v", first.GetValue())
		}

		fn := ee.LookupName(string(name))
		if fn == nil {
			return nil, fmt.Errorf("%s not defined in current scope, evaluation failed", name)
		}

		if fn.GetType() != VTBuiltinFn {
			return nil, fmt.Errorf("%s (type: %s) is not callable", name, fn.GetType())
		}

		if args, err := ee.EvaluateListElements(list[1:]); err != nil {
			return nil, err
		} else {
			return fn.GetValue().(BuiltinFunc)(args)
		}
	}
	return nil, nil
}
