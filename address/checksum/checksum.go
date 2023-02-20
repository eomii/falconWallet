package checksum

import(
        "golang.org/x/crypto/sha3"
        "strings"
)

func ValidChecker(addr string) bool { //Takes a string address as input
        addr_low := strings.ToLower(addr)[2:]

        //Treat the hex address as ascii/utf-8 for keccak256 hashing
        hasher := sha3.NewLegacyKeccak256()
        hasher.Write([]byte(addr_low))
        hash := hasher.Sum(nil)

        prefix := "0x"

        for i, b := range addr_low {
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

func AddressChecksum(addr string) string { //Takes a string address as input
        addr_low := strings.ToLower(addr)[2:]

        //Treat the hex address as ascii/utf-8 for keccak256 hashing
        hasher := sha3.NewLegacyKeccak256()
        hasher.Write([]byte(addr_low))
        hash := hasher.Sum(nil)

        prefix := "0x"

        for i, b := range addr_low {
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
