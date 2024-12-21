Loops in Go are essential for iterative operations. Go simplifies looping by using just one keyword: `for`. Unlike other languages, Go doesn't have constructs like `while` or `do-while`; instead, `for` can handle all looping scenarios.

---

### Key Features of Loops in Go

1. **Single `for` Keyword**:
   - Go uses `for` for all types of loops, including infinite loops, condition-based loops, and range-based loops.

2. **Simpler Syntax**:
   - No parentheses are required around conditions or initialization; just use braces `{}` to define the block.

---

### Examples and Insights from Your Code

#### 1. **Simple `for` Loop**
   ```go
   for i := 0; i < 10; i++ {
       fmt.Println(i)
   }
   ```
   - **Initialization**: `i := 0`
   - **Condition**: `i < 10`
   - **Increment**: `i++`
   - Output: Prints numbers 0 to 9.

---

#### 2. **Range-based `for` Loop**
   ```go
   arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
   for i := range arr {
       fmt.Println(arr[i]) // Access element using index
   }
   ```
   - `range` provides a convenient way to iterate over arrays, slices, maps, strings, and channels.
   - **Index and Element**:
     ```go
     for i, elem := range arr {
         fmt.Println(i, elem) // Prints index and value
     }
     ```

---

#### 3. **Ignoring Values**
   - Use `_` to ignore the index or value:
     ```go
     for _, elem := range arr {
         fmt.Println(elem) // Ignores the index
     }
     ```

---

#### 4. **Infinite Loop**
   - Use `for` without a condition to create an infinite loop:
     ```go
     for {
         fmt.Println("Running forever")
     }
     ```

---

#### 5. **Using Goroutines with Loops**
   - In your example:
     ```go
     const n = 10
     var wg sync.WaitGroup
     wg.Add(10)

     for i := 0; i < n; i++ {
         go func() {
             defer wg.Done()
             fmt.Println(i)
         }()
     }
     wg.Wait()
     ```
   - **Issue**: This introduces a common **closure problem** in Go. Inside the goroutine, `i` is shared across iterations, and the output may not be as expected.
   - **Fix**: Pass `i` as an argument to the goroutine:
     ```go
     go version < 1.22
     for i := 0; i < n; i++ {
         go func(i int) {
             defer wg.Done()
             fmt.Println(i)
         }(i)
     }
     ```
   - This ensures each goroutine gets its own copy of `i`.

---

#### 6. **Loop Over Strings**
   - Use `range` to iterate over strings:
     ```go
     str := "GoLang"
     for i, ch := range str {
         fmt.Printf("Index: %d, Character: %c\n", i, ch)
     }
     ```

---

#### 7. **Breaking or Skipping Iterations**
   - Use `break` to exit a loop and `continue` to skip the current iteration:
     ```go
     for i := 0; i < 10; i++ {
         if i == 5 {
             break // Exit the loop
         }
         if i%2 == 0 {
             continue // Skip even numbers
         }
         fmt.Println(i)
     }
     ```

---

#### 8. **Nested Loops**
   - Use nested loops for multidimensional arrays or complex conditions:
     ```go
     matrix := [2][2]int{{1, 2}, {3, 4}}
     for i := range matrix {
         for j := range matrix[i] {
             fmt.Println(matrix[i][j])
         }
     }
     ```

---

### Common Mistakes and Best Practices

1. **Avoid Infinite Loops Unintentionally**:
   - Ensure loop conditions are correct to prevent infinite loops.
   - Example:
     ```go
     for i := 0; i < 10; i-- { // Infinite loop due to decrement
         fmt.Println(i)
     }
     ```

2. **Closure in Goroutines**:
   - Always pass variables to goroutines explicitly to avoid unintended behaviors.

3. **Efficient Iteration**:
   - Use `range` for concise and readable iterations, especially with collections.
