package main

// #include <stdlib.h>
import "C"
import (
	"time"
)

func makeCString(source string) {
	for j := 0; j < 100000; j++ {
		cstr := C.CString(source)
		_ = cstr
	}
	time.Sleep(time.Millisecond * 500)
}

func main()  {
	for i := 0; i < 10; i++ {
		makeCString("hello world")
	}
}