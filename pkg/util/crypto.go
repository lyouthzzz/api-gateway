package util

import (
	"crypto/hmac"
	"crypto/sha256"
)

func HmacSha256(data []byte, key []byte) ([]byte, error) {
	h := hmac.New(sha256.New, key)
	if _, err := h.Write(data); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}
