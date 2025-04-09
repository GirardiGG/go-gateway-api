package dto

import "github.com/devfullcycle/imersao22/go-gateway/internal/domain"

type CreateAccountInputDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AccountOutputDTO struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Balance   float64 `json:"balance"`
	APIKey    string  `json:"api_key,omitempty"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

func ToAccount(input CreateAccountInputDTO) *domain.Account {
	return domain.CreateAccount(input.Name, input.Email)
}

func FromAccount(account *domain.Account) AccountOutputDTO {
	return AccountOutputDTO{
		ID:        account.ID,
		Name:      account.Name,
		Email:     account.Email,
		Balance:   account.Balance,
		APIKey:    account.APIKey,
		CreatedAt: account.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: account.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
