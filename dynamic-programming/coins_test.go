package main

import (
	"fmt"
	"testing"
)

func xTestCoinCount(t *testing.T) {
	c := NewCoinCounter([]int{1})
	d := NewCoinCounter([]int{5, 1})
	e := NewCoinCounter([]int{2, 6})
	f := NewCoinCounter([]int{1, 2, 3})

	if v := c.Count(1); v != 1 {
		t.Fail()
	}

	if v := c.Count(2); v != 1 {
		t.Fail()
	}

	if v := c.Count(7); v != 1 {
		t.Fail()
	}

	if v := d.Count(5); v != 2 {
		t.Fail()
	}

	if v := e.Count(6); v != 2 {
		t.Fail()
	}

	if v := f.Count(5); v != 5 {
		fmt.Println(v, "!=", 5)
		t.Fail()
	}
}

func TestBroken(t *testing.T) {
	c := NewCoinCounter([]int{1, 2, 3})
	d := NewCoinCounter([]int{2, 3, 5, 6})

	if v := c.Count(5); v != 5 {
		fmt.Println("CoinCount(5, [1, 2, 3]) != 5 ~", v)
		t.Fail()
	}

	if v := c.Count(4); v != 4 {
		fmt.Println("CoinCount(4, [1, 2, 3]) != 4 ~", v)
		t.Fail()
	}

	if v := d.Count(10); v != 5 {
		fmt.Println("CoinCount(10, [2, 3, 5, 6]) != 5 ~", v)
		t.Fail()
	}
}

func x_x() {
	fmt.Println("x_x")
}
