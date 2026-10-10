package openssl

import (
	"testing"

	"github.com/llarhub/openssl"
)

func TestOpenSSLVersionAndTLSContext(t *testing.T) {
	if got := openssl.OPENSSLVersionMajor(); got < 1 {
		t.Fatalf("OpenSSL major version = %d", got)
	}

	method := openssl.TLSMethod()
	if method == nil {
		t.Fatal("TLSMethod returned nil")
	}

	ctx := openssl.SSL_CTXNew(method)
	if ctx == nil {
		t.Fatal("SSL_CTXNew returned nil")
	}
	openssl.SSL_CTXFree(ctx)
}
