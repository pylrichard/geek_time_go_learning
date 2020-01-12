package empty_interface_test

import (
	"fmt"
	"testing"
)

func GetInterfaceType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknown Type")
	}
}

func TestEmptyInterface(t *testing.T) {
	GetInterfaceType(10)
	GetInterfaceType("10")
}