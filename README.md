# :small_red_triangle: Debuggo

![travis-ci status](https://travis-ci.org/negrel/debuggo.svg?branch=master)
[![codebeat badge](https://codebeat.co/badges/9ae45e94-287c-4fcd-9434-0f403657023e)](https://codebeat.co/projects/github-com-negrel-debuggo-master)

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

*Definition of assert from [dart.dev](https://dart.dev/guides/language/language-tour#assert)* :

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

### Simple assert package
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
debuggo gen --src-dir ./debuggo/src/ --out-dir ./debug/
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
  relativePath := getPath()
  assert.True(filepath.IsAbs(relativePath),
    fmt.Sprintf("%v should be an absolute path", relativePath),
  )
  // ...
}
```

Now, if we run the program, even if the path is relative, the program will **not** panic. The reason is that we need to 
**enable** assertions, to do this we need to add a **build tag** when we **run/compile** the program:

```bash
# `assert` build tag
#        ┌--┴--┐
#        |     |
#        |     |
#        v     v
go run -tags assert .

panic: ./your/path/ should be an absolute path

goroutine 1 [running]:
github.com/negrel/beta/assert.True(...)
	/home/negrel/code/golang/src/github.com/negrel/beta/assert/assert.go:9
main.main()
	/home/negrel/code/golang/src/github.com/negrel/beta/main.go:10 +0x87
exit status 2
```
As you may have notice, the build tag is the name of your package.

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
