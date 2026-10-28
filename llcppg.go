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

func (self *MD5_CTX) UpdateBytes(data []byte) c.Int {
	return self.Update(unsafe.Pointer(unsafe.SliceData(data)), uintptr(len(data)))
}

func (self *MD5_CTX) UpdateString(data string) c.Int {
	return self.Update(unsafe.Pointer(unsafe.StringData(data)), uintptr(len(data)))
}

func (self *MD5_CTX) Final(md *byte) c.Int {
	return MD5_Final(md, self)
}

func (self *HMAC_CTX) UpdateBytes(data []byte) c.Int {
	return self.Update((*uint8)(unsafe.SliceData(data)), c.SizeT(len(data)))
}

func (self *HMAC_CTX) UpdateString(data string) c.Int {
	return self.Update((*uint8)(unsafe.StringData(data)), c.SizeT(len(data)))
}

//llgo:link (*BIGNUM).Cmp C.BN_cmp
func (self *BIGNUM) Cmp(b *BIGNUM) c.Int {
	return 0
}

// llgo:link (*BIGNUM).CStr C.BN_bn2dec
func (*BIGNUM) CStr() *c.Char { return nil }
