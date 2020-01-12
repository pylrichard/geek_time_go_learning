package polymorphism_test

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	Write() Code
}

type GoProgrammer struct {
}

func (gp *GoProgrammer) Write() Code {
	return "fmt.Println(\"Hello World!\")"
}

type JavaProgrammer struct {
}

func (jp *JavaProgrammer) Write() Code {
	return "System.out.Println(\"Hello World!\")"
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.Write())
}

func TestPolymorphism(t *testing.T) {
	gp := &GoProgrammer{}
	jp := new(JavaProgrammer)
	writeFirstProgram(gp)
	writeFirstProgram(jp)
}