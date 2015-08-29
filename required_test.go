package validate

import (
	"testing"
	"fmt"
)

func TestRequired(t *testing.T) {
	x := TestStruct{"er", "b", 1}
	err, infractions := Required(x)
	if err != nil {
		fmt.Println(infractions)
		t.Fatal(err)
	}
}

type TestStruct struct {
	A string `validate:"required"`
	B string `validate:"required"`
	C int `validate:"required"`
}
