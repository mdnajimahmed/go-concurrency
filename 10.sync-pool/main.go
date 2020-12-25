package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		fmt.Println("allocating new bytes.Buffer")
		return new(bytes.Buffer)
	},
}

func printLog(w io.Writer, message string) {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	b.WriteString(time.Now().Format("15:04:05"))
	b.WriteString(" : ")
	b.WriteString(message)
	b.WriteString("\n")
	w.Write(b.Bytes())
	bufPool.Put(b)
}

func main() {
	fmt.Println("Hi")
	printLog(os.Stdout, "debug-string-1")
	printLog(os.Stdout, "debug-string-2")
}
