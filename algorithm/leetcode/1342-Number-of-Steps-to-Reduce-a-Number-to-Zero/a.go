package main

func numberOfStepsAa(num int) int {

	var count int

	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num--
		}
		count++
	}

	return count
}
