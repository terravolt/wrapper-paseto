package token

import (
	"errors"
	"github.com/gofrs/uuid"
	"time"
)

// Payload contains the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Place     string    `json:"place"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(userID uuid.UUID, place string, duration time.Duration) (*Payload, error) {

	// Generates a new token ID using the UUID v7 format
	tokenID, err := uuid.NewV7()
	if err != nil {
		return nil, errors.New("failed to generate session ID")
	}

	// Create a new instance of the Payload struct
	payload := &Payload{
		ID:        tokenID,
		UserID:    userID,
		Place:     place,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	// Return the newly created Payload struct and a nil error
	return payload, nil
}
