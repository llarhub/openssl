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

func RAND_Bytes(buf []byte) c.Int {
	return RAND_bytes(unsafe.SliceData(buf), c.Int(len(buf)))
}

func RAND_PrivBytes(buf []byte) c.Int {
	return RAND_priv_bytes(unsafe.SliceData(buf), c.Int(len(buf)))
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

func (self *SHA_CTX) UpdateBytes(data []byte) c.Int {
	return self.Update(unsafe.Pointer(unsafe.SliceData(data)), uintptr(len(data)))
}

func (self *SHA_CTX) UpdateString(data string) c.Int {
	return self.Update(unsafe.Pointer(unsafe.StringData(data)), uintptr(len(data)))
}

func (self *SHA_CTX) Final(md *byte) c.Int {
	return SHA1_Final(md, self)
}

func SHA1_Bytes(data []byte, md *byte) *byte {
	return SHA1((*byte)(unsafe.SliceData(data)), uintptr(len(data)), md)
}

func SHA1_String(data string, md *byte) *byte {
	return SHA1((*byte)(unsafe.StringData(data)), uintptr(len(data)), md)
}

func (c *SHA256_CTX) UpdateBytes(data []byte) c.Int {
	return c.Update(unsafe.Pointer(unsafe.SliceData(data)), uintptr(len(data)))
}

func (c *SHA256_CTX) UpdateString(data string) c.Int {
	return c.Update(unsafe.Pointer(unsafe.StringData(data)), uintptr(len(data)))
}

func (c *SHA256_CTX) Final(md *byte) c.Int {
	return SHA256_Final(md, c)
}

func (c *SHA256_CTX) SHA224_UpdateBytes(data []byte) c.Int {
	return c.SHA224_Update(unsafe.Pointer(unsafe.SliceData(data)), uintptr(len(data)))
}

func (c *SHA256_CTX) SHA224_UpdateString(data string) c.Int {
	return c.SHA224_Update(unsafe.Pointer(unsafe.StringData(data)), uintptr(len(data)))
}

func (c *SHA256_CTX) SHA224_Final(md *byte) c.Int {
	return SHA224_Final(md, c)
}

func SHA256_Bytes(data []byte, md *byte) *byte {
	return SHA256((*byte)(unsafe.SliceData(data)), uintptr(len(data)), md)
}

func SHA256_String(data string, md *byte) *byte {
	return SHA256((*byte)(unsafe.StringData(data)), uintptr(len(data)), md)
}

func SHA224_Bytes(data []byte, md *byte) *byte {
	return SHA224((*byte)(unsafe.SliceData(data)), uintptr(len(data)), md)
}

func SHA224_String(data string, md *byte) *byte {
	return SHA224((*byte)(unsafe.StringData(data)), uintptr(len(data)), md)
}

func (c *SHA512_CTX) UpdateBytes(data []byte) c.Int {
	return c.Update(unsafe.Pointer(unsafe.SliceData(data)), uintptr(len(data)))
}

func (c *SHA512_CTX) UpdateString(data string) c.Int {
	return c.Update(unsafe.Pointer(unsafe.StringData(data)), uintptr(len(data)))
}

func (c *SHA512_CTX) Final(md *byte) c.Int {
	return SHA512_Final(md, c)
}

func (c *SHA512_CTX) SHA384_UpdateBytes(data []byte) c.Int {
	return c.SHA384_Update(unsafe.Pointer(unsafe.SliceData(data)), uintptr(len(data)))
}

func (c *SHA512_CTX) SHA384_UpdateString(data string) c.Int {
	return c.SHA384_Update(unsafe.Pointer(unsafe.StringData(data)), uintptr(len(data)))
}

func (c *SHA512_CTX) SHA384_Final(md *byte) c.Int {
	return SHA384_Final(md, c)
}

func SHA512_Bytes(data []byte, md *byte) *byte {
	return SHA512((*byte)(unsafe.SliceData(data)), uintptr(len(data)), md)
}

func SHA512_String(data string, md *byte) *byte {
	return SHA512((*byte)(unsafe.StringData(data)), uintptr(len(data)), md)
}

func SHA384_Bytes(data []byte, md *byte) *byte {
	return SHA384((*byte)(unsafe.SliceData(data)), uintptr(len(data)), md)
}

func SHA384_String(data string, md *byte) *byte {
	return SHA384((*byte)(unsafe.StringData(data)), uintptr(len(data)), md)
}

func (self *HMAC_CTX) UpdateBytes(data []byte) c.Int {
	return self.Update((*byte)(unsafe.SliceData(data)), c.SizeT(len(data)))
}

func (self *HMAC_CTX) UpdateString(data string) c.Int {
	return self.Update((*byte)(unsafe.StringData(data)), c.SizeT(len(data)))
}

//llgo:link (*BIGNUM).Cmp C.BN_cmp
func (self *BIGNUM) Cmp(b *BIGNUM) c.Int {
	return 0
}

// llgo:link (*BIGNUM).CStr C.BN_bn2dec
func (*BIGNUM) CStr() *c.Char { return nil }
