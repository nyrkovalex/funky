# funky
Funky provides simple functional idioms for golang builtin types by thing and simple wrappers.

# Usage
Every method returns a new object without modifying the source.
```go
import (
    "github.com/nyrkovalex/funky"
    "fmt"
)

goFunky := funky.Slice{1, 2, 3}

// Returns funky.Slice{"1 funky", "2 funky", "3 funky"}
goFunky.Map(func(item interface{}) interface{} {
    return fmt.Sprintf("%d funky", item.(int))
})

// Returns funky.Slice{1}
goFunky.Filter(func(item interface{}) bool {
    return item.(int) < 2
})

goFunky.Contains(2) // Returns true

// Returns funky.Slice{1, 2, 3, 4, 5}
goFunky.Append(4, 5)

// Returns funky.Slice(1, 3}
goFunky.Delete(1)

// Returns 6
goFunky.Reduce(func(left interface{}, right interface{}) interface{} {
    return left.(int) + right.(int)
})
```
