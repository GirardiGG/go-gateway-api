package service

import (
	"github.com/devfullcycle/imersao22/go-gateway/internal/domain"
	"github.com/devfullcycle/imersao22/go-gateway/internal/dto"
)

type AccountService struct {
	repository domain.AccountRepository
}

func CreateAccountService(repository domain.AccountRepository) *AccountService {
	return &AccountService{repository: repository}
}

func (accountService *AccountService) CreateAccount(input dto.CreateAccountInputDTO) (*dto.AccountOutputDTO, error) {
	account := dto.ToAccount(input)
	existAccount, err := accountService.repository.FindAccountByAPIKey(account.APIKey)
	if err != nil && err != domain.ErrAccountNotFound {
		return nil, err
	}
	if existAccount != nil {
		return nil, domain.ErrDuplicateAPIKey
	}

	err = accountService.repository.SaveAccount(account)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return &output, nil
}

func (accountService *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountOutputDTO, error) {
	account, err := accountService.repository.FindAccountByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}
	account.UpdateBalance(amount)
	err = accountService.repository.UpdateBalance(account)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return &output, nil
}

func (accountService *AccountService) FindAccountByAPIKey(apiKey string) (*dto.AccountOutputDTO, error) {
	account, err := accountService.repository.FindAccountByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}
	output := dto.FromAccount(account)
	return &output, nil
}

func (accountService *AccountService) FindAccountByID(id string) (*dto.AccountOutputDTO, error) {
	account, err := accountService.repository.FindAccountByID(id)
	if err != nil {
		return nil, err
	}
	output := dto.FromAccount(account)
	return &output, nil
}
