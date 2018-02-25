package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/iotaledger/giota"
)

const (
	index         int = 0
	securityLevel int = 2
)

func ReadBundle(name string) giota.Bundle {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var txns []giota.Transaction

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		trytes := giota.Trytes(scanner.Text())

		txn, err := giota.NewTransaction(trytes)
		if err != nil {
			panic(err)
		}

		txns = append(txns, *txn)
	}

	return giota.Bundle(txns)
}

func SignBundle(key giota.Trytes, bundle giota.Bundle) []giota.Trytes {
	normalizedBundleHash := bundle.Hash().Normalize()

	sign0 := giota.Sign(normalizedBundleHash[:27], key[:giota.SignatureSize/3])
	sign1 := giota.Sign(normalizedBundleHash[27:54], key[giota.SignatureSize/3:giota.SignatureSize*2/3])

	return []giota.Trytes{sign0, sign1}
}

func BundleTrytes(bundle giota.Bundle) giota.Trytes {
	trytesFrom := ""

	for _, txn := range bundle {
		trytesFrom += string(txn.Trytes())
	}

	return giota.Trytes(trytesFrom)
}

func Sha256Trytes(trytes giota.Trytes) string {
	hash := sha256.Sum256([]byte(trytes))

	return hex.EncodeToString(hash[:])
}

func VerifyBundle(name string, address giota.Address, signatureFragments []giota.Trytes, bundle giota.Bundle) {
	var validity string

	if giota.IsValidSig(address, signatureFragments, bundle.Hash()) {
		validity = "valid"
	} else {
		validity = "invalid"
	}

	fmt.Println()
	fmt.Printf("Verifying %s bundle:\n", name)
	fmt.Printf("SHA-256 hash is %s\n", Sha256Trytes(BundleTrytes(bundle)))
	fmt.Printf("Curl-P hash is %s\n", bundle.Hash())
	fmt.Printf("Signature is %s!\n", validity)
}

func main() {
	signedBundle := ReadBundle("signed_bundle.txt")
	forgedBundle := ReadBundle("forged_bundle.txt")

	seed := giota.NewSeed()
	fmt.Printf("Generated seed: %s\n", seed)

	key := giota.NewKey(seed, index, securityLevel)

	address, err := giota.NewAddress(seed, index, securityLevel)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Generated address: %s\n", address)

	fmt.Println("Signing bundle")
	signatureFragments := SignBundle(key, signedBundle)

	fmt.Printf("Fragment 0: %s\n", signatureFragments[0])
	fmt.Printf("Fragment 1: %s\n", signatureFragments[1])

	VerifyBundle("signed", address, signatureFragments, signedBundle)
	VerifyBundle("forged", address, signatureFragments, forgedBundle)
}
