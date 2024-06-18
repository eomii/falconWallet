package address

import (
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"hash"
)

const (
	// HashLength is the expected length of the hash
	HashLength = 32
	// AddressLength is the expected length of the address
	AddressLength = 20
)

// PubToAddress converts a public key to an Ethereum address.
func PubToAddress(pubKey []byte) ([]byte, error) {
	return crypto.Keccak256(pubKey)[12:], nil
}

// PubkeyToAddress converts a public key to an Ethereum address.
func PubkeyToAddress(p string) Address {
	pubBytes := []byte(p)
	return BytesToAddress(Keccak256(pubBytes[1:])[12:])
}

// Address represents the 20 byte address of an Ethereum account.
type Address [AddressLength]byte

// BytesToAddress returns Address with value b.
func BytesToAddress(b []byte) Address {
	var a Address
	a.SetBytes(b)
	return a
}

// SetBytes sets the address to the value of b.
func (a *Address) SetBytes(b []byte) {
	if len(b) > len(a) {
		b = b[len(b)-AddressLength:]
	}
	copy(a[AddressLength-len(b):], b)
}

// KeccakState interface represents a Keccak state.
type KeccakState interface {
	hash.Hash
	Read([]byte) (int, error)
}

// NewKeccakState creates a new KeccakState.
func NewKeccakState() KeccakState {
	return sha3.NewLegacyKeccak256().(KeccakState)
}

// HashData hashes the provided data using the KeccakState and returns a 32 byte hash.
func HashData(kh KeccakState, data []byte) (h Hash) {
	kh.Reset()
	kh.Write(data)
	kh.Read(h[:])
	return h
}

// Keccak256 calculates and returns the Keccak256 hash of the input data.
func Keccak256(data ...[]byte) []byte {
	b := make([]byte, 32)
	d := NewKeccakState()
	for _, b := range data {
		d.Write(b)
	}
	d.Read(b)
	return b
}

// Keccak256Hash calculates and returns the Keccak256 hash of the input data.
func Keccak256Hash(data ...[]byte) (h Hash) {
	d := NewKeccakState()
	for _, b := range data {
		d.Write(b)
	}
	d.Read(h[:])
	return h
}

// Keccak512 calculates and returns the Keccak512 hash of the input data.
func Keccak512(data ...[]byte) []byte {
	d := sha3.NewLegacyKeccak512()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}

// Hash represents a 32 byte Keccak hash.
type Hash [HashLength]byte

// BytesToHash sets b to hash.
func BytesToHash(b []byte) Hash {
	var h Hash
	h.SetBytes(b)
	return h
}

// SetBytes sets the hash to the value of b.
func (h *Hash) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-HashLength:]
	}
	copy(h[HashLength-len(b):], b)
}
