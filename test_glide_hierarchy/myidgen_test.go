package test_glide_hierarchy

import (
	"testing"
	"fmt"
)

func TestNextId(t *testing.T) {
	err, id := NextId()
	if err != nil {
		t.FailNow()
	}
	if id == 0 {
		t.FailNow()
	}
	fmt.Printf("gen id : %v\n", id)
}