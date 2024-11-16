package main

import (
	"fmt"
	btcec "github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"log"
)

func ECC_Encrypt(pubKeyBytes, message []byte) ([]byte, error) {
	pubKey, err := btcec.ParsePubKey(pubKeyBytes, btcec.S256())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	ciphertext, err := btcec.Encrypt(pubKey, message)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return ciphertext, nil
}

func ECC_Decrypt(privKeyBytes, ciphertext []byte) ([]byte, error) {
	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), privKeyBytes)

	plaintext, err := btcec.Decrypt(privKey, ciphertext)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return plaintext, nil
}

func ECC_Sign(privKeyBytes, message []byte) ([]byte, error) {
	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), privKeyBytes)

	messageHash := chainhash.DoubleHashB(message)

	signature, err := privKey.Sign(messageHash)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return signature.Serialize(), nil
}

func ECC_Verify(pubKeyBytes, message, sigBytes []byte) (bool, error) {
	pubKey, err := btcec.ParsePubKey(pubKeyBytes, btcec.S256())
	if err != nil {
		log.Println(err)
		return false, err
	}

	signature, err := btcec.ParseSignature(sigBytes, btcec.S256())
	if err != nil {
		log.Println(err)
		return false, err
	}

	messageHash := chainhash.DoubleHashB(message)

	verified := signature.Verify(messageHash, pubKey)

	return verified, nil
}

func ECC_GetKeyPair() ([]byte, []byte, error) {
	privKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	pubKey := privKey.PubKey()

	return privKey.Serialize(), pubKey.SerializeCompressed(), nil
}

func main() {
	message := "Hello, ECC"

	privKey, pubKey, err := ECC_GetKeyPair()
	if err != nil {
		log.Println(err)
		return
	}

	ciphertext1, err := ECC_Encrypt(pubKey, []byte(message))
	if err != nil {
		log.Println(err)
		return
	}

	ciphertext2, err := ECC_Encrypt(pubKey, []byte(message))
	if err != nil {
		log.Println(err)
		return
	}

	plaintext1, err := ECC_Decrypt(privKey, ciphertext1)
	if err != nil {
		log.Println(err)
		return
	}

	plaintext2, err := ECC_Decrypt(privKey, ciphertext2)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("message =", message)
	fmt.Println("Encrypt message1 =", ciphertext1)
	fmt.Println("Encrypt message2 =", ciphertext2)
	fmt.Println("Decrypt message1 =", string(plaintext1))
	fmt.Println("Decrypt message2 =", string(plaintext2))

	message = "Hello, ECC_Sign"
	sig, err := ECC_Sign(privKey, []byte("Hello, ECC_Sign"))
	if err != nil {
		log.Println(err)
		return
	}

	verify, err := ECC_Verify(pubKey, []byte(message), sig)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("message =", message)
	fmt.Println("sig message =", sig)
	fmt.Println("verify message =", verify)
}
