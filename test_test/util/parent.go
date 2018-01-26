package util

import (
	"fmt"
	"test/test_test"
)

func Parent() {
	test_test.Child()
	fmt.Println("parent")
}
