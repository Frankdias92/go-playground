### Type systems and constants in Go

```go
package main

import (
	"fmt"
)
```
- The `fmt` package is imported for formatted I/O operations (e.g., `Println`).

---

#### Type Declarations and Type Conversion

```go
// var x int = 10084
// f := float64(x)
// b := bool(x) // Cannot convert x (variable of type int into bool)
```
- **Explicit Conversion**: Go enforces explicit type conversions. For example, an `int` can be converted to `float64` but not to `bool`.

```go
// s := string(x)
```
- **String Conversion**: Direct conversion of integers to strings via `string(x)` results in the Unicode character corresponding to the integer value of `x`.
- Use `strconv.FormatInt` for meaningful numeric string conversions:
  ```go
  y := strconv.FormatInt(int64(x), 10) // Converts int64 to a decimal string.
  ```

---

#### Constants in Go

```go
const x = 10 // Untyped constant
```
- Untyped constants are flexible and can be used with various types where it makes sense.
- Typed constants (e.g., `const y int = 10`) are more restrictive and bound to a specific type.

---

#### Function Signatures

```go
func takeInt32(x int32) {
	fmt.Println(x)
}

func takeInt64(x int64) {
	fmt.Println(x)
}

func takeString(s string) {
	fmt.Println(s)
}
```
- Functions are defined to accept specific types, demonstrating how constants or literals can adapt to different parameter types when passed.


---

### Key Concepts to Study

1. **Typed vs. Untyped Constants**:
   - Untyped constants are more versatile, as they adapt to the type of the variable or function they’re used with.
   - Typed constants are explicitly bound to a type, ensuring strict type enforcement.

2. **Type Conversion**:
   - Always explicit in Go, ensuring type safety.
   - Use packages like `strconv` for non-trivial conversions (e.g., numbers to strings).

3. **Type Systems**:
   - Strongly and statically typed, avoiding common runtime errors.

4. **Literal Types**:
   - Constants can represent literals, such as numeric or string values, used flexibly throughout code.
