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

	val, ok := self.Table[coin][value]

	return val, ok
}

func (self *CoinCounter) Set(value, coin, count int) {
	if _, ok := self.Table[coin]; ok {
		self.Table[coin][value] = count
	}
}

func (self *CoinCounter) Count(amount int) int {

	coins := make([]int, 0)
	coins = append(coins, 0)
	coins = append(coins, self.Coins...)

	height, width := len(self.Coins)+1, amount+1

	solution := make([][]int, height)

	for i := 0; i < height; i++ {
		solution[i] = make([]int, width)
		solution[i][0] = 1
	}

	for i := 1; i < width; i++ {
		solution[0][i] = 0
	}

	for i := 1; i <= len(self.Coins); i++ {
		for j := 1; j <= amount; j++ {
			if coins[i-1] <= j && j-self.Coins[i-1] >= 0 {
				solution[i][j] = solution[i-1][j] + solution[i][j-self.Coins[i-1]]
			} else {
				solution[i][j] = solution[i-1][j]
			}
		}
	}

	return solution[len(self.Coins)][amount]
}

func main() {
	fmt.Println("x_x")
}
