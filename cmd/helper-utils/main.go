package main

import (
	"fmt"
	"github.com/rahul-as-dev/go-helper-utils/cmd/security"
	"log"
)

func main() {
	fmt.Printf("Hello World\n")
	//utils.ConvertIntToBits()
	//utils.ConvertBitsToInt()
	//tree := utils.NewBinaryTree(10)
	//tree.Left = utils.NewBinaryTree(5)
	//tree.Right = utils.NewBinaryTree(15)
	//fmt.Println(tree.NodeCount())
	encryptedToken := "wdPeMY9/xi6ieRfYOW11cF9IOqBDMNraML5WK9huZQ611zT6Q+bfOjIs2R3aiGk8NzqvvSshKaB42DfGXUlviN0sU8Dk/MJu7/+0CyDUbvjmJaxTaUL7MHw5g8bLZkemQcSfHVVXouNmakvhphP9Kg=="
	ivBase64 := "00c5a4f64285ee4a38173a14e9fb6019"
	keyBase64 := "30ee197ffc44308c"

	decryptedToken, err := security.Decrypt(encryptedToken, ivBase64, keyBase64)
	if err != nil {
		log.Fatalf("Decryption failed: %v", err)
	}

	fmt.Println("Decrypted token:", decryptedToken)
}
