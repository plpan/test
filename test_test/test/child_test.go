package child_test

import (
	. "test/test_test"
	"test/test_test/util"
	"testing"
)

func TestChild(t *testing.T) {
	Child()
	util.Parent()
	t.FailNow()
}
