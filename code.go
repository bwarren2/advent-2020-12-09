package advent20201209

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	// mapset "github.com/deckarep/golang-set"
)

func NumsFromFile(filename string) (returnlist []int64) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		input, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		returnlist = append(returnlist, input)
	}
	// spew.Dump([]string)
	return
}

func IsSum(lst []int64, target int64) bool {
	for i := 0; i < len(lst)-1; i++ {
		for j := i + 1; j < len(lst); j++ {
			if lst[i] != lst[j] && lst[i]+lst[j] == target {
				return true
			}
		}
	}
	return false
}

func Part1(filename string, lenPreamble int64) (answer int64) {
	lst := NumsFromFile(filename)

	for idx := lenPreamble; idx < int64(len(lst)); idx++ {

		if !IsSum(lst[idx-lenPreamble:idx], lst[idx]) {
			fmt.Println(IsSum(lst[idx-lenPreamble:idx], lst[idx]))
			return lst[idx]
		}
	}
	return 0
}
func sum(array []int64) int64 {
	result := int64(0)
	for _, v := range array {
		result += v
	}
	return result
}
func MinMax(array []int64) (int64, int64) {
	var max int64 = array[0]
	var min int64 = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
func SumTo(lst []int64, targetNum int64) (works bool, low, high int64) {
	for i := 0; i < len(lst); i++ {
		sums := sum(lst[:i])
		if sums == targetNum {
			min, max := MinMax(lst[:i])
			return true, min, max
		}
		if sums > targetNum {
			return false, 0, 0
		}
	}
	return false, 0, 0
}

func Part2(filename string, targetNum int64) (total int64) {
	lst := NumsFromFile(filename)
	for i := 0; i < len(lst); i++ {
		success, low, high := SumTo(lst[i:], targetNum)
		if success {
			return low + high
		}
	}
	return 0
}
