package main

import "fmt"

func main() {
    fmt.Println("Try programiz.pro")

    arr := make([]int, 0) // Create an empty slice
    list(arr)
    fmt.Println(arr)
    
    helper(&arr) // Pass the slice pointer to the helper function
    fmt.Println(arr)      // Print the modified slice
}

func list(arr []int){
    arr = append(arr,9)
}

func helper(arr *[]int) {
    *arr = append(*arr, 1)
    *arr = append(*arr, 2)
    *arr = append(*arr, 3)
}

/*
output
Try programiz.pro
[]
[1 2 3]
*/
