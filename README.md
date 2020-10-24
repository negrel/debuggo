# :small_red_triangle: Debuggo

<p>
	<a href="https://pkg.go.dev/github.com/negrel/debuggo">
		<img alt="PkgGoDev" src="https://pkg.go.dev/badge/github.com/negrel/debuggo">
	</a>
	<a href="https://goreportcard.com/report/github.com/negrel/debuggo">
		<img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/negrel/debuggo">
	</a>
</p>

*Debugging packages that leverage the power of conditional compilation.*

**Debuggo** made of the following packages:
- [`assert`](#the-assert-package)
- [`log`](#the-log-package)

## Why ?

Recently, I was looking at the source code of [flutter](https://flutter.dev/) to have a deeper understanding of what's
going on inside. This is where I discovered the assert function of dart, *assert* is used everywhere in
[flutter](https://flutter.dev/), every single **class** use it at least once.

*Definition of assert from [dart.dev](https://dart.dev/guides/language/language-tour#assert)* :

> During development, use an assert statement — `assert(condition, optionalMessage)`; — to disrupt normal execution if a
> boolean condition is false.

> In production code, assertions are ignored, and the arguments to `assert` aren’t evaluated.

This function is pretty basic, but it allows developers to make assertions in classes methods without adding overhead
to production binary. Take a look at this part of the definition :

>In **production** code, assertions are **ignored**, and the arguments to `assert` aren’t **evaluated**.

Assertions are removed for the production, thereby, we can write development code without **increasing** the size or
**slowing down** our production binaries.

This is exactly what we tried to reproduce in `pkg/assert`:

## The `assert` package
The `assert` package is a modified version of the excellent [`testify/assert`](https://github.com/stretchr/testify)
package. Thus, you can use assert functions outside tests.

### Features
- Prints friendly, easy to read failure descriptions  
- Allows for very readable code
- Optionally annotate each assertion with a message

Take a look at [examples/assert/main.go](https://github.com/negrel/debuggo/blob/master/examples/assert/main.go). By
default, assertions are disable.

```bash
# Let's run the example
$ go run .
In one hour the temperature will be -1 °C



# The program didn't panic because assertions were disabled
# Let's run the exampe with assertions enabled now.
$ go run -tags assert .
panic: 
%s      Error Trace:    proc.go:204
                                                asm_amd64.s:1374
        Error:          Expected nil, but got: &errors.errorString{s:"unable to get the weather forecast"}
        Messages:       []


goroutine 1 [running]:
github.com/negrel/debuggo/pkg/assert.debuggoGen_Fail(0xc000018360, 0x52, 0xc00014bf38, 0x1, 0x1, 0xc000018360)
        /home/negrel/code/golang/src/github.com/negrel/debuggo/pkg/assert/assertions.go:1071 +0x2a7
github.com/negrel/debuggo/pkg/assert.debuggoGen_Nil(0x582c00, 0xc000012d90, 0xc00014bf38, 0x1, 0x1, 0x5e9260)
        /home/negrel/code/golang/src/github.com/negrel/debuggo/pkg/assert/assertions.go:1179 +0xdb
github.com/negrel/debuggo/pkg/assert.Nil(...)
        /home/negrel/code/golang/src/github.com/negrel/debuggo/pkg/assert/assertions.go:353
main.main()
        /home/negrel/code/golang/src/github.com/negrel/debuggo/examples/assert/main.go:15 +0xda
exit status 2

```

Okay, assertions are great for debugging but logging can also be useful.

## The `log` package
This package is a wrapper around the `log` package of the standard lib. There are 7 log levels, and you can choose one of 
the following using a build tag:
- Panic
- Fatal
- Error
- Warn
- Info
- Debug
- Trace

### Contributing
If you want to contribute to Debuggo to add a feature or improve the code contact me at
[negrel.dev@protonmail.com](mailto:negrel.dev@protonmail.com), open an [issue](https://github.com/negrel/debuggo/issues)
or make a [pull request](https://github.com/negrel/debuggo/pulls).

## :stars: Show your support
Please give a :star: if this project helped you!

#### :scroll: License
GPLv3 © [Alexandre Negrel](https://www.negrel.dev)
