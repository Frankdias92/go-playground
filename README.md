## Go Studies Repository

Welcome to my Go studies repository! This repository is structured by learning topics, each organized into a separate branch. The goal is to make it easier to navigate and reference code examples and explanations for each subject without overloading a single README file.

---

## **Repository Structure**  

| **Category**               | **Description**                                              | **Topic**         |
|----------------------------|-------------------------------------------------------------|----------------------------|
| **Base Concepts**          | Entry point of every Go application, the `main` function.    |`main`
|                            | Declaration and usage of variables in Go.                    |[`variables`](https://github.com/frankdias92/go-playground/tree/variables)
|                            | Data types and constants.                                    |[`typesAndConsts`](https://github.com/frankdias92/go-playground/tree/typesAndConsts)
|                            | Conditional structures like `if` and `else`.                 |[`if-condition`](https://github.com/frankdias92/go-playground/tree/if-condition)
|                            | Repetitive structures such as `for` loops.                   |[`loops`](https://github.com/frankdias92/go-playground/tree/loops)
| **Working with Data**      | Understanding and using arrays.                              |[`arrays`](https://github.com/frankdias92/go-playground/tree/arrays)
|                            | Differences and use cases for arrays and slices.             |[`arrays-and-slices`](https://github.com/frankdias92/go-playground/tree/arrays-and-slices)
|                            | Managing key-value pairs with maps.                          |[`map`](https://github.com/frankdias92/go-playground/tree/map)
| **Code Manipulation**      | Function definitions and usage in Go.                        |[`functions`](https://github.com/frankdias92/go-playground/tree/functions)
|                            | Using `defer` to schedule tasks.                             |[`defer`](https://github.com/frankdias92/go-playground/tree/defer)
|                            | Error handling and best practices.                          |[`errorhandler`](https://github.com/frankdias92/go-playground/tree/errohandler)
| **Advanced Structures**    | Creating and using structs in Go.                            |[`structs`](https://github.com/frankdias92/go-playground/tree/structs)
|                            | Managing visibility of identifiers in Go packages.           |[`private-and-puclic-names`](https://github.com/frankdias92/go-playground/tree/private-and-puclic-names)
|                            | Type parameters and generics.                               |[`type-parameter`](https://github.com/frankdias92/go-playground/tree/type-parameters)
| **Packages and Resources** | Organizing code into packages.                               |[`packages`](https://github.com/frankdias92/go-playground/tree/packages)
|                            | Working with pointers in Go.                                 |[`points`](https://github.com/frankdias92/go-playground/tree/points)
| **Concurrency**            | Managing concurrency and resource access.                    |[`competition-and-resources`](https://github.com/frankdias92/go-playground/tree/competition-and-resources)
| **Projects**               | A simple number guessing game project.                       |[`first-project-guessing-game`](https://github.com/frankdias92/go-playground/tree/first-project-guessing-game)
|                            | Interactive quiz project to reinforce learning.              |[`quiz-project`](https://github.com/frankdias92/go-playground/tree/quiz-project)
|                            | A small API project for a book catalog.                      |[`get-a-book`](https://github.com/frankdias92/go-playground/tree/get-a-book)
|                            | Building and structuring REST APIs.                          |[`api-rest`](https://github.com/frankdias92/go-playground/tree/api-rest)
|                            | Building an API shorten URL                                  |[`api-project`](https://github.com/frankdias92/go-playground/tree/api-project)
|                            | API in Go that allows you to search for movie informatio     |[`api-project-movies`](https://github.com/frankdias92/go-playground/tree/api-project-movies)
|                            | Multi Database integration showcasing the use of MySQL, PostgreSQL, SQLite, and Redis in Go, with Docker Compose for managing database containers. |[`api-project-redis`](https://github.com/frankdias92/go-playground/tree/multi-database-integration)
|                            | This project is the continuation of the api-project branch that I use memory storage. Here in this I migrate to the use of Redis as a backend system. |[`api-project-redis`](https://github.com/frankdias92/go-playground/tree/api-project-redis)
 
---

### How to Navigate the Repository

1. Click on one of the links in the table above to directly access the desired topic.
2. Each branch contains:
   - **Code Examples**: Practical demonstrations of the concept.
   - **Explanations**: Detailed and theoretical descriptions of the topic.
   - **Best Practices**: Tips to apply these concepts in real-world scenarios.

3. If you want to explore all the branches locally:
   ```bash
   # List all available branches
   git branch -r
   
   # switch a specific branch
   git switch <branch-name>

   ```

---

### Contributions and Improvements

If you have improvement suggestions, find any errors, or want to add new examples, feel free to open an issue or submit a pull request.

---

### About the Author

This repository was created as part of my continuous learning in Go. The structure is organized to help both myself and other developers studying the language. I hope you find it useful!

