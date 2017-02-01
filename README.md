# leri

[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser)

a small (and currently quite limited) command line tool to generate markdown files from source code comments

## why?

I used *leri* mainly to generate some tutorial-like documents from heavily documented source code.

## features

*leri* currently only supports `.go` and `.sas` files, but could be extended easily by defining own regular expressions to match comments that should be interpreted as documentation, (see: `lib/parsing/parser.go` and parser usage in file `lib/commands/genmd.go`).

## installation

```
go get github.com/blblblu/leri/cmd/leri
```

(or download a pre-compiled version from the [releases page](https://github.com/blblblu/leri/releases))

## usage example

Let's assume you have a file `lorem.go` with following content:

```go
package ecs

import "github.com/blblblu/reba/env"

// Initer is used for Systems or Entities that want to initialized based on the env.Context before the main loop
type Initer interface {
	Init(ctx *env.Context)
}

// Updater is used for Systems that want to be involved in the main loop
type Updater interface {
	Update(ctx *env.Context)
}

// Cleanuper is used for Systems that want to free resources (like e.g. deleting OpenGL buffers etc.)
type Cleanuper interface {
	Cleanup(ctx *env.Context)
}
```

When running `leri gen -i lorem.go -o lorem.md`, *leri* will recognize each `//` comment that starts at the beginning of the line as documentation, and will create the following `lorem.md` file:

    ```go
    package ecs

    import "github.com/blblblu/reba/env"

    ```

    Initer is used for Systems or Entities that want to initialized based on the env.Context before the main loop

    ```go
    type Initer interface {
      Init(ctx *env.Context)
    }

    ```

    Updater is used for Systems that want to be involved in the main loop

    ```go
    type Updater interface {
      Update(ctx *env.Context)
    }

    ```

    Cleanuper is used for Systems that want to free resources (like e.g. deleting OpenGL buffers etc.)

    ```go
    type Cleanuper interface {
      Cleanup(ctx *env.Context)
    }

    ```

which will look as follows:

---


```go
package ecs

import "github.com/blblblu/reba/env"

```

Initer is used for Systems or Entities that want to initialized based on the env.Context before the main loop

```go
type Initer interface {
  Init(ctx *env.Context)
}

```

Updater is used for Systems that want to be involved in the main loop

```go
type Updater interface {
  Update(ctx *env.Context)
}

```

Cleanuper is used for Systems that want to free resources (like e.g. deleting OpenGL buffers etc.)

```go
type Cleanuper interface {
  Cleanup(ctx *env.Context)
}

```
