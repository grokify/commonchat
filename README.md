# CommonChat

[![Go CI][go-ci-svg]][go-ci-url]
[![Go Lint][go-lint-svg]][go-lint-url]
[![Go SAST][go-sast-svg]][go-sast-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![Visualization][viz-svg]][viz-url]
[![License][license-svg]][license-url]

 [go-ci-svg]: https://github.com/grokify/commonchat/actions/workflows/go-ci.yaml/badge.svg?branch=main
 [go-ci-url]: https://github.com/grokify/commonchat/actions/workflows/go-ci.yaml
 [go-lint-svg]: https://github.com/grokify/commonchat/actions/workflows/go-lint.yaml/badge.svg?branch=main
 [go-lint-url]: https://github.com/grokify/commonchat/actions/workflows/go-lint.yaml
 [go-sast-svg]: https://github.com/grokify/commonchat/actions/workflows/go-sast-codeql.yaml/badge.svg?branch=main
 [go-sast-url]: https://github.com/grokify/commonchat/actions/workflows/go-sast-codeql.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/commonchat
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/commonchat
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/commonchat
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/commonchat
 [viz-svg]: https://img.shields.io/badge/visualizaton-Go-blue.svg
 [viz-url]: https://mango-dune-07a8b7110.1.azurestaticapps.net/?repo=grokify%2Fcommonchat
 [loc-svg]: https://tokei.rs/b1/github/grokify/commonchat
 [repo-url]: https://github.com/grokify/commonchat
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/commonchat/blob/master/LICENSE

CommonChat is an abstraction library for chat / team messaging services like Glip and Slack. It currently includes two parts:

* Common message format - After converting a message to the `commonchat.Message` format, the libraries can be used to convert to individula chat services.
* Webhook clients - Given a webhook URL, the clients use the `commonchat.Adapter` interface to enable webhook API calls using the `commonchat.Message` format.

It is currently used with the Chathooks webhook formatting service:

[https://github.com/grokify/chathooks](https://github.com/grokify/chathooks)
