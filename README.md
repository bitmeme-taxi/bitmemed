
Bitmemed
====

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/bitmeme-taxi/bitmemed/cmd/bitmemed)

Bitmemed is the reference full node Bitmeme implementation written in Go (golang).

## What is Bitmeme Network

BTM is decentralized cryptocurrency, that is making waves in the world of BlockChain.
This innovative project is based on the Proof of Work consensus algorithm and uses the Blake3 hashing function to ensure the security and efficiency of its network.

## Requirements

Go 1.18 or later.

## Installation

#### Build from Source

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Ensure Go was installed properly and is a supported version:

```bash
$ go version
```

- Run the following commands to obtain and install bitmemed including all dependencies:

```bash
$ git clone https://github.com/bitmeme-taxi/bitmemed/cmd/bitmemed
$ cd bitmemed
$ [go install . ./cmd/...]
$ build.sh

```

- Bitmemed (and utilities) should now be installed in `$(go env GOPATH)/bin`. If you did
  not already add the bin directory to your system path during Go installation,
  you are encouraged to do so now.


## Getting Started

Bitmemed has several configuration options available to tweak how it runs, but all
of the basic operations work with zero configuration.

```bash
$ cd ~/go/bin
$ bitmemed --utxoindex
```



## Website
Join our website server using the following link: https://bitmeme.world/

## Twitter
Join our twitter server using the following link: https://twitter.com/BTMCurrency

## Discord
Join our discord server using the following link: https://discord.gg/

## Telegram
Join our telegram server using the following link: https://t.me/BTMcurrency
