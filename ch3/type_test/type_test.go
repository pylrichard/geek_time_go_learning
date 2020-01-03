package type_test

import (
	"math"
	"runtime"
	"strconv"
	"testing"
)

func TestInt(t *testing.T) {
	cpu := runtime.GOARCH
	size := strconv.IntSize

	t.Log(cpu, size)
}

func TestString(t *testing.T) {
	var s string
	//string是值类型，默认初始化值是空字符串，不是nil
	t.Log("*" + s + "*")
	t.Log(len(s))
	t.Logf("%T %q", s, s)
}

func TestConvert(t *testing.T) {
	i := int64(10)
	j := int8(i)
	t.Log(j)

	a := int64(math.MaxInt64)
	b := int32(a)
	t.Log(a, b)
}

func TestPtr(t *testing.T) {
	a := 1
	ptr := &a
	//不支持指针运算
	//ptr = ptr + 1
	t.Log(a, ptr)
	t.Logf("%T %T", a, ptr)
}