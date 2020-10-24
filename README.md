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

Each log level have 4 method, for example the `info` level have the following:
- Info("my message")
- Infoln("my message")
- Infof("My message is %v:", "hello world")
- Infofn(myfunction)

Take a look at [examples/log/main.go](https://github.com/negrel/debuggo/blob/master/examples/log/main.go). By
default, logs are disable.

```bash
# Let's run the example
$ go run .

# Nothing happend, let's try using the `info` build tag.
$ go run -tags info .
2020/10/24 10:19:59 [INFO] - Info log
2020/10/24 10:19:59 [WARN] - Warning log
2020/10/24 10:19:59 [ERROR] - Error log
2020/10/24 10:19:59 [FATAL] - Fatal log
exit status 1
```

## Analysing binaries

If you have taken a look at the generated code in pkg/* you must have notice that many functions are empty. Thus, we may
ask ourselves if there is an overhead to use Debuggo packages ?

### Disassemble binaries

To answer this question, we should take a look at the produced binaries:
```bash
# Go to an example
$ cd examples/log/

# Let's build the example with the `info` build tag
$ go build -tags info .

# Let's disassemble the file.
$ objdump log -d > log_info.asm

# Number of line
$ wc -l log_info.asm
182342 log_info.asm

# Let's build the example without any build tag
$ go build .

# Let's disassemble the file.
$ objdump log -d > log.asm

# Number of line
$ wc -l log.asm
149495 log.asm
``` 

There is a difference of **32847** lines between the two disassembled files. This show us that the produced binaries are
different but that's not enough to prove that all call to `log` have been removed.

### Static Single Assignment (SSA)

In order to get a better understanding of the produced binaries we're going to analyse generated SSA:

```bash
# So, we're still in the examples/log directory
# Let's generate the SSA with the `info` build tag 
$ GOSSAFUNC=main go build -tags info .

# An ssa.html file have been generated
$ mv ssa.html ssa_info.html

# Let's generate the SSA without any build tag 
$ GOSSAFUNC=main go build .
```

Now tht we have two HTML files containing the compilation phases of our example, let's analyse them.

Let's start with `ssa.html`:  
Since we only care about the result, we should analyse the last phase (*genssa*):
```
# /home/negrel/code/golang/src/github.com/negrel/debuggo/examples/log/main.go
00000 (7) TEXT "".main(SB), ABIInternal
00001 (7) FUNCDATA $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
00002 (7) FUNCDATA $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
b22
00003 (+15) RET
00004 (?) END
```
This is the expected result since without any build flag, the main function is empty.

Let's compare with the `ssa_info.html`:
```
# /home/negrel/code/golang/src/github.com/negrel/debuggo/examples/log/main.go
00000 (7) TEXT "".main(SB), ABIInternal
...
00014 (10) PCDATA $1, $0
v70
00015 (10) CALL github.com/negrel/debuggo/pkg/log.Infoln(SB)
v120
00016 (+11) XORPS X0, X0
...
v92
00026 (11) CALL github.com/negrel/debuggo/pkg/log.Warnln(SB)
...
00037 (12) CALL github.com/negrel/debuggo/pkg/log.Errorln(SB)
...
00048 (13) CALL github.com/negrel/debuggo/pkg/log.Fatalln(SB)
...
# /home/negrel/code/golang/src/github.com/negrel/debuggo/pkg/log/exported.info.go
v159
00055 (+24) MOVQ github.com/negrel/debuggo/pkg/log.std(SB), AX
v157
...
b19
00065 (+24) RET
00066 (?) END
```

This result also match our expectations since we have call to `log` functions. You may have notice that only call to
the available functions at the `info` level are present.

### Contributing
If you want to contribute to Debuggo to add a feature or improve the code contact me at
[negrel.dev@protonmail.com](mailto:negrel.dev@protonmail.com), open an [issue](https://github.com/negrel/debuggo/issues)
or make a [pull request](https://github.com/negrel/debuggo/pulls).

## :stars: Show your support
Please give a :star: if this project helped you!

#### :scroll: License
GPLv3 © [Alexandre Negrel](https://www.negrel.dev)
