### what is Packages?
Visibility Rules
Public (Exported) Names:

Start with a capital letter.
They are visible and accessible outside the package in which they were defined.

```go
// public
func Add(a, b int) int {
    return a + b
}
```

Private Names (Not exported):

Start with a lowercase letter.
Can only be accessed within the same package.

```go
// Private
func subtract(a, b int) int {
    return a - b
}
```


### Best Practice Principles
1. Encapsulation:
- Use private names to hide implementation details that should not be exposed.

2. Minimalism:
- Expose only what is necessary for the program to function.
This keeps your package's API clean and easy to use.

3. Documentation:
- Public names should be well documented so that other developers understand their usage.