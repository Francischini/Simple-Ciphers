package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ASCII_START uint8 = 65 // uppercase A
var ASCII_END uint8 = 90   // uppercase Z

func caesar(rawStringUpper string, shiftKey int) string {

	var cipher string
	cipherSlice := make([]byte, len(rawStringUpper))

	for i := 0; i < len(rawStringUpper); i++ {
		if rawStringUpper[i] < ASCII_START || rawStringUpper[i] > ASCII_END {
			cipherSlice[i] = rawStringUpper[i]
		} else {
			b := byte(shiftKey)
			cipherSlice[i] = ((rawStringUpper[i] - ASCII_START + b) % 26) + ASCII_START
		}
	}

	cipher = string(cipherSlice)
	return cipher
}

func vigenere(rawStringUpper string, shiftKey string) string {

	var cipher string
	cipherSlice := make([]byte, len(rawStringUpper))
	z := 0

	for i := 0; i < len(rawStringUpper); i++ {
		if z == len(shiftKey) {
			z = 0
		}
		if rawStringUpper[i] < ASCII_START || rawStringUpper[i] > ASCII_END {
			cipherSlice[i] = rawStringUpper[i]
		} else {
			cipherSlice[i] = (rawStringUpper[i]-ASCII_START+byte(shiftKey[z])-ASCII_START)%26 + ASCII_START
			z = z + 1
		}
	}

	cipher = string(cipherSlice)
	return cipher
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	var op int
	fmt.Print("Press 1 for Caesar Cipher or Press 2 for Vigenere Cipher: ")
	input, err := reader.ReadString('\n')
	checkError(err)
	op, err = strconv.Atoi(strings.TrimSpace(input))
	checkError(err)

	var rawString string
	fmt.Print("Enter the raw string: ")
	rawString, err = reader.ReadString('\n')
	checkError(err)

	rawString = strings.TrimSpace(rawString)
	rawStringUpper := strings.ToUpper(rawString)

	var cipher string

	if op == 1 {
		var shift int
		fmt.Print("Enter the shift amount: ")
		input, err := reader.ReadString('\n')
		checkError(err)

		shift, err = strconv.Atoi(strings.TrimSpace(input))
		checkError(err)

		cipher = caesar(rawStringUpper, shift)
	}
	if op == 2 {
		var shiftString string
		fmt.Print("Enter the shift string: ")
		shiftString, err = reader.ReadString('\n')
		checkError(err)
		shiftString = strings.TrimSpace(shiftString)
		shiftStringUpper := strings.ToUpper(shiftString)
		cipher = vigenere(rawStringUpper, shiftStringUpper)
	}

	fmt.Print("Cipher:\n", cipher)

}
