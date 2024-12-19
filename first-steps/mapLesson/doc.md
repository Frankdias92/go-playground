## Lesson: Maps in Go

### Overview

In this lesson, we learned about maps, which are fundamental data structures in the Go language. Key points covered include:

- A map must be initialized before inserting elements, which can be done using a map literal or the `make` function.
- We explored how to access and delete keys from a map and how to iterate over a map.
- It was highlighted that in Go, it is possible to delete keys during iteration, which is not common in other implementations of hash maps.
- The lesson emphasized that slices are passed by value but contain a pointer to the data, while arrays are passed by value, copying all elements.

### Code Example

Here is the provided code with explanations:

```go
package maplesson

import (
	"fmt"
	"math"
)

// Function to demonstrate map creation and manipulation
func GetMap() {
	f := math.NaN() // Create NaN values
	f2 := math.NaN()
	m := map[float64]string{
		f:  "Frank",
		f2: "Git",
	}

	fmt.Println(m) // Print the map

	valor, ok := m[f] // Accessing value with NaN key
	println(valor, ok) // Prints the value and existence status

	delete(m, f) // Delete key f from the map
	fmt.Print(m) // Print the map after deletion

	clear(m) // This will cause an error since clear is not a valid function for maps
	fmt.Println(m)
}

// Function to demonstrate iterating over a map
func ForMap() {
	m := map[string]string{
		"Frank": "Person",
		"Git":   "Frank",
	}

	for k, v := range m {
		fmt.Println(k, v) // Print key-value pairs
	}
	fmt.Println(m)

	for k := range m {
		if k == "Git" {
			delete(m, k) // Delete key "Git" during iteration
		}
	}
	fmt.Println(m)
}

// Note: The 'clear' function does not exist for maps in Go and will cause an error if called.
```

### Explanation of the Code

#### Function `GetMap`
- Two NaN values are created using `math.NaN()`.
- A map `m` is initialized with float64 keys and string values. Since NaN is not equal to itself in floating-point representation, both keys will actually refer to the same entry in the map.
- The value associated with one of the NaN keys is accessed and printed. The `ok` variable indicates whether the key exists in the map.
- The key `f` is deleted from the map using `delete()`.
- An attempt is made to call `clear(m)`, which will result in an error because there is no built-in `clear` function for maps in Go.

#### Function `ForMap`
- A string-to-string map is created with two entries.
- The first loop iterates over the map and prints each key-value pair.
- The second loop demonstrates deleting a key during iteration. In this case, if the key `"Git"` is found, it is deleted from the map.

### Important Considerations

1. **Initialization**: Always initialize a map before adding elements. You can use:
   - `make(map[KeyType]ValueType)` for dynamic creation.
   - Map literals for initialization with predefined values.

2. **Accessing Values**: When accessing values from a map, use the comma-ok idiom to check if a key exists:
   ```go
   value, ok := m[key]
   ```

3. **Deleting Keys**: You can safely delete keys from a map while iterating through it. However, be cautious as this can lead to unexpected behavior if not handled properly.

4. **Nil Maps**: A nil map behaves like an empty map when reading but causes a panic when writing. Always initialize your maps before use.

5. **Performance**: Maps provide average-case constant time complexity for lookups, inserts, and deletes due to their underlying hash table implementation.
