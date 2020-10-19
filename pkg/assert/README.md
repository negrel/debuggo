# assert

<p align="center">
	<a href="https://pkg.go.dev/github.com/negrel/debuggo/pkg/assert">
		<img alt="PkgGoDev" src="https://pkg.go.dev/badge/github.com/negrel/debuggo">
	</a>
	<a href="https://goreportcard.com/report/github.com/negrel/debuggo">
		<img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/negrel/debuggo/pkg/assert">
	</a>
</p>

The **assert** package provide many helpful functions to perform one line assertions.

### Features

- Can be turned off (as all **debuggo** packages)
- Prints friendly, easy to read failure descriptions

- Allows for very readable code
- Optionally annotate each assertion with a message



## Example

*See [`examples/assert`](https://github.com/negrel/debuggo/blob/master/examples/assert/main.go)*

```go
package main

import (
	"errors"
	"fmt"

	"github.com/negrel/debuggo/pkg/assert"
)

func main() {
	days, err := cureCOVID()
	assert.Nil(err, err)

	fmt.Printf("it took %v days to cure the COVID.\n", days)
}
```



## Why

The `assert` package is a **modified** version of the [`testify/assert`](https://github.com/stretchr/testify) package (see `codegen/assert/main`). The code is generated using the tools from the  [`asttk`](https://github.com/negrel/asttk) module.

This package is **inspired** by the [*assert*](https://dart.dev/guides/language/language-tour#assert) function of the [Dart language](https://dart.dev/):

> During development, use an assert statement — `assert(condition, optionalMessage)`; — to disrupt normal execution if a
> boolean condition is false.

> In production code, assertions are ignored, and the arguments to `assert` aren’t evaluated.



The `assert` package is an attempt to reproduce the *assert* function for **Go**