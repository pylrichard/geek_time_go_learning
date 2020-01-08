package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func calculateTime(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())

		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)

	return op
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}

	return ret
}

func Clear() {
	fmt.Println("Clear resources")
}

func TestReturnValue(t *testing.T) {
	ret, _ := returnMultiValues()
	t.Log(ret)
}

func TestTime(t *testing.T) {
	fn := calculateTime(slowFunc)
	t.Log(fn(10))
}

func TestSum(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	panic("err")
}



