package domain

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Name      string
	Email     string
	Balance   float64
	APIKey    string
	mu        sync.RWMutex // Mutex to protect concurrent access
	CreatedAt time.Time
	UpdatedAt time.Time
}

func generatAPIKey() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func CreateAccount(name, email string) *Account {
	return &Account{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Balance:   0.0,
		APIKey:    generatAPIKey(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (account *Account) UpdateBalance(amount float64) {
	account.mu.Lock()         // Lock the account for writing
	defer account.mu.Unlock() // Unlock the account, defer to ensure it happens after the function returns
	account.Balance += amount
	account.UpdatedAt = time.Now()
}
