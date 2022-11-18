package app

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func MoneyTransfer(ctx workflow.Context, input PaymentDetails) (string, error) {

	// create RetryPolicy for when activity fail
	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:        time.Second,
		BackoffCoefficient:     2.0,
		MaximumInterval:        100 * time.Second,
		MaximumAttempts:        0, // 0 = unlimited retries
		NonRetryableErrorTypes: []string{"InvalidAccountError", "InsufficientFundsError"},
	}

	// create ActivityOptions with Timeout & RetryPolicy
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute, // timeout in 1 minute
		RetryPolicy:         retrypolicy,
	}

	// apply ActivityOptions
	ctx = workflow.WithActivityOptions(ctx, options)

	// create result variable
	var withdrawOutput string

	// run withdraw activity
	withdrawErr := workflow.ExecuteActivity(ctx, Withdraw, input).Get(ctx, &withdrawOutput)
	if withdrawErr != nil {
		return "", withdrawErr
	}

	// create deposit result variable
	var depositOutput string

	// run deposit activity
	depositErr := workflow.ExecuteActivity(ctx, Deposit, input).Get(ctx, &depositOutput)
	if depositErr != nil {
		// create refund result variable
		var result string

		// if deposit failed, run refund activity
		refundErr := workflow.ExecuteActivity(ctx, Refund, input).Get(ctx, &result)
		if refundErr != nil {
			return "",
				fmt.Errorf("Deposit: failed to deposit money into %v: %v. Money could not be returned to %v: %w",
					input.TargetAccount, depositErr, input.SourceAccount, refundErr)
		}

		return "", fmt.Errorf("Deposit: failed to deposit money into %v: Money returned to %v: %w",
			input.TargetAccount, input.SourceAccount, depositErr)
	}

	result := fmt.Sprintf("Transfer complete (transaction IDs: %s, %s)", withdrawOutput, depositOutput)
	return result, nil
}
