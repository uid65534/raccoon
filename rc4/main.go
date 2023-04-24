package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
	"crypto/rc4"
)

var (
	decodeBase64 bool
	key string
)

func printUsage() {
	fmt.Fprintf(os.Stderr, "Command-line RC4 stream cipher.\n")
	fmt.Fprintf(os.Stderr, "Usage: %s [-k key] [-b64] [data]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Reads from standard input if data is not provided.\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = printUsage
	flag.StringVar(&key, "k", "", "The key used to initialize the cipher state." +
		"\nWill be read from the RC4_KEY environment variable if not specified.")
	flag.BoolVar(&decodeBase64, "b64", false, "Decode input data from base64.")
	flag.Parse()

	args := flag.Args()
	if len(args) > 1 {
		flag.Usage()
		return
	}
	
	if len(key) == 0 {
		key = os.Getenv("RC4_KEY")
		if len(key) == 0 {
			flag.Usage()
			return
		}
	}

	var reader io.Reader
	if len(args) == 1 {
		var data []byte
		var err error
		if decodeBase64 {
			data, err = base64.StdEncoding.DecodeString(args[0])
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
				return
			}
		} else {
			data = []byte(args[0])
		}
		reader = bytes.NewReader(data)
	} else {
		reader = os.Stdin
	}

	cipher, err := rc4.NewCipher([]byte(key))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
		return
	}

	buffer := make([]byte, 4096)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
			}
			break
		}
		if n <= 0 {
			break
		}
		cipher.XORKeyStream(buffer[:n], buffer[:n])
		os.Stdout.Write(buffer[:n])
	}
}
