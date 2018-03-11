package main

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

type User struct {
	Name string
	Age  int
}

func test_m() []byte {
	u := User{
		"stupig",
		12,
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	buf, _ := json.Marshal(u)
	fmt.Println(buf)

	return buf
}

func test_u(buf []byte) {
	var u User

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(buf, &u)

	fmt.Printf("%#v\n", u)
}

func main() {
	res := test_m()
	test_u(res)
}
