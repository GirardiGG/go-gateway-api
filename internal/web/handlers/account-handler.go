package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/devfullcycle/imersao22/go-gateway/internal/dto"
	"github.com/devfullcycle/imersao22/go-gateway/internal/service"
)

type AccountHandler struct {
	accountService *service.AccountService
}

// constructor function
// CreateAccountHandler initializes a new AccountHandler with the provided AccountService.
func CreateAccountHandler(accountService *service.AccountService) *AccountHandler {
	return &AccountHandler{accountService: accountService}
}

func (handler *AccountHandler) CreateAccount(writer http.ResponseWriter, request *http.Request) {
	var input dto.CreateAccountInputDTO

	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := handler.accountService.CreateAccount(input)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(output)
}

func (handler *AccountHandler) GetAccounts(writer http.ResponseWriter, request *http.Request) {
	apiKey := request.Header.Get("x-api-key")
	if apiKey == "" {
		http.Error(writer, "API key is required", http.StatusUnauthorized)
		return
	}
	output, err := handler.accountService.FindAccountByAPIKey(apiKey)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(output)
}
