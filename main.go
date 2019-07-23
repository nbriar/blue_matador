package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/nbriar/blue_matador/chuck_norris"
	"github.com/nbriar/blue_matador/palindrome"
	"github.com/nbriar/blue_matador/stats"
)

func main() {
	command := os.Args[1]

	if command == "chucky-category-joke" {
		fmt.Println(chuck_norris.JokeByCategory(os.Args[2]))
		return
	}
	if command == "chucky-search-joke" {
		fmt.Println(chuck_norris.JokeSearch(os.Args[2]))
		return
	}
	if command == "validate-integer-palindrome" {
		i, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Input os not an integer!")
		}
		fmt.Println(palindrome.IsPalindrome(int64(i)))
		return
	}
	if command == "calculate-stats" {
		stats.CalculateStats(os.Args[2])
	}

}
