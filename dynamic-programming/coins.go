package main

import (
	"fmt"
	"sort"
)

type CoinCounter struct {
	// Table is structured like table[Coin][Value]
	Table map[int]map[int]int
	// Coins will be sorted
	Coins []int
}

func NewCoinCounter(coins []int) CoinCounter {
	sort.Ints(coins)

	table := make(map[int]map[int]int)
	table[0] = make(map[int]int)
	for _, c := range coins {
		table[c] = make(map[int]int)
	}

	return CoinCounter{table, coins}
}

func (self *CoinCounter) Contains(value, coin int) bool {
	if _, ok := self.Table[coin]; !ok {
		return false
	}

	_, ok := self.Table[coin][value]

	return ok
}

func (self *CoinCounter) Get(value, coin int) (int, bool) {
	if _, ok := self.Table[coin]; !ok {
		return 0, false
	}

	return self.Table[coin][value]
}

func (self *CoinCounter) Set(value, coin, count int) {
	if _, ok := self.Table[coin]; ok {
		self.Table[coin][value] = count
	}
}

func (self *CoinCounter) Count(value int) int {
	return 0
}

func main() {
	fmt.Println("x_x")
}
