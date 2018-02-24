package main

/*
#include <stdint.h>

// 因为go语言不支持联合体类型，因此联合体在go中用uint8数组表示
// 如果必须要使用联合体，可以使用unsafe包
union B1 {
	int i;
	float j;
};

union B2 {
	int64_t i;
	double j;
};
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	var b1 C.union_B1
	fmt.Printf("%T\n", b1)

	var b2 C.union_B2
	fmt.Printf("%T\n", b2)
	fmt.Printf("%#v\n", *(*C.double)(unsafe.Pointer(&b2)))
}
