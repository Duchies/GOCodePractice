issue 1)
❌ Why this is not valid:

func solve(A []int , B int ) (int) {
    list := [len(A)]int{}
}
This gives a compile-time error:

this is valid:
list := [5]int{}

Because 5 is a constant known at compile time, and Go requires that array lengths be constants when defining arrays (not slices).

The Rule in Go:
Array sizes must be constant expressions known at compile time.

var x [10]int ✅ valid (10 is constant)

var x [len(A)]int ❌ invalid (len(A) is only known at runtime)
Go requires that array lengths be constants when defining arrays (not slices).



issue 2 ) with 2D array
Here's the core issue with your code:

var row, col int
row = len(matrix)
col = len(matrix[0])
var visited [row][col]bool // ❌ INVALID
❌ Why This Fails
In Go, array sizes must be compile-time constants, not variables.

But in your case:

row = len(matrix) and col = len(matrix[0]) are runtime values (you only know them when the program runs).

So visited [row][col]bool is not allowed — Go needs to know array sizes at compile time.

✅ Solutions
✅ Option 1: Use slices (preferred in Go for dynamic size)

visited := make([][]bool, row)
for i := range visited {
    visited[i] = make([]bool, col)
}
This is idiomatic Go — use make for dynamic 2D slices.

