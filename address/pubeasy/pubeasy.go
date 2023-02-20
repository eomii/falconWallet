package pubeasy

import (
	"github.com/ethereum/go-ethereum/crypto"
)

func PubToAddress(pubKey []byte) ([]byte, error) {
	return crypto.Keccak256(pubKey)[12:], nil

}
