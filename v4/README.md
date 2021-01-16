# workfx
> Module enabling your `Fx` application to effectively track and commit
changes to your entities.

[![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Coverage Status][coverage-img]][coverage] [![Release][release-img]][release] [![License][license-img]][license] [![Blog][blog-img]][blog]

## Example Usage

### Provide the Module

```go
package main

import (
	"github.com/freerware/workfx/v4"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		// 🎉
		workfx.Module,
		// ... other modules.
	).Run()
}
```

### Dependencies

The only dependency of `workfx` is `[]unit.Option`. These options are utilized
to create work units in your [`Fx`][fx] application. `workfx` pulls in these options
using the `group:"unitOptions"` [value group][value-groups].

[fx]: https://github.com/uber-go/fx
[value-groups]: https://godoc.org/go.uber.org/fx#hdr-Value_Groups
[doc-img]: https://godoc.org/github.com/freerware/workfx?status.svg
[doc]: https://godoc.org/github.com/freerware/workfx
[ci-img]: https://travis-ci.org/freerware/workfx.svg?branch=master
[ci]: https://travis-ci.org/freerware/workfx
[coverage-img]: https://coveralls.io/repos/github/freerware/workfx/badge.svg?branch=master
[coverage]: https://coveralls.io/github/freerware/workfx?branch=master
[release]: https://github.com/freerware/workfx/releases
[release-img]: https://img.shields.io/github/tag/freerware/workfx.svg?label=version
[blog]: https://medium.com/@freerjm/work-units-ec2da48cf574
[blog-img]: https://img.shields.io/badge/blog-medium-lightgrey
