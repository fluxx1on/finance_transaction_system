package serial

// RabbitMQ serializer
type TransactionInfo struct {
	SenderID         int `json:"sender_id"`
	RecipientID      int `json:"recipient_id"`
	AmountToTransfer int `json:"amountToTransfer"`
}
