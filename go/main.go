package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1]

	input := []byte(args)
	bytes, err := decodeHex(input)

	if err != nil  {
		fmt.Printf("Failed to decode hex", err)
		return
	}

	output := base64Encode(bytes)
	fmt.Printf("%s", output)
}

func decodeHex(input []byte) ([]byte, error) {
	bytes := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(bytes, input)

	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func base64Encode(input []byte) ([]byte) {
	encodedBytes := make([]byte, base64.StdEncoding.EncodedLen(len(input)))
	base64.StdEncoding.Encode(encodedBytes, input)

	return encodedBytes
}
