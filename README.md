### Study Structure with Improvements and Explanations

1. **Variable Declaration in Go**:
   - **Explicit Type**:
     ```go
     var x int8 = 10
     ```
     Here, the type (`int8`) is explicitly stated, ensuring clarity.

   - **Type Inference**:
     ```go
     year := "1992"
     ```
     The type is automatically inferred based on the assigned value. This method is concise but depends on the clarity of the assigned value.

   - **Declaration Without Initialization**:
     ```go
     var age int
     ```
     The variable is initialized with its type's default value (`0` for integers, `""` for strings, etc.).

2. **Grouping Variable Declarations**:
   - Grouping makes the code more organized and readable:
     ```go
     var (
         name     = "Franklin"
         lastName = "Macedo"
     )
     ```

3. **Global vs. Local Variables**:
   - `age` is a **global variable**, accessible throughout the package.
   - `x`, `year`, `name`, and `lastName` are **local variables**, available only within the `main` function.
