# foobar

A sample Go library.

## Installation

```bash
go get github.com/cubahno/foobar
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/cubahno/foobar"
)

func main() {
    // Basic greeting
    greeting := foobar.Foo("Alice")
    fmt.Println(greeting) // Output: Hello, Alice!

    // Transform to uppercase
    upper := foobar.Bar("hello world")
    fmt.Println(upper) // Output: HELLO WORLD

    // Combined operation
    result := foobar.FooBar("Bob")
    fmt.Println(result) // Output: HELLO, BOB!
}
```

## License

MIT

