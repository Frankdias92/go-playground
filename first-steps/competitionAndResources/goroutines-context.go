package competitionandresources

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func GoRoutines() {
	GetRequest()
	GetGoRoutine()
}

func GetRequest() {
	start := time.Now()

	for range 10 {
		resp, err := http.Get("https://google.com")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		fmt.Println("ok")
	}
	fmt.Println("Get a request", time.Since(start))
}

func GetGoRoutine() {
	start := time.Now()
	const n = 10
	var wg sync.WaitGroup
	wg.Add(n)

	for range n {
		go func() {
			defer wg.Done()
			resp, err := http.Get("https://google.com")
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			fmt.Println("ok")
		}()
	}

	wg.Wait()
	fmt.Println("Get a Go req with Goroutine", time.Since(start))
}
