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
