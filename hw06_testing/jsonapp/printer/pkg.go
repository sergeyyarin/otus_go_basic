package printer

import (
	"fmt"

	"github.com/sergeyyarin/otus_go_basic/hw06_testing/jsonapp/types"
)

func PrintStaff(staff []types.Employee) string {
	var printed string
	for i := 0; i < len(staff); i++ {
		printed += fmt.Sprint(staff[i])
	}
	return printed
}
