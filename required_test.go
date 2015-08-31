package validate

import (
	"fmt"
	"testing"
	"time"
)

func TestBasicRequired(t *testing.T) {
	fmt.Println("Test Basic")
	x := testStructA{"er", "b", 1}
	err, _ := Required(x)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRequiredWithSomeOptionalFields(t *testing.T) {
	fmt.Println("Test Optional Fields")
	x := testStructB{99, ""}
	err, _ := Required(x)
	if err != nil {
		t.Fatal(err)
	}
}

func TestZeroValuesInRequired(t *testing.T) {
	fmt.Println("Test Zero Values")
	x := testStructA{}
	err, _ := Required(x)
	if err == nil {
		t.Fatal(err)
	}

}

func TestNonBasicFields(t *testing.T) {
	fmt.Println("Test Non Basic Fields: Structs")
	x := testStructC{time.Now()}
	err, _ := Required(x)
	if err != nil {
		t.Fatal(err)
	}
}

type testStructA struct {
	A string `validate:"required"`
	B string `validate:"required"`
	C int    `validate:"required"`
}

type testStructB struct {
	A int `validate:"required"`
	B string
}

type testStructC struct {
	B time.Time `validate:required"`
}
