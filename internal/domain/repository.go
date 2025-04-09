package domain

type AccountRepository interface {
	SaveAccount(account *Account) error
	FindAccountByID(id string) (*Account, error)
	FindAccountByAPIKey(apiKey string) (*Account, error)
	UpdateBalance(account *Account) error
}
