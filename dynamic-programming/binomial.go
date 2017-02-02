package main

func Binomial(power int) []int {
	// ...
	lookup := make(map[int]map[int]int)

	lookup[0] = make(map[int]int)

	// ...
	lookup[0][0] = 1

	// ...
	for i := 1; i < power+1; i++ {
		lookup[i] = make(map[int]int)
		for j := 0; j <= i; j++ {
			lookup[i][j] = lookup[i-1][j-1] + lookup[i-1][j]
		}
	}

	coefficients := make([]int, power+1)

	for i, _ := range coefficients {
		coefficients[i] = lookup[power][i]
	}

	// ...
	return coefficients
}
