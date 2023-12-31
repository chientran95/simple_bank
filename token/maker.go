package token

import "time"

type Maker interface {
	// Create a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, *Payload, error)

	// Checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
