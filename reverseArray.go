package main

import "fmt"

func main() {
	fmt.Println("Main started...")
	// program to reverse the array 
	var arr = []int{1,2,3,4,5}

	fmt.Println("input array : ",arr)

	var j = len(arr)-1

	for i := 0;i < len(arr)/2;i++ {
		swap(&arr[i], &arr[j])
        j--
	}

    fmt.Println("output array : ",arr)
    
	fmt.Println("Main eneded...")
}
func swap(a, b *int) {
	*a, *b = *b, *a
}

// func swap(a,b *int)  {
// 	var c *int
// 	c = a
// 	a = b
// 	b = c
// }
