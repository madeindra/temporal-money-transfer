package app

import (
	"context"
	"fmt"
	"log"
)

func Withdraw(ctx context.Context, data PaymentDetails) (string, error) {
	log.Printf("Withdrawing $%d from account %s.\n\n",
		data.Amount,
		data.SourceAccount,
	)

	// connect to banking service
	bank := BankingService{"bank-api.example.com"}

	// run withdraw method from banking service
	referenceID := fmt.Sprintf("%s-withdrawal", data.ReferenceID)
	confirmation, err := bank.Withdraw(data.SourceAccount, data.Amount, referenceID)

	return confirmation, err
}

func Deposit(ctx context.Context, data PaymentDetails) (string, error) {
	log.Printf("Depositing $%d into account %s.\n\n",
		data.Amount,
		data.TargetAccount,
	)

	// connect to banking service
	bank := BankingService{"bank-api.example.com"}

	// run deposit method from banking service
	referenceID := fmt.Sprintf("%s-deposit", data.ReferenceID)

	// this method will fail
	// confirmation, err := bank.DepositThatFails(data.TargetAccount, data.Amount, referenceID)

	// this method will succeed
	confirmation, err := bank.Deposit(data.TargetAccount, data.Amount, referenceID)

	return confirmation, err
}

func Refund(ctx context.Context, data PaymentDetails) (string, error) {
	log.Printf("Refunding $%v back into account %v.\n\n",
		data.Amount,
		data.SourceAccount,
	)

	// connect to banking service
	bank := BankingService{"bank-api.example.com"}

	// run refund method from banking service
	referenceID := fmt.Sprintf("%s-refund", data.ReferenceID)
	confirmation, err := bank.Deposit(data.SourceAccount, data.Amount, referenceID)

	return confirmation, err
}
