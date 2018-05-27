package main

import "fmt"

func test_slice1() {
	fmt.Println("---test_slice1---")
	s := []byte("")
	s1 := append(s, 'a')
	s2 := append(s, 'b')
	fmt.Println(string(s1) + " --- " + string(s2))
}

func test_slice2() {
	fmt.Println("\n---test_slice2---")
	s := []byte("")
	s1 := append(s, 'a')
	s2 := append(s, 'b')
	fmt.Println(s1, " --- ", s2)
	fmt.Println(string(s1) + " --- " + string(s2))
}

func test_slice3() {
	fmt.Println("\n---test_slice3---")
	s := make([]byte, 0)
	s1 := append(s, 'a')
	s2 := append(s, 'b')
	fmt.Println(string(s1) + " --- " + string(s2))
}

func test_slice4() {
	fmt.Println("\n---test_slice4---")
	s := make([]byte, 0)
	s1 := append(s, 'a')
	s2 := append(s, 'b')
	fmt.Println(s1, " --- ", s2)
	fmt.Println(string(s1) + " --- " + string(s2))
}

func main() {
	test_slice1()
	test_slice2()
	test_slice3()
	test_slice4()
}
