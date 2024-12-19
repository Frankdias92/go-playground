package competitionandresources

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"time"
)

func GetContextRequest() {
	start := time.Now()
	const n = 10
	var wg sync.WaitGroup
	wg.Add(n)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second) // Set a timeout of 5 seconds
	defer cancel()                                         // Ensure that the cancel function is called to release resources

	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(10 * time.Second)
			fmt.Fprintln(w, "hello, world!")
		}),
	)

	for range n {
		go func(ctx context.Context) { // Start a goroutine for each request

			defer wg.Done() // Signal that this goroutine is done when it finishes

			/* handle with the request */
			req, err := http.NewRequestWithContext(
				ctx,        // context
				"GET",      // Method
				server.URL, // URL
				nil,        // body
			)
			if err != nil {
				panic(err)
			}

			/* handle with the response */
			resp, err := http.DefaultClient.Do(req) // send the request
			if err != nil {
				if errors.Is(err, context.DeadlineExceeded) {
					fmt.Println("timeout")
					return
				}
				panic(err)
			}
			defer resp.Body.Close() // Close response body after reading (not shown in this example)
		}(ctx)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Waint for server", time.Since(start))
}
