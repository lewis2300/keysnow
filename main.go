package main

import (
	"fmt"
	"math/big"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

func main() {
	// Print header
	fmt.Printf("%64s %34s %34s\n", "Private", "Public", "Public Compressed")

	// Initialise big numbers with small numbers
	count, one := big.NewInt(0), big.NewInt(1)

	// Create a slice to pad our count to 32 bytes
	padded := make([]byte, 32)

	// Loop forever because we're never going to hit the end anyway
	for {
		// Increment our counter
		count.Add(count, one)

		// Copy count value's bytes to padded slice
		copy(padded[32-len(count.Bytes()):], count.Bytes())

		// Get public key
		privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), padded)
		pubKey := privKey.PubKey()

		// Get compressed and uncompressed addresses
		caddr, _ := btcutil.NewAddressPubKey(pubKey.SerializeCompressed(), &chaincfg.MainNetParams)
		uaddr, _ := btcutil.NewAddressPubKey(pubKey.SerializeUncompressed(), &chaincfg.MainNetParams)

		// Print keys
		fmt.Printf("%x %34s %34s\n", padded, uaddr.EncodeAddress(), caddr.EncodeAddress())
	}
}



