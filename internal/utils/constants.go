package utils

// grpc services

const (
	balance   = "BalanceActionService"
	transfer  = "TransferService"
	operation = "OperationService"
	user      = "UserService"
)

// grpc methods

const (
	BalanceFill = balance + " - Fill"
	BalanceGet  = balance + " - Get"

	TransferTransfer = transfer + " - Transfer"

	OperationOperationList = operation + " - OperationList"

	UserRegister = user + " - Register"
	UserLogin    = user + " - Login"
)
