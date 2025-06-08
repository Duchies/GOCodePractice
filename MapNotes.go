package main

import (
	"fmt"
)

func main() {
	// 1. map[string]int
	intMap := map[string]int{}
	fmt.Println("intMap[\"foo\"]:", intMap["foo"])                  // 0
	fmt.Println("intMap[\"bar\"] == 0:", intMap["bar"] == 0)        // true

	// 2. map[string]byte
	byteMap := map[string]byte{}
	fmt.Println("byteMap[\"foo\"]:", byteMap["foo"])                // 0
	fmt.Println("byteMap[\"bar\"] == 0:", byteMap["bar"] == 0)      // true

	// 3. map[string]string
	stringMap := map[string]string{}
	fmt.Println("stringMap[\"foo\"]:", stringMap["foo"])            // ""
	fmt.Println("stringMap[\"bar\"] == \"\":", stringMap["bar"] == "") // true

	// 4. map[string]bool
	boolMap := map[string]bool{}
	fmt.Println("boolMap[\"foo\"]:", boolMap["foo"])                // false
	fmt.Println("boolMap[\"bar\"] == false:", boolMap["bar"] == false) // true

	// 5. map[string]*int
	pointerMap := map[string]*int{}
	fmt.Println("pointerMap[\"foo\"]:", pointerMap["foo"])          // <nil>
	fmt.Println("pointerMap[\"bar\"] == nil:", pointerMap["bar"] == nil) // true

	// Optional: using comma-ok to check if key exists
	fmt.Println("\n--- Using comma-ok idiom ---")
	_, ok := stringMap["foo"]
	fmt.Println("stringMap[\"foo\"] exists?:", ok)                  // false
}
