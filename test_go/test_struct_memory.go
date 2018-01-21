package main

import (
	"fmt"
	"reflect"
)

type MyData struct {
	// offset at 0
	aByte byte
	// offset at 8
	aLong int64
	// offset at 16
	aShort int16
	// offset at 20
	aInt32 int32
	// offset at 24
	aSlice []byte
}

var myData MyData

func main() {
	typ := reflect.TypeOf(myData)
	fmt.Printf("struct is %d bytes long\n", typ.Size())

	n := typ.NumField()
	for i := 0; i < n; i++ {
		field := typ.Field(i)
		fmt.Printf("%s at offset %v, size=%d, align=%d\n", field.Name, field.Offset, field.Type.Size(), field.Type.Align())
	}
}
