package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	proof, _ := os.ReadFile("./testing_data/sp1_fibonacci.b64")
	proof_buf := make([]byte, base64.StdEncoding.EncodedLen(len(proof)))
	base64.StdEncoding.Encode(proof_buf, proof)
	os.WriteFile("testing_data/sp1_fibonacci.b64", proof_buf, 0644)
}
