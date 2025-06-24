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
	newMK := flag.Uint64("new-mk", 0, "create master key with bit-size")
	target := flag.String("target", "", "password for the target service")

	flag.Parse()

	if bits := *newMK; bits != 0 {
		fmt.Println(generateMasterKey(bits))
		return
	}

	if t := *target; t != "" {
		fmt.Print("Master-Key: ")
		mk := readStdinUntilEOL()
		fmt.Println("Please wait a few seconds...")
		fmt.Println("Password:", generatePassword(mk, t))
		return
	}

	panic("target is null")
}

func generateMasterKey(bits uint64) string {
	newMK, err := gorfc1751.NewMnemonic(rand.Reader, bits)
	if err != nil {
		panic(err)
	}
	return newMK
}

func generatePassword(mk, t string) string {
	key, err := scrypt.Key([]byte(mk), []byte(t), 1<<20, 8, 1, 32)
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
