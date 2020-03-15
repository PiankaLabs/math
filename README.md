<img src="https://www.piankalabs.com/static/media/piankalabs_title.svg" alt="Hugo" width="565"><br/>

[![Build Status](https://travis-ci.com/piankalabs/math.svg?branch=master)](https://travis-ci.org/piankalabs/math)
[![codecov](https://codecov.io/gh/piankalabs/math/branch/master/graph/badge.svg)](https://codecov.io/gh/piankalabs/math)
[![codebeat badge](https://codebeat.co/badges/f1db1585-3835-4c83-bda4-4dad2c499b33)](https://codebeat.co/projects/github-com-piankalabs-math-master)
[![Go Report Card](https://goreportcard.com/badge/github.com/piankalabs/math)](https://goreportcard.com/report/github.com/piankalabs/math)
[![Go 1.14](https://img.shields.io/badge/go-1.14-blue.svg)](http://java.oracle.com)
[![License](https://img.shields.io/badge/license-Apache_2.0-blue.svg)](https://raw.githubusercontent.com/piankalabs/math/master/LICENSE)

# Math Expressions
A command line interface for evaluating simple mathematical expressions using [ANTLR4](https://github.com/antlr/antlr4) in Go.

## Functionality
The following statement types are supported:

* Number literals (integer and decimal)
* Unary operators (negative numbers)
* Infix operators (add, subtract, multiply, divide)
* Parentheticals (putting any statement within parentheses)

The order of operations follows [PEMDAS](https://www.wikiwand.com/en/Order_of_operations#/Mnemonics) for the subset of operations supported.

## Examples
Below are some examples of supported mathematical expressions:

`2` outputs `2`

`2 + 2` outputs `4`

`5 - 4` outputs `1`

`2 * -8` outputs `-16`

`10 * (1 / 2)` outputs `5`

`((1 + 2) - 3) - (4 - 5)` outputs `1`

## Installation
Installing the command line interface is as simple as:

`./scripts/install.sh`

## Usage
Invoking the command line interface is as simple as:

`math 2 + 2`

### Shells
One important caveat is that shell semantics can interfere with mathematical notation.
For example, using [Zsh](https://www.zsh.org/) for the following invocation:

`math (1 + 1)` produces the error `zsh: unknown file attribute: 1`

And similarly, the CLI library used in this project ([urfave/cli](https://github.com/urfave/cli))
will occasionally have some parsing issues, such as:

`math -2 + 0` produces the error `Incorrect Usage. flag provided but not defined: -2`

These issues can be overcome by using quotes and parentheses:

`math "(1 + 1)"` outputs `2`

`math "(-2 + 0)"` outputs `-2`