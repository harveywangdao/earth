package main

import (
	"crypto"
	//"crypto/dsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//sudo openssl genrsa -out rsa_private_key.pem 1024
//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem

//openssl pkcs8 -topk8 -inform PEM -in rsa_private_key.pem -outform PEM -nocrypt -out pkcs8.pem
//openssl rsa -in pkcs8.pem -out pkcs1.pem
//openssl x509 -in cert.pem -text –noout
//openssl req -new -x509 -newkey rsa:1024 -keyout CA.key -out CA.pem
/*var privatekey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQDFsNRpwxr8UkNYO/guYeYjKnI6nDClixwD5TTJqu8Z1J+oe2vi
NsdgJffkpq8GnR2Luk5uspNx0lkbEI/nqi13bnZRZsc351GrO76kI1+jZooSqkLx
zEpYojyQ2G35MUvlNW3vrYScS92QfKefSnsUdPuY082d+/+4LBva7lc+ewIDAQAB
AoGAFctFz5cTzYdgJ0UNOkfOsEfIqg94CXgJkaBeLnFqxKU9KLzaiujRGBd3ebyq
hQcjL7lNVkTXnm+7JRGxW6/WLSfgC+Er44+A+OadNC+3uP+5uOTmdaXc7VaHznIr
NM0UmRJpq9a1pJshM+NTISMYUQFh03rhCcWZAFTH7Q1Rl/kCQQDneVcsA0Fb3ZGW
P3eX4WrSW3oF1PkH8LTz9hPJC4BQFE8AKceTVDKtPm20659nEA4W7O5v7uwNGv+y
KhXUscLfAkEA2qMhF2BoO9pqPBTQCImj7Gs8Gx2ymWKz9qisQdFFULWLQlCDgTW3
XduVR4nz1VXVKY1gjyZsQBPQHhgdJC6z5QJAelqne3pVCAw8sqDlxIDBCGSWQZLr
+AiCfhprJvBPaQTzJXDwL65oAy9mqiWKYt4XtIKVHtG6MMs+sZyRYZZ9HQJAK/ul
LI64aVIHDQ8iypTl5SmtbccLps+0ZGqKPLNcvl/HJBAOZG/p83e5ECswYjpiJ3nJ
cwLHxBw4QdWYBy9eAQJBAJCYdYPesTIrYqV8K3oV9FU4qX6u0pxGYJ49ItJQlxJ5
DF93bbi7CoVx4DTnRwqR8wKWnMphTuMh4ZXsO5JLQFw=
-----END RSA PRIVATE KEY-----
`)

var publickey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDFsNRpwxr8UkNYO/guYeYjKnI6
nDClixwD5TTJqu8Z1J+oe2viNsdgJffkpq8GnR2Luk5uspNx0lkbEI/nqi13bnZR
Zsc351GrO76kI1+jZooSqkLxzEpYojyQ2G35MUvlNW3vrYScS92QfKefSnsUdPuY
082d+/+4LBva7lc+ewIDAQAB
-----END PUBLIC KEY-----
`)
*/
func RSA_Encrypt(pubkey, msg []byte) ([]byte, error) {
	block, _ := pem.Decode(pubkey)
	if block == nil {
		log.Println("public key error")
		return nil, errors.New("public key error")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	pub, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		log.Println("public key error")
		return nil, errors.New("public key error")
	}

	return rsa.EncryptPKCS1v15(rand.Reader, pub, msg)
}

func RSA_Decrypt(privkey, ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privkey)
	if block == nil {
		log.Println("private key error")
		return nil, errors.New("private key error")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func RSA_Sign(privkey, msg []byte) ([]byte, error) {
	block, _ := pem.Decode(privkey)
	if block == nil {
		log.Println("private key error")
		return nil, errors.New("private key error")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	msgHash := sha256.Sum256(msg)

	sig, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, msgHash[:])
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return sig, nil
}

func RSA_Verify(pubkey, msg, sig []byte) (bool, error) {
	block, _ := pem.Decode(pubkey)
	if block == nil {
		log.Println("public key error")
		return false, errors.New("public key error")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Println(err)
		return false, err
	}

	pub, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		log.Println("public key error")
		return false, errors.New("public key error")
	}

	msgHash := sha256.Sum256(msg)

	err = rsa.VerifyPKCS1v15(pub, crypto.SHA256, msgHash[:], sig)
	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, nil
}

func RSA_GenerateKeyPair(password []byte, privKeyFile, pubKeyFile string) error {
	privKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		log.Println(err)
		return err
	}

	privKeyData := x509.MarshalPKCS1PrivateKey(privKey)

	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privKeyData,
	}

	file, err := os.Create(privKeyFile)
	if err != nil {
		log.Println(err)
		return err
	}

	err = pem.Encode(file, block)
	if err != nil {
		log.Println(err)
		return err
	}

	pubKey := privKey.Public()

	pubKeyData, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		log.Println(err)
		return err
	}

	block = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubKeyData,
	}

	file, err = os.Create(pubKeyFile)
	if err != nil {
		log.Println(err)
		return err
	}

	err = pem.Encode(file, block)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func main() {
	err := RSA_GenerateKeyPair([]byte("123456"), "priv.pem", "pub.pem")
	if err != nil {
		log.Println(err)
		return
	}

	privatekey, err := ioutil.ReadFile("priv.pem")
	if err != nil {
		log.Println(err)
		return
	}

	publickey, err := ioutil.ReadFile("pub.pem")
	if err != nil {
		log.Println(err)
		return
	}

	msg := "RSA MESSAGE"

	ciphertext, err := RSA_Encrypt(publickey, []byte(msg))
	if err != nil {
		log.Println(err)
		return
	}

	plaintext, err := RSA_Decrypt(privatekey, ciphertext)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("message =", msg)
	fmt.Println("Encrypt message =", ciphertext)
	fmt.Println("Decrypt message =", string(plaintext))

	msg = "RSA SIGN MESSAGE"
	sig, err := RSA_Sign(privatekey, []byte(msg))
	if err != nil {
		log.Println(err)
		return
	}

	v, err := RSA_Verify(publickey, []byte(msg), sig)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("message =", msg)
	fmt.Println("sig message =", sig)
	fmt.Println("verify message =", v)
}
