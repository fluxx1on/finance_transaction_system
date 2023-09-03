package serial

// RabbitMQ serializer
type TransactionInfo struct {
	SenderID         uint64 `json:"sender_id"`
	RecipientID      uint64 `json:"recipient_id"`
	AmountToTransfer int32  `json:"amountToTransfer"`
}
