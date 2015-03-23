# funky

[![Build Status](https://travis-ci.org/nyrkovalex/funky.svg)](https://travis-ci.org/nyrkovalex/funky)

Funky provides simple functional idioms for golang builtin types by thin and simple wrappers.

# Usage

## Chan
Every method returns new `funky.Chan` that can be used to chain calls
```go
import (
    "github.com/nyrkovalex/funky"
    "fmt"
)

func main() {
    c := funky.ChanOf(funky.Slice{1, 2, 3})
    result := c.Filter(func(item interface{}) bool {
        return item.(int) != 2
    }).Map(func(item interface{}) interface{} {
        return fmt.Sprintf("%d", item.(int))
    }).Reduce(func(left interface{}, right interface{}) interface{} {
        return left.(string) + right.(string)
    })
    fmt.Println(result) // Outputs "13"
}
```

## Slice
*Slice should not probaby be used anymore as `funky.Chan` should perform much better*

Every method returns a new object without modifying the source.
```go
import (
    "github.com/nyrkovalex/funky"
    "fmt"
)

func main() {
    fs := funky.Slice{1, 2, 3}

    // Returns funky.Slice{"1 funky", "2 funky", "3 funky"}
    fs.Map(func(item interface{}) interface{} {
        return fmt.Sprintf("%d funky", item.(int))
    })

    // Returns funky.Slice{1}
    fs.Filter(func(item interface{}) bool {
        return item.(int) < 2
    })

    fs.Contains(2) // Returns true

    // Returns funky.Slice{1, 2, 3, 4, 5}
    fs.Append(4, 5)

    // Returns funky.Slice(1, 3}
    fs.Delete(1)

    // Returns 6
    fs.Reduce(func(left interface{}, right interface{}) interface{} {
        return left.(int) + right.(int)
    })
}
```

