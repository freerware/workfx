# workfx
> Modules enabling your `Fx` application to effectively track and commit changes to your entities.

[![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Coverage Status][coverage-img]][coverage] [![Release][release-img]][release] [![License][license-img]][license] [![Blog][blog-img]][blog]

## What is it?

`workfx` empowers your [`Fx`][fx] application with the ability to track and
commit atomic changes to your entities. It essentially defines core 
module(s) that can be imported into your `Fx` application so that it can
leverage [`work.Uniter`][uniter-doc] instances.

## Why use it?

With `workfx` you can seamlessly integrate work units into your `Fx`
application with one line of code.

## Release information

### [4.0.0-beta][v4.0.0-beta]

- Leverage `v4.0.0-beta` of `work`.

### 2.x.x

- NO LONGER SUPPORTED.

### 1.x.x

- NO LONGER SUPPORTED.

## Dependancy Information

As of [`v4.0.0-beta`][modules-release], the project utilizes [modules][modules-doc].
Prior to `v4.0.0-beta`, the project utilized [`dep`][dep] for dependency management.

In order to transition to modules gracefully, we adhered to the
[best practice recommendations][modules-wiki] authored by the Golang team.

## Release information

Versions `1.x.x` and `2.x.x` are no longer supported. Please upgrade to
`4.x.x+` to receive the latest and greatest features of work units!

## Contribute

Want to lend us a hand? Check out our guidelines for [contributing][contributing].

## License

We are rocking an [Apache 2.0 license][apache-license] for this project.

## Code of Conduct

Please check out our [code of conduct][code-of-conduct] to get up to speed how we do things.

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
[release]: https://github.com/freerware/workfx/releases
[release-img]: https://img.shields.io/github/tag/freerware/workfx.svg?label=version
[blog]: https://medium.com/@freerjm/work-units-ec2da48cf574
[blog-img]: https://img.shields.io/badge/blog-medium-lightgrey
[v4.0.0-beta]: https://github.com/freerware/workfx/releases/tag/v4.0.0-beta
[modules-doc]: https://golang.org/doc/go1.11#modules
[modules-wiki]: https://github.com/golang/go/wiki/Modules#releasing-modules-v2-or-higher
[modules-release]: https://github.com/freerware/workfx/releases/tag/v4.0.0-beta
[dep]: https://golang.github.io/dep/
