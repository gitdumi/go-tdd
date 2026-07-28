package depinj

import (
	"fmt"
	"io"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
