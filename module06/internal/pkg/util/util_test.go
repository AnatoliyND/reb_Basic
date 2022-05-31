package util

import (
	"fmt"
	"testing"
)

func TestReversInt(t *testing.T) {
	res0, err := ReverseInt(123)
	if err != nil {
		t.Error("Fail to get input: %w\n", err)
	}
	if res0 != 321 {
		t.Error("incorrect result\n")
	}
	res1, _ := ReverseInt(-649)
	fmt.Println(res1)
	if res1 != -946 {
		t.Fatal("incorrect negativ result\n")
	}
	res2, _ := ReverseInt(0)
	if res2 != 0 {
		t.Error("Revers 0 != 0\n")
	}

}
