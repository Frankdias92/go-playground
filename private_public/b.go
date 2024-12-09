package private_public

import (
	"fmt"
	"myFirstProject/private_public/internal/foo"
)

var Bar string = "hello, bar"

// Method exported
func MyTest() {
	fmt.Println(foo.MyTest)
}

// Method not exported
// func notExport() {
// 	fmt.Println("Cant export this function", foo.MyTest)
// }
