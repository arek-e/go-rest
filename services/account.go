package services

import (
	"github.com/google/uuid"

	"github.com/arek-e/lanexpense/domain"
	"github.com/arek-e/lanexpense/repositories"
)

// AccountService
type AccountService struct {
	Repository *repositories.AccountRepository
}

// NewServices
func NewServices() *AccountService {
	return &AccountService{
		Repository: repositories.NewAccountRepository(),
	}
}

// GetAllAccounts
func (s *AccountService) GetAllAccounts() ([]domain.Account, error) {
	return s.Repository.GetAllAccounts()
}

// GetAccount
func (s *AccountService) GetAccount(id string) (domain.Account, error) {
	return s.Repository.GetAccount(id)
}

// CreateAccount
func (s *AccountService) CreateAccount(account domain.Account) (domain.Account, int, error) {
	account.ID = uuid.New().String()

	account, err := s.Repository.CreateAccount(account)
	if err != nil {
		return domain.Account{}, 400, err
	}

	return account, 201, nil
}

// UpdateAccount
func (s *AccountService) UpdateAccount(account domain.Account) (domain.Account, error) {
	if err := s.Repository.UpdateAccount(account); err != nil {
		return domain.Account{}, err
	}

	return account, nil
}

// DeleteAccount
func (s *AccountService) DeleteAccount(id string) error {
	if err := s.Repository.DeleteAccount(id); err != nil {
		return err
	}

	return nil
}
