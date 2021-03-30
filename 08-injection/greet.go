package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	_, _ = fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Jack")
}

func main() {
	Greet(os.Stdout, "Jack")

	_ = http.ListenAndServe(":8089", http.HandlerFunc(MyGreeterHandler))
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

将 Greet 的参数定义为 io.Writer 类型，更加通用

curl 127.0.0.1:8089
> Hello, Jack
*/