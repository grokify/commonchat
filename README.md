# CommonChat

[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

CommonChat is an abstraction library for chat / team messaging services like Glip and Slack. It currently includes two parts:

* Common message format - After converting a message to the `commonchat.Message` format, the libraries can be used to convert to individula chat services.
* Webhook clients - Given a webhook URL, the clients use the `commonchat.Adapter` interface to enable webhook API calls using the `commonchat.Message` format.

It is currently used with the Chathooks webhook formatting service:

[https://github.com/grokify/chathooks](https://github.com/grokify/chathooks)

 [build-status-svg]: https://github.com/grokify/commonchat/workflows/test/badge.svg
 [build-status-url]: https://github.com/grokify/commonchat/actions
 [coverage-status-svg]: https://coveralls.io/repos/grokify/commonchat/badge.svg?branch=master
 [coverage-status-url]: https://coveralls.io/r/grokify/commonchat?branch=master
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/commonchat
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/commonchat
 [codeclimate-status-svg]: https://codeclimate.com/github/grokify/commonchat/badges/gpa.svg
 [codeclimate-status-url]: https://codeclimate.com/github/grokify/commonchat
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/commonchat
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/commonchat
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/commonchat/blob/master/LICENSE
