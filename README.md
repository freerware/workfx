# workfx
> A module enabling your Fx application to effectively track and commit changes to your entities.

[![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Coverage Status][coverage-img]][coverage]

## What is it?

`workfx` empowers your [Fx](https://github.com/uber-go/fx) application with the ability to track and commit atomic changes to your entities. It essentially defines a core set of Fx modules that can be imported into your Fx application so that it can leverage `work.Uniter` instances.

## Why use it?

With `workfx`, you can seamlessly integrate work units into your Fx application. On top of the various benefits of using work units in general, the `workfx` module provides:

- a well defined set of modules that can be used in isolation or in conjuction.
- integrations with Fx dependency management, reducing the necessary code to create `work.Uniter` instances in your Fx application.

## Example Usage

```golang
package main

import (
	"github.com/freerware/workfx"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		... // modules for your Fx application.
		workfx.Modules.SQLWorkUnit,
	).Run()
}
```

[doc-img]: https://godoc.org/github.com/freerware/workfx?status.svg
[doc]: https://godoc.org/github.com/freerware/workfx
[ci-img]: https://travis-ci.org/freerware/workfx.svg?branch=master
[ci]: https://travis-ci.org/freerware/workfx
[coverage-img]: https://coveralls.io/repos/github/freerware/workfx/badge.svg?branch=master
[coverage]: https://coveralls.io/github/freerware/workfx?branch=master
