package main

import (
	"fmt"
	"strings"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func statName(name string) (sum int) {
	for _, litter := range name {
		if litter == 'e' {
			sum += 1
		} else if litter == 'i' {
			sum += 2
		} else if litter == 'o' {
			sum += 3
		} else if litter == 'u' {
			sum += 4
		}
	}
	return
}

func dispatchCoin(num int, people []string) (sum int) {
	for _, name := range people {
		num := statName(strings.ToLower(name))
		distribution[name] = num
		sum += num
	}
	return
}

func main() {
	left := dispatchCoin(coins, users)
	fmt.Println("剩下：", left)
	fmt.Println(distribution)
}
