package arrays

import "fmt"

// Main function that demonstrates the creation and modification of slices
func GetArrayAndSlice() {
	slice := []int{1, 2, 3} // Creates an initial slice
	fmt.Println("Original Slice:", slice)

	slice = append(slice, 4) // Adds an element to the slice
	fmt.Println("Slice after append:", slice)

	// Multiplies each element by two
	for i := range slice {
		slice[i] *= 2
	}
	fmt.Println("Modified Slice:", slice)

	// Creation of a matrix (slice of slices)
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("matrix:", matrix)
}

// Function that demonstrates out-of-bounds access behavior
func IsInbounds(slice []int) {
	_ = slice[3] // Out-of-bounds access to check bounds check
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
	fmt.Println(slice[3]) // This will cause a panic if there is no element at index 3
}

// To run with bounds check:
// => go run -gcflags=!-d=ssa/check_bce
