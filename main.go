package main

import (
	"fmt"
	"myFirstProject/private_public"
)

func main() {
	fmt.Println("Hello, word!")

	fmt.Println(private_public.Foo)

	private_public.MyTest()
}
