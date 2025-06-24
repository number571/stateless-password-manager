package main

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"os"

	gorfc1751 "github.com/number571/go-rfc1751"
	"golang.org/x/crypto/scrypt"
)

func main() {
	var (
		flagV = flag.Bool("v", false, "print version of application")
		flagM = flag.Uint64("m", 0, "create master key with a bit-size")
		flagT = flag.String("t", "", "name of the target service")
	)

	flag.Parse()

	if *flagV {
		fmt.Println("v1.0.1")
		return
	}

	if bits := *flagM; bits != 0 {
		fmt.Println(generateMasterKey(bits))
		return
	}

	if t := *flagT; t != "" {
		fmt.Print("Master-Key: ")
		mk := loadMasterKey(readStdinUntilEOL())
		fmt.Println("Please wait a few seconds...")
		fmt.Println("Password:", generatePassword(mk, t))
		return
	}

	flag.Usage()
}

func loadMasterKey(s string) []byte {
	b, err := gorfc1751.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}

func generateMasterKey(bits uint64) string {
	newMK, err := gorfc1751.NewMnemonic(rand.Reader, bits)
	if err != nil {
		panic(err)
	}
	return newMK
}

func generatePassword(mk []byte, t string) string {
	key, err := scrypt.Key(mk, []byte(t), 1<<20, 8, 1, 32)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(key)
}

func readStdinUntilEOL() string {
	res, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		panic(err)
	}
	return string(res)
}
