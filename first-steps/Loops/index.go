package loops

import (
	"fmt"
	"sync"
)

func SuperLoop() {
	var res int
	for i := 0; i < 10; i++ {
		res++
		fmt.Println(i)
	}
	fmt.Println("result", res)

	// for i := range 10 {
	// 	fmt.Println(i)
	// }

	// arr := [3]int{1, 2, 3}
	// for i, elem := range arr {
	// 	fmt.Println(&i, &elem)
	// }

	// str := "GoLang"
	// for i, ch := range str {
	// fmt.Printf("Index: %d, Character: %c\n", i, ch)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break // Exit the loop
	// 	}
	// 	if i%2 == 0 {
	// 		continue // Skip even numbers
	// 	}
	// 	fmt.Println(i)
	// }

}

func RangeBased() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range arr {
		fmt.Println(arr[i])
	}
	for _, elem := range arr {
		fmt.Println("Range Based", elem)
	}
}

func GoRoutinesLoop() {
	const n = 10
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("in Go routines", i)
		}()
	}
	wg.Wait()
}

func Matrix() {
	matrix := [2][2]int{{1, 2}, {3, 4}}
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Println("print matrix in range", matrix[i][j])
		}
	}
}
