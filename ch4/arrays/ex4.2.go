package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func mainE42() {

	var hash string
	flag.StringVar(&hash, "hash", "SHA256", "Available options: SHA256 (default), SHA384, SHA512")

	flag.Parse()
	args := flag.Args()
	fmt.Printf("Using %v checksum\n", hash)

	checksumFn, ok := map[string]func([]byte) []byte{
		"SHA256": func(b []byte) []byte {
			cs := sha256.Sum256(b)
			return cs[:]
		},
		"SHA384": func(b []byte) []byte {
			cs := sha512.Sum384(b)
			return cs[:]
		},
		"SHA512": func(b []byte) []byte {
			cs := sha512.Sum512(b)
			return cs[:]
		},
	}[hash]

	if !ok {
		fmt.Printf("Hash function: %v is not supported", hash)
		os.Exit(1)
	}

	for _, arg := range args {
		checksum := checksumFn([]byte(arg))
		fmt.Printf("%x\t%s\n", checksum, arg)
	}
}
