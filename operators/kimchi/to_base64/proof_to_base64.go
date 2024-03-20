package main

import (
	"encoding/base64"
	"log"
	"os"
)

func main() {
	contents, err := os.ReadFile("../lib/kimchi_ec_add.proof")
	if err != nil {
		log.Printf("Error reading file")
	}

	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(contents)))
	base64.StdEncoding.Encode(encoded, contents)

	err = os.WriteFile("../lib/proof_base64", encoded, 0644)

	if err != nil {
		log.Printf("Error writing to file")
	}
}
