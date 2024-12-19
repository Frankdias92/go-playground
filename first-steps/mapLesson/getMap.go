package maplesson

import (
	"fmt"
	"math"
)

func GetMap() {
	f := math.NaN()
	f2 := math.NaN()
	m := map[float64]string{
		f:  "Frank",
		f2: "Git",
	}

	fmt.Println(m)

	valor, ok := m[f]
	println(valor, ok)

	delete(m, f)
	fmt.Print(m)

	clear(m)
	fmt.Println(m)
}

func ForMap() {
	m := map[string]string{
		"Frank": "Person",
		"Git":   "Frank",
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println(m)

	for k := range m {
		if k == "Git" {
			delete(m, k)
		}
	}
	fmt.Println(m)
}
