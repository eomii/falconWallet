package main

import (
	"bufio"
	"encoding/hex"
	"falconWallet/address"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/open-quantum-safe/liboqs-go/oqs"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("This is an experimental mini-wallet using Falcon-512 with NIST round 2 standardization. Do not use in production.")

	const sigName = "Falcon-512"
	var signer oqs.Signature
	defer signer.Clean()

	for {
		fmt.Println("1. Generate a new address")
		fmt.Println("2. Import an address")
		fmt.Println("3. Exit")

		choice, err := getUserChoice()
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

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
			addr, err := address.PubToAddress(pubKey)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Signer Address: ", hex.EncodeToString(addr))
			goto mainLoop
		case 2:
			fmt.Println("Address import will be supported in the next update.")
		case 3:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 3.")
		}
	}

mainLoop:
	for {
		fmt.Println("1. Sign a message")
		fmt.Println("2. Verify a message")
		fmt.Println("3. Export private key")
		fmt.Println("4. Exit")

		choice, err := getUserChoice()
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			message := generateMessage()
			signature, err := signer.Sign(message)
			if err != nil {
				log.Fatal(err)
			}
			strSignature := hex.EncodeToString(signature)
			fmt.Println("Your signature is: ", strSignature)
		case 2:
			message := generateMessage()
			verifier := oqs.Signature{}
			defer verifier.Clean()

			if err := verifier.Init(sigName, nil); err != nil {
				log.Fatal(err)
			}
			signature := getSignature()
			pubKey := getPubKey()
			isValid, err := verifier.Verify(message, signature, pubKey)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Valid signature?", isValid)
		case 3:
			secretKey := signer.ExportSecretKey()
			strSecretKey := hex.EncodeToString(secretKey)
			fmt.Println("Your secret key is: ", strSecretKey)
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
		}
	}
}

func getUserChoice() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your choice: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	return strconv.Atoi(input)
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

	hashedMessage := crypto.Keccak256([]byte(message))
	prefix := []byte("Lattice Signed Message:")
	finalMessage := crypto.Keccak256(append(prefix, hashedMessage...))
	return finalMessage
}
