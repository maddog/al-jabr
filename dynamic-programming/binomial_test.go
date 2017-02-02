package main

import (
	"testing"
)

func arrayEquals(lhs, rhs []int) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for i := 0; i < len(lhs); i++ {
		if lhs[i] != rhs[i] {
			return false
		}
	}

	return true
}

func TestArrayEquals(t *testing.T) {
	if arrayEquals([]int{1, 2, 3}, []int{1, 2, 4}) {
		t.Fail()
	}
	if !arrayEquals([]int{1, 2, 3}, []int{1, 2, 3}) {
		t.Fail()
	}
}

func TestBinomial(t *testing.T) {
	if !arrayEquals(Binomial(0), []int{1}) {
		t.Fail()
	}
	if !arrayEquals(Binomial(1), []int{1, 1}) {
		t.Fail()
	}
	if !arrayEquals(Binomial(2), []int{1, 2, 1}) {
		t.Fail()
	}
	if !arrayEquals(Binomial(3), []int{1, 3, 3, 1}) {
		t.Fail()
	}
	if !arrayEquals(Binomial(4), []int{1, 4, 6, 4, 1}) {
		t.Fail()
	}
}
