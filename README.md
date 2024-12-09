### what is Packages?
A package is a collection of Go code files that together form a module or feature. It is identified by the name declared at the top of a .go file.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

**In the example above**:

The package is main, which is the entry point for running Go programs.
The code uses the fmt package, which is Go's standard library for formatted input/output.
Types of Packages
Standard Packages:

Go provides several standard packages out of the box.
Examples:
fmt: For formatted input and output.
os: For interacting with the operating system.
math: For advanced mathematical calculations.
net/http: For creating and consuming HTTP services.

