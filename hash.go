package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"io"
	"os"

	"github.com/spaolacci/murmur3"
)

type Hash []byte

func (h Hash) B64() string {
	return base64.RawURLEncoding.EncodeToString(h)
}

func (h Hash) Bytes() []byte {
	return append([]byte{}, h...)
}

func (h Hash) Hex() string {
	return hex.EncodeToString(h[:])
}

func Murmur3Hash(data []byte) Hash {
	hasher := murmur3.New128()
	hasher.Write(data)
	u1, u2 := hasher.Sum128()

	b := []byte{}

	b = binary.BigEndian.AppendUint64(b, u1)
	b = binary.BigEndian.AppendUint64(b, u2)

	return b
}

func SHA224(data []byte) Hash {
	h := sha256.New224()
	h.Write(data)
	buf := h.Sum(nil)
	return buf
}
func SHA224String(s string) Hash {
	return SHA224([]byte(s))
}

func HMACSHA224(data []byte, key []byte) Hash {
	h := hmac.New(sha256.New224, key)
	h.Write(data)
	buf := h.Sum(nil)
	return buf
}

func HMACSHA224String(data, key string) Hash {
	return HMACSHA224([]byte(data), []byte(key))
}

func SHA256(data []byte) Hash {
	h := sha256.New()
	h.Write(data)
	buf := h.Sum(nil)
	return buf
}
func SHA256String(s string) Hash {
	return SHA256([]byte(s))
}

func HMACSHA256(data []byte, key []byte) Hash {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	buf := h.Sum(nil)
	return buf
}

func HMACSHA256String(data, key string) Hash {
	return HMACSHA256([]byte(data), []byte(key))
}

func SHA224File(path string) Hash {
	h := sha256.New()
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	io.Copy(h, file)

	return h.Sum(nil)
}
