package main

import (
	"fmt"
)

func main() {
	fmt.Println("Main started...")
	// 345. Reverse Vowels of a String
	// Input: s = "IceCreAm"
	// Output: "AceCreIm"

	fmt.Println(reverseVowels("IceCreAm"))
	fmt.Println(reverseVowels("leetcode"))
	fmt.Println("Main ended...")
}

var vowelsMap = map[string]bool{
	`a`: true,
	`e`: true,
	`o`: true,
	`i`: true,
	`u`: true,
	`A`: true,
	`E`: true,
	`O`: true,
	`I`: true,
	`U`: true,
}

//Strings in Go are immutable, which means you cannot modify them directly

//Go does not have a built-in Set data structure like some other languages (e.g., Python or Java). However, you can simulate a Set using a map

func reverseVowels(s string) string {

	vowelIdxSet := make(map[string]int)
	var vowelSeqArr []string

	for i := 0; i < len(s); i++ {
		if _, ok := vowelsMap[string(s[i])]; ok {
			vowelIdxSet[string(s[i])] = i
			vowelSeqArr = append(vowelSeqArr, string(s[i]))
		}
	}

	vowelSeqArr = reverseArray(vowelSeqArr)
	var j = 0

	var result string

	for i := 0; i < len(s); i++ {

		if _, ok := vowelIdxSet[string(s[i])]; ok {
			result = result + vowelSeqArr[j]
			j++
		} else {
			result = result + string(s[i])
		}
	}

	return result

}

func reverseArray(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
