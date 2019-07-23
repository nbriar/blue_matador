package stats

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// CalculateStats - Prints out the mean, min, max, median, sum and stddev of a list of integers input as a space separated string
func CalculateStats(list string) {
	il := convertStringListToIntegerList(list)
	if len(il) == 0 {
		fmt.Println("You must supply at least 1 number.")
		return
	}

	fmt.Println("Input: ", il)
	sum, mean := calcMean(il)
	fmt.Println("Sum: ", sum)
	fmt.Println("Mean: ", mean)

	calcMinMax(il)
	calcMedian(il)
	calcStdDev(il)
}

func convertStringListToIntegerList(list string) (il []int) {
	sl := strings.Split(list, " ")
	for _, s := range sl {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Part of the string you provided is not an integer!")
			return
		}

		il = append(il, i)
	}
	return il
}

func calcMean(il []int) (float64, float64) {
	var total float64
	for _, value := range il {
		total += float64(value)
	}
	mean := total / float64(len(il))

	return total, mean
}

func calcMedian(il []int) {
	sort.Ints(il)
	l := len(il)
	middle := l / 2

	if l%2 != 0 {
		fmt.Println("Median: ", il[middle])
		return
	}

	median := float64(il[middle-1]+il[middle]) / 2
	fmt.Println("Median: ", median)
}

func calcMinMax(il []int) {
	// initialize with the first value
	min := il[0]
	max := il[0]
	for _, value := range il {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}

	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}

func calcStdDev(il []int) {
	// shamelessly stolen and modified from https://github.com/montanaflynn/stats
	_, m := calcMean(il)

	var variance float64
	for _, n := range il {
		variance += (float64(n) - m) * (float64(n) - m)
	}

	vp := variance / float64((len(il)))

	fmt.Println("StdDev: ", math.Pow(vp, 0.5))
}
