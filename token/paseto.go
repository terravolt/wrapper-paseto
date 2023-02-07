package token

import (
	"errors"
	"fmt"
	"github.com/aead/chacha20poly1305"
	"github.com/gofrs/uuid"
	"github.com/o1egl/paseto"
	"time"
)

// PasetoMaker is a struct that implements PASETO token generation and verification.
// It contains the paseto instance and the symmetric key used for encryption and decryption.
type PasetoMaker struct {
	// paseto is an instance of PASETO v2 implementation
	paseto *paseto.V2
	// symmetricKey is a secret key used for encrypting and decrypting the tokens
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	// Check if the key size is exactly equal to the required size
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}
	// Initialize a new PasetoMaker with the provided symmetric key
	maker := &PasetoMaker{
		// Create a new V2 paseto instance
		paseto: paseto.NewV2(),
		// Store the symmetric key as a byte slice
		symmetricKey: []byte(symmetricKey),
	}
	return maker, nil
}

// CreateToken creates a new token for a specific userID and duration
func (maker *PasetoMaker) CreateToken(userID uuid.UUID, duration time.Duration) (string, *Payload, error) {
	// Create a new Payload with the userID and duration
	payload, err := NewPayload(userID, duration)
	if err != nil {
		// Return an error if there was an issue creating the Payload
		return "", payload, err
	}
	// Encrypt the payload using the paseto package and the symmetric key
	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	// Return the encrypted token and the Payload
	return token, payload, err
}

// VerifyToken checks the validity of the given token
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	// Initialize a Payload struct to store the decrypted data
	payload := &Payload{}

	// Decrypt the token using the paseto library
	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)

	// If the decryption fails, return an error
	if err != nil {
		return nil, errors.New("token is invalid")
	}

	// Check if the token has expired
	if time.Now().After(payload.ExpiredAt) {
		return nil, errors.New("token has expired")
	}

	// Return the payload and a nil error if the token is valid
	return payload, nil
}
