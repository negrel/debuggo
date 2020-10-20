# :small_red_triangle: Debuggo


<p align="center">
	<a href="https://pkg.go.dev/github.com/negrel/debuggo">
		<img alt="PkgGoDev" src="https://pkg.go.dev/badge/github.com/negrel/debuggo">
	</a>
	<a href="https://goreportcard.com/report/github.com/negrel/debuggo">
		<img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/negrel/debuggo">
	</a>
</p>

*Debugging package that leverage the power of conditional compilation.*

**Debuggo** is a module made of the following *switchable* package for debugging:
- `assert`
- `log`

We say "switchable" because you can activate/deactivate them using build tags. 

## Why

Recently, I was looking at the source code of [flutter](https://flutter.dev/) to have a deeper understanding of what's
going on inside. This is where I discovered the assert function of dart, *assert* is used everywhere in
[flutter](https://flutter.dev/), every single **class** use it at least once.

*Definition of assert from [dart.dev](https://dart.dev/guides/language/language-tour#assert)* :

> During development, use an assert statement — `assert(condition, optionalMessage)`; — to disrupt normal execution if a
> boolean condition is false.

> In production code, assertions are ignored, and the arguments to `assert` aren’t evaluated.

This function is pretty basic but and allows developer to make assertions in classes methods but without adding overhead
to production binary. Take a look at this part of the definition :

>In **production** code, assertions are **ignored**, and the arguments to `assert` aren’t **evaluated**.

Assertions are removed for the production, thereby, we can write development code without **increasing** the size or
**slowing down** our production binaries.



This is exactly what we tried to reproduce in `pkg/assert`:

## The assert package
The assert package is modified version of the excellent [`testify/assert`](https://github.com/stretchr/testify) package.
Thus you can use assert functions outside tests using the `assert` build tag. 

- Prints friendly, easy to read failure descriptions  

- Allows for very readable code
- Optionally annotate each assertion with a message

### Contributing
If you want to contribute to Debuggo to add a feature or improve the code contact me at
[negrel.dev@protonmail.com](mailto:negrel.dev@protonmail.com), open an [issue](https://github.com/negrel/debuggo/issues)
or make a [pull request](https://github.com/negrel/debuggo/pulls).

## :stars: Show your support
Please give a :star: if this project helped you!

#### :scroll: License
GPLv3 © [Alexandre Negrel](https://www.negrel.dev)
