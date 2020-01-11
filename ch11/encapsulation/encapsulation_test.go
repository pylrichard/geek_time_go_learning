package encapsulation_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id		string
	Name	string
	Age		int
}

func (e Employee) String() string {
	fmt.Printf("address: %x\n", unsafe.Pointer(&e.Name))

	return fmt.Sprintf("id:%s-name:%s-age:%d", e.Id, e.Name, e.Age)
}

func TestEmployee(t *testing.T) {
	e1 := Employee{"0", "Bob", 20}
	e2 := Employee{Name: "Mike", Age: 30}
	//new()返回指针
	e3 := new(Employee)
	e3.Id = "2"
	e3.Name = "Rose"
	e3.Age = 22

	t.Log(e1)
	t.Logf("e1 is %T", e1)
	t.Log(e2)
	t.Log(e2.Id)
	t.Log(e3)
	t.Logf("e3 is %T", e3)
}