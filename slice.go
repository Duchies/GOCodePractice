package main

import "fmt"

func main() {
    fmt.Println("Main started..")
    var arr []int
    arr = append(arr,2)
    var arr2 []int
    
    arr2 = append(arr2,7)
    arr2 = append(arr2,9)

    arr2 = append(arr2,arr...)
    
    fmt.Println(arr)
    fmt.Println(arr2)

    //learning is append only work with []int. not with [len]int
    
    var a [2]string
  	a[0] = "Hello"
  	a[1] = "World"
  	fmt.Println(a[0], a[1])
  	fmt.Println(a)
  
  	primes := [6]int{2, 3, 5, 7, 11, 13}
  	fmt.Println(primes)

    //https://leetcode.com/problems/car-pooling/submissions/1553921473
    
    fmt.Println("Main ended...")
}
