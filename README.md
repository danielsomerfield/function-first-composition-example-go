# TDD and function-first composition with Go

This repository is a companion to the
[Dependency Composition](https://martinfowler.com/articles/dependency-composition.html)
that discusses an alternative approach to building services using a somewhat unconventional combination of TDD,
partial application, and functions as the primary compositional unit. The article is written using Typescript
examples, but this is a port for Go, which means the pattern is necessarily a little bit different.

# Points of interest

- Each component module defines a type with its own dependent functions.
- At the top level of the injection is done via the `injection.go`
- Rather than constructors for components, modules expose a factory function that takes dependencies and
  returns either a "configured" function or component.

# Differences from the Typescript Example

Typescript's structure typing makes some things possible that are not in a language like Go. Go has structural typing but 
it's very different. Fulfilling an interface contract consists of implementing matching functions for a given type. 
Unfortunately, in order to implement an interface, you need an underlying struct and that struct cannot be anonymous, 
requiring a number of structs that have no purpose other to bind to functions. The result was a ton of template code.
Instead, I used structs and assigned functions to them. This is closer to how the Typescript code was implemented. The 
downside is, however, that structs are nominally typed, meaning two types are considered different even if they are 
identical in composition. The means I have to use literally the same type and since Go doesn't provide any mechanism to 
prevent Nil-assigned fields, the type system won't help you if you forget to bind a function. The resulting composition
is clear, but it's non-idiomatic and the type safety a bit weak.

In summary, I think the basic TDD and module injection approach works, but I'd try to achieve it in a somewhat different.
Given that I don't work with Go very much, I will most likely leave that as an exercise to the reader.

# Instructions

## Running the tests

    make test

## Running the application

    make run

You should then be able to access the service endpoint at `http://localhost:8080/vancouverbc/restaurants/recommended`

