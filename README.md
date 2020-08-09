# :small_red_triangle: Debuggo

*Debuggo generate optimized debugging packages for minimal performance / size overhead in production.*

## Installation

```
# Install the CLI tool
go install github.com/negrel/debuggo/cmd/debuggo
```

## Why ?

Recently, I was looking at the source code of [flutter](https://flutter.dev/) to have a deeper understanding of what's
going on inside. In the codebase, I discovered the `assert` function of dart, `assert` is used everywhere in
flutter, every single **class** use it at least once.

*Definition of assert from [dart.(https://flutter.dev/)dev](https://dart.dev/guides/language/language-tour#assert)* :

> During development, use an assert statement — `assert(condition, optionalMessage)`; — to disrupt normal execution if a
>boolean condition is false.
>
> In production code, assertions are ignored, and the arguments to `assert` aren’t evaluated.

This function is pretty basic but and allows developer to make assertions in classes methods but without adding overhead
to production binary. Take a look at this part of the definition :

>In **production** code, assertions are **ignored**, and the arguments to `assert` aren’t **evaluated**.

Call to `assert` are removed for the production, thereby, we can write development code without **increasing** the size
or **slowing down** our production binaries.

### Introducing Debuggo

Debuggo rely on [**compiler optimization**](https://en.wikipedia.org/wiki/Optimizing_compiler) and
[**conditional compilation**](https://en.wikipedia.org/wiki/Conditional_compilation) to generate production code for
your debugging packages.

Let's write our own `assert` function now:

### Assert package
The following example is inspired by the [**`assert`**](https://dart.dev/guides/language/language-tour#assert) function
from the Dart language:

```go
package assert

// True assert that the given bool is true otherwise,
// the program will panic.
func True(isTrue bool, optionalMsg string) {
  if !isTrue {
    panic(optionalMsg)
  }
}
```

Then we can generate a new package with the following command:
```bash
# The ./debuggo/src is the source directory for the generator
# The ./debug/ is the output directory of the generator 
$ debuggo gen --src-dir ./debuggo/src/ --out-dir ./debug/
``` 

or using `go generate` with a file containing the following comment:
```go
// gen.go
package your_package

//go:generate debuggo gen --src-dir ./debuggo/src/ --out-dir ./debug/
```

Once the package is generated you can start using it:

```go
// main.go
package main

// Path to your generated package.
import (
  "fmt"
"path/filepath"

  "github.com/negrel/beta/debug/assert"
)

func main() {
  path := getPath()
  assert.True(filepath.IsAbs(path),
    fmt.Sprintf("%v should be an absolute path", path),
  )
  // ...
}
```

If we run the program written above, the `assert.True` call will be removed and the program will not panic. If we want
to enable assertions of the `assert` package, we must use the `assert` build tag:
```bash
# Run the program
$ go run .
# Run the program with assertions enabled
$ go run -tags assert .
```
The **name** of your **debugging package** define the build tag to enable it. 


## TODO

- [ ] Custom function for production build.
- [ ] Placeholder variable to customize package (e.g `debuggo.PkgName`)
- [ ] Detect source file change to regenerate package

### Contributing
If you want to contribute to Debuggo to add a feature or improve the code contact me at
[negrel.dev@protonmail.com](mailto:negrel.dev@protonmail.com), open an [issue](https://github.com/negrel/debuggo/issues)
or make a [pull request](https://github.com/negrel/debuggo/pulls).

## :stars: Show your support
Please give a :star: if this project helped you!

#### :scroll: License
MIT © [Alexandre Negrel](https://www.negrel.dev)
