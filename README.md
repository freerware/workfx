<p align="center"><img src="https://gophercises.com/img/gophercises_jumping.gif" width="360"></p>

# workfx
> Modules enabling your `Fx` application to effectively track and commit changes to your entities.

[![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Coverage Status][coverage-img]][coverage] [![Release][release-img]][release] [![License][license-img]][license] [![Blog][blog-img]][blog]

## What is it?

`workfx` empowers your [Fx][fx] application with the ability to track and commit atomic changes to your entities. It essentially defines a core set of Fx modules that can be imported into your Fx application so that it can leverage [`work.Uniter`][uniter-doc] instances.

## Why use it?

With `workfx`, you can seamlessly integrate work units into your Fx application. On top of the various benefits of using work units in general, the `workfx` module provides:

- a well defined set of modules that can be used in isolation or in conjuction.
- integrations with Fx dependency management, reducing the necessary code to create `work.Uniter` instances in your Fx application.

## Example Usage

### Provide the Module

```go
package main

import (
	"github.com/freerware/workfx"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		... // modules for your Fx application.
		workfx.Modules.SQLUnit,
	).Run()
}
```

### Named Values

To support Fx applications that may leverage more than one work unit type, `workfx` utilizes named values. This means that in order to inject `work.Uniter` instances, you must add the `name` Go tag. In addition, it's common for applications leveraging SQL databases to support configurations with multiple database instances, such as primary (read-write) and replicas (read-only). In order to accommodate these applications, the `name` tag is used as well.

| Name                   | Type          | Description        |
| ---------------------- | ------------- | ------------------ |
| `sqlWorkUniter`        | `work.Uniter` | SQL Uniter         |
| `bestEffortWorkUniter` | `work.Uniter` | Best Effort Uniter |
| `rwDB`                 | `*sql.DB`     | Read-Write DB      |

```go
type Parameters struct {
	fx.In

	uniter work.Uniter `name:"sqlWorkUniter"`
}
```

### Value Groups

In addition to named values, `workfx` also makes use of value groups as a convenience for injecting multiple instances of the same type.

| Name                   | Type          | Description       |
| ---------------------- | ------------- | ----------------- |
| `workUniter`           | `work.Uniter` | Work Uniter Group |

```go
type Parameters struct {
	fx.In

	uniters []work.Uniter `group:"workUniter"`
}
```

## Contribute

Want to lend us a hand? Check out our guidelines for [contributing][contributing].

## License

We are rocking an [Apache 2.0 license][apache-license] for this project.

## Code of Conduct

Please check out our [code of conduct][code-of-conduct] to get up to speed how we do things.

## Artwork

Discovered via the interwebs, the artwork was created by Marcus Olsson and Jon Calhoun for [Gophercises][gophercises].

[fx]: https://github.com/uber-go/fx
[uniter-doc]: https://godoc.org/github.com/freerware/work#Uniter
[doc-img]: https://godoc.org/github.com/freerware/workfx?status.svg
[doc]: https://godoc.org/github.com/freerware/workfx
[ci-img]: https://travis-ci.org/freerware/workfx.svg?branch=master
[ci]: https://travis-ci.org/freerware/workfx
[coverage-img]: https://coveralls.io/repos/github/freerware/workfx/badge.svg?branch=master
[coverage]: https://coveralls.io/github/freerware/workfx?branch=master
[license]: https://opensource.org/licenses/Apache-2.0
[license-img]: https://img.shields.io/badge/License-Apache%202.0-blue.svg
[contributing]: https://github.com/freerware/workfx/blob/master/CONTRIBUTING.md
[apache-license]: https://github.com/freerware/workfx/blob/master/LICENSE.txt
[code-of-conduct]: https://github.com/freerware/workfx/blob/master/CODE_OF_CONDUCT.md
[gophercises]: https://gophercises.com
[release]: https://github.com/freerware/workfx/releases
[release-img]: https://img.shields.io/github/tag/freerware/workfx.svg?label=version
[blog]: https://medium.com/@freerjm/work-units-ec2da48cf574
[blog-img]: https://img.shields.io/badge/blog-medium-lightgrey
