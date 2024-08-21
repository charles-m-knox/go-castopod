package castopod

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"log"
)

const (
	// Used for generating [RandomAlphanumeric] characters.
	AlphanumericCharacters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// Precalculated length of [AlphanumericCharacters] as a constant for the sake of efficiency.
	AlphaNumericCharactersLength = 62
)

// RandomAlphanumeric returns a random number of characters up to the provided
// length. See [AlphanumericCharacters] for the selection of characters that
// can appear.
func RandomAlphanumeric(length int) ([]byte, error) {
	r := []byte{}

	// var sb strings.Builder
	for i := 0; i < length; i++ {
		b := make([]byte, 1)
		_, err := rand.Read(b)
		if err != nil {
			return r, err
		}

		// returns a value between 0 and 255
		randomness := b[0]

		// retrieves a value between 0 and 61 at most (length of the available
		// characters from above)
		index := int(randomness) % AlphaNumericCharactersLength

		// pick a random character from the character selection
		char := AlphanumericCharacters[index]

		r = append(r, char)

		// sb.WriteByte(chars[int(b[0])%len(chars)])
	}
	return r, nil
}

// NewToken takes a sha256 of 8 alphanumeric characters (upper and lowercase
// OK). This is what Castopod uses for its token generator.
func NewToken() string {
	r, err := RandomAlphanumeric(8)
	if err != nil {
		log.Fatalf("failed to get random number: %v", err.Error())
	}

	h := sha256.Sum256(r)

	return hex.EncodeToString(h[:])
}
