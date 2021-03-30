package main

import (
	"fmt"
	"io"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Jack")
}

/*NOTE
package fmt
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	//...
}

package io
type Writer interface {
    Write(p []byte) (n int, err error)
}
io.Writer 是一个很好的通用接口，用于「将数据放在某个地方」

fmt.Fprintf 允许传入一个 io.Writer 接口，我们知道 os.Stdout 和 bytes.Buffer 都实现了它
*/