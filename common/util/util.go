package util

import (
	"fmt"
	"os"
)

// ErrPrint calls fmt.Fprint to stderr
func ErrPrint(a ...interface{}) (n int, err error) {
	n, err = fmt.Fprint(os.Stderr, a...)
	return
}

// ErrPrintln calls fmt.Fprintln to stderr
func ErrPrintln(a ...interface{}) (n int, err error) {
	n, err = fmt.Fprintln(os.Stderr, a...)
	return
}

// ErrPrintf calls fmt.Fprintf to stderr
func ErrPrintf(format string, a ...interface{}) (n int, err error) {
	n, err = fmt.Fprintf(os.Stderr, format, a...)
	return
}
