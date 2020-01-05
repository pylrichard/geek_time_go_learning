package condition

import (
	"errors"
	"testing"
)

func testIf() (int, error) {
	return 1, errors.New("error")
}

func TestIfCondition(t *testing.T) {
	if v, err := testIf(); err != nil {
		t.Log(err)
	} else {
		t.Log(v)
	}
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("4, 5")
		}
	}
}

func TestSwitchCaseConditions(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i % 2 == 0:
			t.Log("Even")
		case i % 2 == 1:
			t.Log("Odd")
		default:
			t.Log("Unknown")
		}
	}
}