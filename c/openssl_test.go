package openssl

import (
	"testing"

	"github.com/llarhub/openssl"
)

func TestOpenSSLVersionAndTLSContext(t *testing.T) {
	if got := openssl.VersionMajor(); got < 1 {
		t.Fatalf("OpenSSL major version = %d", got)
	}

	method := openssl.TLS_method()
	if method == nil {
		t.Fatal("TLS_method returned nil")
	}

	ctx := method.CTX_new()
	if ctx == nil {
		t.Fatal("CTX_new returned nil")
	}
	ctx.Free()
}
