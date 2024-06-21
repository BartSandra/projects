package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	flagMean := flag.Bool("mean", false, "Print mean")
	flagMedian := flag.Bool("median", false, "Print median")
	flagMode := flag.Bool("mode", false, "Print mode")
	flagSD := flag.Bool("sd", false, "Print standard deviation")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	var data []int
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		num, err := strconv.Atoi(scanner.Text())
		if err != nil || num < -100000 || num > 100000 {
			fmt.Println("Enter valid values in the range from -100.000 to 100.000")
			// return
		}
		data = append(data, num)
	}
	sort.Ints(data)

	mean := calculateMean(data)
	if *flagMean || (!*flagMean && !*flagMedian && !*flagMode && !*flagSD) {
		fmt.Println("Mean:", mean)
	}
	median := calculateMedian(data)
	if *flagMedian || (!*flagMean && !*flagMedian && !*flagMode && !*flagSD) {
		fmt.Println("Median:", median)
	}
	mode := calculateMode(data)
	if *flagMode || (!*flagMean && !*flagMedian && !*flagMode && !*flagSD) {
		fmt.Println("Mode:", mode)
	}
	sd := calculateStandardDeviation(data, mean)
	if *flagSD || (!*flagMean && !*flagMedian && !*flagMode && !*flagSD) {
		fmt.Println("SD:", sd)
	}
}

func calculateMean(data []int) float64 {
	sum := 0
	for _, num := range data {
		sum += num
	}
	return float64(sum) / float64(len(data))
}

func calculateMedian(data []int) float64 {
	size := len(data)
	if size%2 == 0 {
		return float64(data[size/2-1]+data[size/2]) / 2.0
	}
	return float64(data[size/2])
}

func calculateMode(data []int) int {
	counts := map[int]int{}

	maxCount := 0
	mode := 100001

	for _, num := range data {
		counts[num]++
		if counts[num] > maxCount || (counts[num] == maxCount && num < mode) {
			maxCount = counts[num]
			mode = num
		}
	}
	return mode
}

func calculateStandardDeviation(data []int, mean float64) float64 {
	sumOfSquaredDiffs := 0.0
	for _, num := range data {
		diff := float64(num) - mean
		sumOfSquaredDiffs += diff * diff
	}
	variance := sumOfSquaredDiffs / float64(len(data)-1)
	return math.Sqrt(variance)
}
