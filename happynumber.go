package main

import (
	"strconv"
)

func isHappy(n int) bool {
	set := make(map[int]bool)

	s := strconv.Itoa(n)
	sum := 0
	for sum != 1 {
		sum = 0
		for _, v := range s {
			a, _ := strconv.Atoi(string(v))
			sum += a * a
		}
		if _, ok := set[sum]; ok {
			return false
		}

		set[sum] = true

		s = strconv.Itoa(sum)
		

	}

	return sum == 1
}
