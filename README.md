# Money Transfer with Temporal

This project mimicks bank service that allow user to deposit and withdraw money.

## Start Temporal

1. Clone temporal docker repo
```
git clone https://github.com/temporalio/docker-compose.git
```

2. Start Temporal
```
cd docker-compose
docker-compose up
```

## Get the Go SDK

1. Get the SDK
```
go get go.temporal.io/sdk
```

2. or clone the Go SDK repo
```
git clone https://github.com/temporalio/sdk-go.git
```

## Run Application

1. Start the worker
```
go run worker/main.go
```

2. Start the initiator
```
go run start/main.go
```

## Error on Activity

Try toggling between `bank.Deposit` and `bank.DepositThatFails` in [activity.go](./activity.go) to see the retry in action.

Remember to restart the worker after changing the code.

## References

This project is created based on tutorial on [Temporal](https://learn.temporal.io/getting_started/go/first_program_in_go/).