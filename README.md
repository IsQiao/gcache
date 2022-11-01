gcache
========
[![GoDoc](https://godoc.org/github.com/maemual/go-cache?status.svg)](https://pkg.go.dev/github.com/IsQiao/gcache)

An in-memory K/V cache library. support generic!

## Documentation

[API Reference](https://pkg.go.dev/github.com/IsQiao/gcache)

## Installation

Install gcache using the "go get" command:

```
$ go get -u github.com/IsQiao/gcache
```

## Example

```
package main

import (
	"fmt"
	"time"

	"github.com/IsQiao/gcache"
)

type testItem struct {
	ColumnA string
}

func main() {
	c := gcache.NewDefault[testItem](time.Second)

	key1 := "key1"
	val1 := testItem{
		ColumnA: "val1",
	}

	key2 := "key2"
	val2 := testItem{
		ColumnA: "val2",
	}

	c.Set(key1, val1)
	c.Set(key2, val2)

	resultVal1 := c.Get(key1)
	fmt.Println(resultVal1.ColumnA)

	resultVal2 := c.Get(key2)
	fmt.Println(resultVal2.ColumnA)
}
```
