# go-inputKit

A simple library for generating CLI-prompt based applications

## Why?

When people are creating prompts in golang(as well as most other programming langauges), they must use a slew of `if`s,`switch`es,`for`s and other annoying control structures. It makes _NO_ sense why you have to design a hierarchical prompt without using a hierarchical structure. This made me think, _Gee, it would really be nice if you could just put all of your stuff into a struct, then have it auto-generate all of the boilerplate for you_. So, that is what I have made.

## How?

`import ik "github.com/liamnaddell/go-inputKit"`

## Example

It is in `cmd/ik/main.go`. This thing demonstrates all of the stuff available in `go-inputKit`
