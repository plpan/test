package main

// #cgo LDFLAGS: -L. -lsayhello
// void SayHello(const char *s);
import "C"

func main() {
	C.SayHello(C.CString("Hello, World."))
}
