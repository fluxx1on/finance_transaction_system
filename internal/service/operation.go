package service

import (
	"context"

	"github.com/fluxx1on/finance_transaction_system/internal/auth"
	"github.com/fluxx1on/finance_transaction_system/internal/repo"
	"github.com/fluxx1on/finance_transaction_system/internal/transport/rpc/pb/transfer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OperationFetcher struct {
	db *repo.CreditDB
}

func NewOperationFetcher(db *repo.CreditDB) *OperationFetcher {

	return &OperationFetcher{
		db: db,
	}
}

func (s *OperationFetcher) FetchOperationList(ctx context.Context, req *transfer.EmptyRequest) (
	*transfer.OperationResponseList, error,
) {
	var (
		sttMsg bool = true
		errMsg string
	)

	user, err := auth.GetUserFromContext(ctx)
	if err != nil {
		sttMsg = false
		errMsg = err.Error()
	}

	var operations []*transfer.TransferInfo
	if sttMsg {
		transfers, err := s.db.CompletedTransferList(user.ID)
		if err != nil {

			sttMsg = false
			errMsg = status.Error(codes.Internal, "Unexpected error").Error()

		} else {

			operations = make([]*transfer.TransferInfo, 0, len(transfers))
			for key := range transfers {
				t := transfers[key]
				operations = append(operations, &transfer.TransferInfo{
					TransferSum:   int32(t.Amount),
					RecipientName: t.Receiver.Username,
					TimeCompleted: timestamppb.New(t.Completed),
				})
			}

		}
	}

	resp := &transfer.OperationResponseList{
		Status:       sttMsg,
		ErrorMessage: errMsg,
		Operations:   operations,
	}

	return resp, err
}
