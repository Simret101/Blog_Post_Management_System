package rand

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func generateRandomSecret() string {
	secret := make([]byte, 32) // 32 bytes = 256 bits
	_, err := rand.Read(secret)
	if err != nil {
		panic("failed to generate random secret")
	}
	return base64.StdEncoding.EncodeToString(secret)
}

func main() {
	fmt.Println("AccessTokenSecret:", generateRandomSecret())
	fmt.Println("RefreshTokenSecret:", generateRandomSecret())
}
