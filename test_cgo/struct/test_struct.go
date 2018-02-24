package main

/*
struct A {
	int type; // 因为type是go关键字，在访问这个字段时需要加前缀下划线
};
*/
import "C"
import "fmt"

func main() {
	var a C.struct_A
	fmt.Println(a._type)
}
