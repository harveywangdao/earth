package main

import (
	"fmt"
)

func RsaWithSHA256Base64(origData string, block []byte) (sign string, err error) {
	blocks, _ := pem.Decode(block)
	privateKey, _ := x509.ParsePKCS8PrivateKey(blocks.Bytes)
	h := sha256.New()
	h.Write([]byte(origData))
	digest := h.Sum(nil)
	s, _ := rsa.SignPKCS1v15(nil, privateKey.(*rsa.PrivateKey), crypto.SHA256, digest)
	sign = base64.StdEncoding.EncodeToString(s)
	return
}

func main() {

}
