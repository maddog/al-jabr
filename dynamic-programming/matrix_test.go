package main

import (
	"fmt"
	"testing"
)

func TestMatrix(t *testing.T) {
	m := NewMatrix(3, 4)

	m.Set(2, 3, 1.0)

	if m.Get(2, 3) != 1.0 {
		t.Fail()
	}

	if m.table[11] != 1.0 {
		t.Fail()
	}

	m.Set(1, 2, 5.0)

	if m.table[6] != 5.0 {
		fmt.Println(m.table)
		t.Fail()
	}
}

func _() {
	fmt.Println("x_x")
}
