# Go Substrate RPC Client (GSRPC)

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![GoDoc Reference](https://godoc.org/github.com/centrifuge/go-substrate-rpc-client?status.svg)](https://godoc.org/github.com/centrifuge/go-substrate-rpc-client)
[![Build Status](https://travis-ci.com/centrifuge/go-substrate-rpc-client.svg?branch=master)](https://travis-ci.com/centrifuge/go-substrate-rpc-client)
[![codecov](https://codecov.io/gh/centrifuge/go-substrate-rpc-client/branch/master/graph/badge.svg)](https://codecov.io/gh/centrifuge/go-substrate-rpc-client)
[![Go Report Card](https://goreportcard.com/badge/github.com/centrifuge/go-substrate-rpc-client)](https://goreportcard.com/report/github.com/centrifuge/go-substrate-rpc-client)

Substrate RPC client in Go. It provides APIs and types around Polkadot and any Substrate-based chain RPC calls.
This client is modelled after [polkadot-js/api](https://github.com/polkadot-js/api).

## State

This package is actively developed. The first stable release is expected in November 2019.

## Documentation & Usage Examples

Please refer to https://godoc.org/github.com/centrifuge/go-substrate-rpc-client

## Contributing

1. Install dependencies by running `make` followed by `make install`
1. Run tests `make test`
1. Lint `make lint` (you can use `make lint-fix` to automatically fix issues)

## Run tests in a Docker container against the Substrate Default Docker image

1. Run the docker container `make test-dockerized`

## Run tests locally against the Substrate Default Docker image

1. Start the Substrate Default Docker image: `make run-substrate-docker`
1. In another terminal, run the tests against that image: `make test`
1. Visit https://polkadot.js.org/apps for inspection

## Run tests locally against any substrate endpoint

1. Set the endpoint: `export RPC_URL="http://example.com:9933"`
1. Run the tests `make test`
