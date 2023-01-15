package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/open-quantum-safe/liboqs-go/oqs"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// set up liboqs

	fmt.Println("This is an experimental mini wallet. \nIt uses Falcon-512 with NIST round 2 standartisation level. \nDo not use in production.")
	sigName := "Falcon-512"
	signer := oqs.Signature{}
	defer signer.Clean() // clean up even in case of panic
	i := 0
setup:
	for i < 3 {

		// print menu options
		fmt.Println("1. Generate a new address")
		fmt.Println("2. Import an address")
		fmt.Println("3. Exit")

		// read user input
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your choice: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// convert input to integer
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		// initiate signer with a fresh keypair or an imported private key
		switch choice {
		case 1:
			if err := signer.Init(sigName, nil); err != nil {
				log.Fatal(err)
			}
			pubKey, err := signer.GenerateKeyPair()
			if err != nil {
				log.Fatal(err)
			}
			hexPubKey := hex.EncodeToString(pubKey)
			fmt.Println("Signer public key: ", hexPubKey)
			break setup
		case 2:
			println("Address import will be supported in the next update.")
			//        fmt.Print("Enter your private key: ")
			//        secretKey, _ := reader.ReadString('\n')
			//        secretKey = strings.TrimSpace(input)
			//        secretKeyBytes = []byte(secretKey)
			//
			//      if err := signer.Init(sigName, secretKeyBytes); err != nil {
			//           log.Fatal(err)
			//    }

		case 3:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 3.")
			continue

		}
	}
	// main loop
	for {
		// print menu options
		fmt.Println("1. Sign a message")
		fmt.Println("2. Verify a message")
		fmt.Println("3. Export private key")
		fmt.Println("4. Exit")

		// read user input
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your choice: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// convert input to integer
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		// execute function based on user input
		switch choice {
		case 1:
			message := generateMessage()
			signature, _ := signer.Sign(message)
			strSignature := hex.EncodeToString(signature)
			fmt.Println("Your signatur is: ", strSignature)
		case 2:
			message := generateMessage()

			verifier := oqs.Signature{}
			defer verifier.Clean() // clean up even in case of panic

			if err := verifier.Init(sigName, nil); err != nil {
				log.Fatal(err)
			}
			signature := getSignature()
			pubKey := getPubKey()
			isValid, err := verifier.Verify(message, signature, pubKey)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("\nValid signature?", isValid)

		case 3:
			secretKey := signer.ExportSecretKey()
			strSecretKey := hex.EncodeToString(secretKey)
			fmt.Println("Your secret Key is: ", strSecretKey)
		case 4:
			os.Exit(0)

		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
		}
	}
}

func getPubKey() []byte {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a Falcon-512 public key: ")
	pubKey, _ := reader.ReadString('\n')
	pubKey = strings.TrimSpace(pubKey)
	bytePubKey, _ := hex.DecodeString(pubKey)
	return bytePubKey
}

func getSignature() []byte {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a signature to be verified: ")
	signature, _ := reader.ReadString('\n')
	signature = strings.TrimSpace(signature)
	byteSignature, _ := hex.DecodeString(signature)
	return byteSignature
}

func generateMessage() []byte {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a message: ")
	message, _ := reader.ReadString('\n')
	message = strings.TrimSpace(message)

	// Hash the message
	hashedMessage := crypto.Keccak256([]byte(message))
	// The prefix will look in solidity like: "\x19Lattice Signed Message:\n32")
	prefix := []byte("Lattice Signed Message:")
	finalMessage := crypto.Keccak256(append(prefix, hashedMessage...))
	return finalMessage
}
