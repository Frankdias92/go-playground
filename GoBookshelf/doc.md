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

### Create a new method to get book
```
package service

type Book struct {
	ID     int
	Title  string
	Author string
	Genre  string
}

func (b Book) GetFullBook() string {
	return b.Title + " by " + b.Author
}
```
### Execute this method on the main
```
package main

import (
	"errors"
	"fmt"
	"gobooks/internal/service"
)

func main() {
	book := service.Book{
		ID:     1,
		Title:  "New title",
		Author: "Fredrik",
		Genre:  "Action",
	}

	fmt.Println(book.GetFullBook())
}
```
output: New title by Fredrik