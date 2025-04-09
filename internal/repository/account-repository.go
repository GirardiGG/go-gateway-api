package repository

import (
	"database/sql"
	"time"

	"github.com/devfullcycle/imersao22/go-gateway/internal/domain"
)

type AccountRepository struct {
	db *sql.DB
}

func CreateAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}
func (repository *AccountRepository) SaveAccount(account *domain.Account) error {
	statement, err := repository.db.Prepare("INSERT INTO accounts (id, name, email, balance, api_key, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(
		account.ID,
		account.Name,
		account.Email,
		account.Balance,
		account.APIKey,
		account.CreatedAt,
		account.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repository *AccountRepository) FindAccountByAPIKey(apiKey string) (*domain.Account, error) {
	var account domain.Account
	var CreatedAt, UpdatedAt time.Time

	err := repository.db.QueryRow(`
		SELECT id, name, email, api_key, balance, created_at, updated_at
		FROM accounts
		WHERE api_key = $1
	`, apiKey).Scan( // Acessa a variavel direto na memoria e subistitui as variaveis direto na memoria
		&account.ID,
		&account.Name,
		&account.Email,
		&account.APIKey,
		&account.Balance,
		&CreatedAt,
		&UpdatedAt,
	)

	if err == sql.ErrNoRows { // erro de busca sem resultado
		return nil, domain.ErrAccountNotFound
	}
	if err != nil {
		return nil, err
	}
	account.CreatedAt = CreatedAt
	account.UpdatedAt = UpdatedAt
	return &account, nil
}

func (repository *AccountRepository) FindAccountByID(id string) (*domain.Account, error) {
	var account domain.Account
	var CreatedAt, UpdatedAt time.Time

	err := repository.db.QueryRow(`
		SELECT id, name, email, api_key, balance, created_at, updated_at
		FROM accounts
		WHERE id = $1
	`, id).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.APIKey,
		&account.Balance,
		&CreatedAt,
		&UpdatedAt,
	)

	if err == sql.ErrNoRows { // erro de busca sem resultado
		return nil, domain.ErrAccountNotFound
	}
	if err != nil {
		return nil, err
	}
	account.CreatedAt = CreatedAt
	account.UpdatedAt = UpdatedAt
	return &account, nil
}

func (repository *AccountRepository) UpdateBalance(account *domain.Account) error {
	tx, err := repository.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var currentBalance float64
	err = tx.QueryRow("SELECT balance FROM accounts WHERE id = $1 FOR UPDATE", account.ID).Scan(&currentBalance)

	if err == sql.ErrNoRows {
		return domain.ErrAccountNotFound
	}
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE accounts SET balance = $1, updated_at = $2 WHERE id = $3", account.Balance, time.Now(), account.ID)
	if err != nil {
		return err
	}
	return tx.Commit()
}
