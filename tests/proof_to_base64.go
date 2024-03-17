package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	files, err := os.ReadDir("./testing_data")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), "proof") {
			continue
		}

		filePath := filepath.Join("./testing_data", file.Name())
		contents, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("Error reading file %s: %s\n", file.Name(), err)
			continue
		}

		encoded := make([]byte, base64.StdEncoding.EncodedLen(len(contents)))
		base64.StdEncoding.Encode(encoded, contents)

		outputFilePath := filepath.Join("./testing_data", strings.TrimSuffix(file.Name(), "proof")+"base64")
		err = os.WriteFile(outputFilePath, encoded, 0644)
		if err != nil {
			log.Printf("Error writing file %s: %s\n", outputFilePath, err)
			continue
		}
	}
}
