# OpenSSL

[![Go Reference](https://pkg.go.dev/badge/github.com/llarhub/openssl.svg)](https://pkg.go.dev/github.com/llarhub/openssl)
[![LLGo](https://img.shields.io/badge/powered_by-LLGo-green.svg)](https://github.com/xgo-dev/llgo)
[![XGo](https://img.shields.io/badge/project-XGo-blue.svg)](https://github.com/goplus/xgo)

LLGo bindings for the OpenSSL C API. The generated package includes TLS/SSL, EVP, BIO, X.509, RSA, EC, BN, hashing, and related public APIs reachable from `openssl/ssl.h`.

## Installation

Install LLGo, OpenSSL's development package, and `pkg-config`, then add the module:

```bash
go get github.com/llarhub/openssl
```

On macOS with Homebrew:

```bash
brew install openssl pkg-config
```

On Debian or Ubuntu:

```bash
sudo apt update
sudo apt install libssl-dev pkg-config
```

The binding links with `pkg-config --libs openssl`, `-lssl`, and `-lcrypto`. Make sure `pkg-config` resolves the same OpenSSL installation as the headers used by your LLGo build.

## Usage

Create and release a TLS context:

```go
package main

import "github.com/llarhub/openssl"

func main() {
	method := openssl.TLSMethod()
	ctx := openssl.SSL_CTXNew(method)
	if ctx == nil {
		panic("could not create OpenSSL TLS context")
	}
	openssl.SSL_CTXFree(ctx)
}
```

The generated declarations track the installed OpenSSL headers. If the library version or header layout changes, update the inputs on the `c` branch and regenerate the binding with `llcppg`.

Run the binding test from the `c` worktree with:

```bash
cd _out/c
llgo test ./...
```
