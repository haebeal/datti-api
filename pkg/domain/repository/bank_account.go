package repository

import (
	"context"

	"github.com/datti-api/pkg/domain/model"
)

type BankAccountRepository interface {
	UpsertBankAccount(c context.Context, bank *model.BankAccount) (*model.BankAccount, error)
	GetBankAccountById(c context.Context, user *model.User) (*model.BankAccount, error)
}
