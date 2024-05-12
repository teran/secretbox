# secretbox

![Test & Build status](https://github.com/teran/secretbox/actions/workflows/go.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/teran/secretbox)](https://goreportcard.com/report/github.com/teran/secretbox)
[![Go Reference](https://pkg.go.dev/badge/github.com/teran/secretbox.svg)](https://pkg.go.dev/github.com/teran/secretbox)

Trivial library, example server and client applications to provide secrets for
applications from third party password manager authenticated once.

## Use case

Secretbox provides ability to run any other application which can obtain
secrets from a command (like [restic](https://restic.net) or
[rclone](https://rclone.org)) to get the secrets it needs with one-time
token through unix-socket. In that scenario service side must be used
to serve unix-socket connections by application runner.

Examples can be found in cmd: cli and server side.
