### **What Are Functions in Go?**

In Go, functions are reusable blocks of code designed to perform specific tasks. They can:
- Accept parameters as input.
- Return values as output.
- Be named or anonymous.
- Support features like variadic parameters and higher-order functions.

---

### **Basic Structure of a Function**
```go
func functionName(parameter1 type, parameter2 type) returnType {
    // function body
    return value
}
```

- **`func`**: Declares a function.
- **`functionName`**: Name of the function.
- **Parameters**: Function inputs.
- **`returnType`**: Type of the return value (optional).
- **Function Body**: The block of code that implements the logic.


### **Explaining Specific Functions**

1. **`sums` Function** (Variadic):
   - Accepts a variable number of `int` arguments.
   - Internally, the arguments are treated as a **slice**.
   - Allows flexibility by enabling calls with different numbers of arguments.
   ```go
   sums(10, 20, 30) // Adds 10 + 20 + 30
   sums(5, 15)      // Adds 5 + 15
   ```

2. **`swap` Function**:
   - Demonstrates multiple return values.
   - Swaps the values of `a` and `b`.
   ```go
   a, b := swap(1, 2) // a = 2, b = 1
   ```

3. **`divid` Function**:
   - Names the return values for better code clarity.
   - Calculates the quotient and remainder of a division.
   ```go
   res, rem := divid(10, 3) // res = 3 (quotient), rem = 1 (remainder)
   ```

4. **`sumHighOrder` Function**:
   - A **higher-order function** that returns another function.
   - Enables "partial application" by fixing certain values.
   ```go
   addTwo := sumHighOrder(2) // Returns a function that adds 2 to a number.
   result := addTwo(3)       // result = 5 (2 + 3)
   ```

---

### **Concepts Demonstrated in the Code**
- **Variadic Functions**: Allow dynamic arguments (`sums`).
- **Multiple Return Values**: Return more than one value (`swap`, `divid`).
- **Higher-Order Functions**: Return or accept other functions (`sumHighOrder`).
- **Local Scope**: Variables defined inside a function are local and don’t affect the rest of the program.
