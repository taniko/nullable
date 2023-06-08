# nullable

## Installation
```sh
go get github.com/taniko/nullable
```

## Usage
```go
package main

import (
	"fmt"

	"github.com/taniko/nullable"
)

func main() {
	var v nullable.Nullable[string]
	fmt.Println(v.IsNull()) // true

	v = nullable.New("text")
	fmt.Println(v.IsNull()) // false
	fmt.Println(v.Value())  // text
}
```