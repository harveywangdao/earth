package main

import (
	"crypto/dsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
)

func DSA_Sign(privkey *dsa.PrivateKey, msg []byte) (*big.Int, *big.Int, error) {
	msgHash := sha256.Sum256(msg)
	r, s, err := dsa.Sign(rand.Reader, privkey, msgHash[:])
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	return r, s, nil
}

func DSA_Verify(pubkey *dsa.PublicKey, msg []byte, r, s *big.Int) bool {
	msgHash := sha256.Sum256(msg)
	return dsa.Verify(pubkey, msgHash[:], r, s)
}

func DSA_GenerateKeyPair() (*dsa.PrivateKey, error) {
	priv := &dsa.PrivateKey{}
	err := dsa.GenerateParameters(&priv.Parameters, rand.Reader, dsa.L1024N160)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = dsa.GenerateKey(priv, rand.Reader)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return priv, nil
}

func main() {
	privKey, err := DSA_GenerateKeyPair()
	if err != nil {
		log.Println(err)
		return
	}

	msg := "DSA SIGN MESSAGE"
	r, s, err := DSA_Sign(privKey, []byte(msg))
	if err != nil {
		log.Println(err)
		return
	}

	v := DSA_Verify(&privKey.PublicKey, []byte(msg), r, s)

	fmt.Println("message =", msg)
	fmt.Println("sig message =", r, s)
	fmt.Println("verify message =", v)
}
