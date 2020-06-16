# :small_red_triangle: Debugo

*An optimized debug library for go*

Debugo is a dead simple debugging library that doesn't increase your programs size thanks to the compiler optimizations.

## Installation

```
go get github.com/negrel/debugo
```

## Hello world

During development, use the assert statement — debugo.Assert(condition, panicMessage); — to disrupt normal execution if a boolean condition is false.

Take a look at the following example :

![example](https://github.com/negrel/debugo/raw/master/.github/carbon.png?sanitize=true)

Let's run the program:

```
go run .
'Hello world!' == 'Bye Bye' is false
panic: error, the given string are not equal

goroutine 1 [running]:
github.com/negrel/debugo.Assert(...)
        /home/negrel/code/golang/src/github.com/negrel/beta/vendor/github.com/negrel/debugo/assert_d.go:10
main.main()
        /home/negrel/code/golang/src/github.com/negrel/beta/main.go:17 +0x146
exit status 2
```

If we take a look at the **Static Single Assignment**, we will notice that the **debugo** function are removed by the compiler.


### Contributing
If you want to contribute to Ginger to add a feature or improve the code contact me at [negrel.dev@protonmail.com](mailto:negrel.dev@protonmail.com), open an [issue](https://github.com/negrel/debugo/issues) or make a [pull request](https://github.com/negrel/debugo/pulls).

## :stars: Show your support
Please give a :star: if this project helped you!

#### :scroll: License
MIT © Alexandre Negrel
