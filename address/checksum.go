package address

import (
	"strings"

	"golang.org/x/crypto/sha3"
)

// ValidChecker checks if the address is valid.
func ValidChecker(addr string) bool {
	addrLow := strings.ToLower(addr)[2:]

	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte(addrLow))
	hash := hasher.Sum(nil)

	prefix := "0x"

	for i, b := range addrLow {
		c := string(b)
		if b < '0' || b > '9' {
			if hash[i/2]&byte(128-i%2*120) != 0 {
				c = string(b - 32)
			}
		}
		prefix += c
	}
	return addr == prefix
}

// AddressChecksum computes the checksum of the address.
func AddressChecksum(addr string) string {
	addrLow := strings.ToLower(addr)[2:]

	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte(addrLow))
	hash := hasher.Sum(nil)

	prefix := "0x"

	for i, b := range addrLow {
		c := string(b)
		if b < '0' || b > '9' {
			if hash[i/2]&byte(128-i%2*120) != 0 {
				c = string(b - 32)
			}
		}
		prefix += c
	}
	return prefix
}
