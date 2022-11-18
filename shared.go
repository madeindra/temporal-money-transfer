package app

type PaymentDetails struct {
	SourceAccount string
	TargetAccount string
	Amount        int
	ReferenceID   string
}

const MoneyTransferTaskQueueName = "MONEY_TRANSFER_TASK_QUEUE"
