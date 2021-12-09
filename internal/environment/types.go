package environment

import (
	"fmt"
)

type ExecutionEnv struct {
	symbols map[string]Value
	parent  *ExecutionEnv
}

func NewExecutionEnv(parent *ExecutionEnv) *ExecutionEnv {
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
	if ee.parent != nil {
		return ee.parent.LookupName(name)
	}
	return nil
}

func (ee *ExecutionEnv) LookupLocalName(name string) Value {
	if val, ok := ee.symbols[name]; ok {
		return val
	}
	return nil
}

func (ee *ExecutionEnv) EvaluateList(list List) (Value, error) {
	if len(list) > 0 {
		first, err := ee.Evaluate(list[0])
		if err != nil {
			return nil, err
		}

		if fn, ok := first.(BuiltinFunc); !ok {
			return nil, fmt.Errorf("TypeError: %v(%T) is not callable", first, first)
		} else {
			args := make([]Value, 0, len(list[1:]))
			for _, v := range list[1:] {
				arg, err := ee.Evaluate(v)
				if err != nil {
					return nil, err
				}

				args = append(args, arg)
			}
			return fn(ee, args)
		}
	}
	return nil, nil
}

func (ee *ExecutionEnv) Evaluate(val Value) (Value, error) {
	switch v := val.(type) {
	case *List:
		return ee.EvaluateList(*v)
	case Identifier:
		return ee.LookupName(string(v)), nil
	case *Quoted:
		return v.GetValue().(Value), nil
	default:
		return v, nil
	}
}
