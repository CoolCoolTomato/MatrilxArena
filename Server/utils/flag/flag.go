package flag

import (
    "crypto/rand"
    "fmt"
    "math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

func GenerateFlag() string {
    result := make([]byte, 32)
    for i := range result {
        randomInt, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
        if err != nil {
            return ""
        }
        result[i] = charset[randomInt.Int64()]
    }
    return fmt.Sprintf("flag{%s}", result)
}
