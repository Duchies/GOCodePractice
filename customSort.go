package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//Sorting a slice of Strings
	strs := []string{"quick", "brown", "fox", "jumps"}
	sort.Strings(strs)
	fmt.Println("Sorted strings: ", strs)

	// Sorting a slice of Integers
	ints := []int{56, 19, 78, 67, 14, 25}
	sort.Ints(ints)
	fmt.Println("Sorted integers: ", ints)

	// Sorting a slice of Floats
	floats := []float64{176.8, 19.5, 20.8, 57.4}
	sort.Float64s(floats)
	fmt.Println("Sorted floats: ", floats)

	strs2 := []string{"United States", "India", "France", "United Kingdom", "Spain"}
	sort.Slice(strs2, func(i, j int) bool {
		return len(strs2[i]) < len(strs2[j])
	})

	// Stable sort
	sort.SliceStable(strs2, func(i, j int) bool {
		return len(strs2[i]) < len(strs2[j])
	})
	fmt.Println("[Stable] Sorted strings by length: ", strs2)

	// custom sort example question 
	//https://leetcode.com/problems/largest-number/
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
}

func largestNumber(nums []int) string {

	var arr []string

	for _, num := range nums {
		arr = append(arr, strconv.Itoa(num))
	}

	sort.Slice(arr, func(i, j int) bool {
		return isLarger(arr[i], arr[j])
	})

	fmt.Println("this is array after soring", arr)

	return strings.Join(arr, "")

}

func isLarger(str1, str2 string) bool {
	if str1 == str2 {
		return false
	}

	length := len(str1)

	if len(str2) < length {
		length = len(str2)
	}

	i := 0
	for ; i < length; i++ {
		if str1[i] > str2[i] {
			return true
		} else if str1[i] < str2[i] {
			return false
		}
	}

	if i == len(str1) {
		return isLarger(str1, str2[i:])
	}

	return isLarger(str1[i:], str2)
}
