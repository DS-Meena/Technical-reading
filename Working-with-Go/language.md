# How to Use Golang

Capture in Notion. Then organize here since related to this project. 

# Few confusing things 😕

### Function Literals (Anonymous Functions) 🚯

Function literals are functions without names that can be assigned to variables or passed as arguments. They can access variables from the enclosing function (closure).

```go
// Example of a function literal
square := func(x int) int {
    return x * x
}
result := square(5) // Returns 25
```

### Closures 🔐

Closures are function literals that can reference variables from their enclosing scope, even after the outer function returns.

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
// count maintains its value between calls
```

### Serial and Parallel Consistency 🏗️

Serial consistency (also known as sequential consistency) means that operations appear to execute in the order specified in the program, while parallel consistency deals with how concurrent operations are ordered across multiple goroutines.

**Serial Consistency ⛎**

- Operations within a single goroutine are always serially consistent
- The program executes instructions in the order they appear in the code
- Reads and writes in the same goroutine maintain their program order

**Parallel Consistency 🈂️**

- Deals with ordering of operations across multiple concurrent goroutines
- Go provides sync primitives (mutexes, channels) to maintain consistency
- Without proper synchronization, parallel operations may have race conditions

Example using channels for consistency (parallel):

```go
ch := make(chan int)
go func() {
    // Operations in goroutine 1
    x := 42
    ch <- x // Ensures operation ordering
}()
// Operations in main goroutine
val := <-ch // Guarantees consistent view of x
```

Example using mutex for consistency:

```go
var mu sync.Mutex
var count int

func increment() {
    mu.Lock()
    count++ // Protected access ensures consistency
    mu.Unlock()
}
```

This is how you can handle concurrent operations to achieve parallel consistency. 