package child_test

import (
	. "test/test_test/test_normal"
	"test/test_test/test_normal/util"
	"testing"
)

func TestChild(t *testing.T) {
	Child()
	util.Parent()
	t.FailNow()
}
