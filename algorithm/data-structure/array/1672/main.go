package main

func maximumWealth(accounts [][]int) (ans int) {

	for _, account := range accounts {

		sum := 0

		for _, val := range account {
			sum += val
		}

		if sum > ans {
			ans = sum
		}
	}

	return
}
