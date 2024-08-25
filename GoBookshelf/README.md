_my first code on Go!_
```
package main

import (
	"errors"
	"fmt"
)

func sum(x, y int) (int, error) {
	if x+y > 5 {
		return x + y, nil
	}
	return 0, errors.New("x + y is less than 5")
}

func main() {
	x, err := sum(4, 3)
	if err != nil {
		panic(err)
	}

	fmt.Println(x)
}
```
Function: Return the sum of two numbers.<br>
Output: 7