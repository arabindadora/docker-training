package main

import (
    "crypto/sha256"
    "encoding/base64"
    "fmt"
)

func GetSHA256(text string) string {
    hash := sha256.Sum256([]byte(text))
    return fmt.Sprintf("%x", hash)
}

func GetBase64(text string) string {
    return base64.StdEncoding.EncodeToString([]byte(text))
}
