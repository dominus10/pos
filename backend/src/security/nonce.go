package security

import (
	"sync"
	"time"
)

var (
	nonceStore = make(map[string]time.Time)
	mu         sync.Mutex
)

// ValidateNonce checks if a nonce is valid and prevents replay attacks.
func ValidateNonce(nonce string) bool {
	mu.Lock()
	defer mu.Unlock()

	// Reject if the nonce was already used (Replay Attack Prevention)
	if _, exists := nonceStore[nonce]; exists {
		return false
	}

	// Store nonce with expiration (e.g., 5 min)
	nonceStore[nonce] = time.Now().Add(5 * time.Minute)

	return true
}

// CleanupExpiredNonces removes expired nonces
func CleanupExpiredNonces() {
	mu.Lock()
	defer mu.Unlock()

	now := time.Now()
	for nonce, expiry := range nonceStore {
		if now.After(expiry) {
			delete(nonceStore, nonce)
		}
	}
}
