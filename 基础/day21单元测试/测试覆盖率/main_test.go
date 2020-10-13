package 测试覆盖率

import (
	"mytest"
	"testing"
)

func TestTag(t *testing.T) {
	mytest.Tag(1)
	mytest.Tag(2)
	mytest.Tag(3)
	mytest.Tag(6)

}
