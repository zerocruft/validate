package validate

import (
	"testing"
)

func TestRequired(t *testing.T) {
	x := TestStruct{"a", "b"}
	err, _ := Required(x)
	if err != nil {
		t.Fatal(err)
	}
}

type TestStruct struct {
	A string `validate:"required"`
	B string `validate:"required"`
}
