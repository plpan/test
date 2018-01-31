package util

import (
	"fmt"
	"test/test_test/test_normal"
)

func Parent() {
	test_test.Child()
	fmt.Println("parent")
}
