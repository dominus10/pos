package security

import (
	"sync"
	"time"
)

var (
	nonceStore = make(map[string]time.Time)
	mu         sync.Mutex
)

// ValidateNonce checks if a nonce is valid
func ValidateNonce(nonce string) bool {
	mu.Lock()
	defer mu.Unlock()

	// If nonce exists, reject it (replay detected)
	if _, exists := nonceStore[nonce]; exists {
		return false
	}

	// Store nonce with expiration time (5 minutes)
	nonceStore[nonce] = time.Now().Add(5 * time.Minute)

	return true
}

// CleanupExpiredNonces clears old nonces periodically
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
