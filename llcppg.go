package openssl

import (
	"unsafe"

	"github.com/goplus/lib/c"
)

func Free(ptr c.Pointer) {
	CRYPTO_free(ptr, c.Str(""), 0)
}

func FreeCStr(ptr *c.Char) {
	CRYPTO_free(c.Pointer(ptr), c.Str(""), 0)
}

func (ctx *HMAC_CTX) UpdateBytes(data []byte) c.Int {
	return ctx.Update((*uint8)(unsafe.SliceData(data)), c.SizeT(len(data)))
}

func (ctx *HMAC_CTX) UpdateString(data string) c.Int {
	return ctx.Update((*uint8)(unsafe.StringData(data)), c.SizeT(len(data)))
}

//llgo:link (*BIGNUM).Cmp C.BN_cmp
func (self *BIGNUM) Cmp(b *BIGNUM) c.Int {
	return 0
}

// llgo:link (*BIGNUM).CStr C.BN_bn2dec
func (*BIGNUM) CStr() *c.Char { return nil }
