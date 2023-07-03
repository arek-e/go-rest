package repositories

import (
	"errors"

	"github.com/arek-e/lanexpense/domain"
)

type AccountRepository struct {
	DB map[string]domain.Account
}

func NewAccountRepository() *AccountRepository {
	example := domain.Account{
		ID:            "1",
		Name:          "RepoTest",
		AccountType:   nil,
		AccountNumber: nil,
		StartBalance:  nil,
	}

	return &AccountRepository{
		DB: map[string]domain.Account{example.ID: example},
	}
}

func (repo *AccountRepository) GetAllAccounts() ([]domain.Account, error) {
	var accounts []domain.Account

	for _, account := range repo.DB {
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (repo *AccountRepository) GetAccount(id string) (domain.Account, error) {
	if account, exist := repo.DB[id]; exist {
		return account, nil
	}

	return domain.Account{}, errors.New("Account not found")
}

func (repo *AccountRepository) CreateAccount(account domain.Account) (domain.Account, error) {
	if _, exist := repo.DB[account.ID]; exist {
		return domain.Account{}, errors.New("Already exist")
	}

	repo.DB[account.ID] = account

	return account, nil
}

func (repo *AccountRepository) UpdateAccount(account domain.Account) error {
	if _, exist := repo.DB[account.ID]; !exist {
		return errors.New("Not found")
	}

	repo.DB[account.ID] = account

	return nil
}

func (repo *AccountRepository) DeleteAccount(id string) error {
	if _, exist := repo.DB[id]; !exist {
		return errors.New("Not found")
	}

	delete(repo.DB, id)

	return nil
}
