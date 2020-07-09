package stdio

import "fmt"

// Printf - stdio.h/printf
func Printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func Sprintf(format string, a []interface{}) {
}

func Puts(s string) {
}
