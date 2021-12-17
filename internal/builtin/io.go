package builtin

import (
	"fmt"
	"strings"

	"github.com/nnnewb/minilang/internal/vm"
)

func RegisterIO(m *vm.MiniVM) {
	RegisterBuiltinFunc(m, "display", func(m *vm.MiniVM, argc int, argv []vm.Object) (vm.Object, error) {
		sb := strings.Builder{}
		for _, v := range argv {
			sb.WriteString(fmt.Sprintf("%v", v))
			sb.WriteString(" ")
		}
		println(sb.String())
		return vm.Nil{}, nil
	})
}
