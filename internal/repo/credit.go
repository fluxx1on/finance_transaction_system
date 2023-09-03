package repo

import (
	"context"
	"errors"
	"sync"

	"github.com/fluxx1on/finance_transaction_system/internal/mq/serial"
	"github.com/fluxx1on/finance_transaction_system/pkg/logger"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/exp/slog"
)

const (
	unavailable = "service unavailable"
)

type CreditDB struct {
	conn *pgxpool.Pool
}

func NewCreditDB(conn *pgxpool.Pool) *CreditDB {
	return &CreditDB{
		conn: conn,
	}
}

func (db *CreditDB) userByID(id uint64) (*Person, error) {
	var recipient Person

	query := "SELECT p.id, p.balance FROM Person AS p WHERE p.id = $1"
	err := db.conn.QueryRow(context.Background(), query, id).Scan(&recipient.ID, &recipient.Balance)
	if err != nil {
		return nil, errors.New("invalid username")
	}

	return &recipient, nil
}

func (db *CreditDB) userByName(username string) (*Person, error) {
	var recipient Person

	query := "SELECT p.id FROM Person AS p WHERE p.username = $1"
	err := db.conn.QueryRow(context.Background(), query, username).Scan(&recipient.ID)
	if err != nil {
		return nil, errors.New("invalid username")
	}

	return &recipient, nil
}

// PreTransfer gets transfer sides. Return senderID, recipientID, error(for responseMessage)
func (db *CreditDB) PreTransfer(senderID uint64, recipientUsername string, amountToTransfer int32) (
	uint64, uint64, error,
) {
	var (
		sender, recipient       *Person
		senderErr, recipientErr error
		wg                      = sync.WaitGroup{}
	)

	wg.Add(1)
	go func() {
		sender, senderErr = db.userByID(senderID)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		recipient, recipientErr = db.userByName(recipientUsername)
		wg.Done()
	}()

	wg.Wait()

	if senderErr != nil {
		slog.Debug("Error searching by id", senderErr)
		return 0, 0, senderErr
	}

	if recipientErr != nil {
		slog.Debug("Error searching by username", recipientErr)
		return 0, 0, recipientErr
	}

	// Checking that we have enough sum on senderBalance
	if sender.Balance < amountToTransfer {
		slog.Debug("Insufficient funds")
		return 0, 0, errors.New("insufficient funds")
	}

	return sender.ID, recipient.ID, nil
}

// TransferTransaction start up transaction between transfer sides
func (db *CreditDB) TransferTransaction(ctx context.Context, tn serial.TransactionInfo) error {

	tx, err := db.conn.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.RepeatableRead,
		AccessMode: pgx.ReadWrite,
	})
	if err != nil {
		slog.Error("Error starting transaction", err)
		return errors.New(unavailable)
	}

	// Updating sender and recipient Balances
	_, err = tx.Exec(ctx,
		"UPDATE Person SET balance = balance - $1 WHERE id = $2",
		tn.AmountToTransfer, tn.SenderID,
	)
	if err != nil {
		tx.Rollback(ctx)
		slog.Warn("Error updating sender balance:", err)
		return errors.New(unavailable)
	}

	_, err = tx.Exec(ctx,
		"UPDATE Person SET balance = balance + $1 WHERE id = $2",
		tn.AmountToTransfer, tn.RecipientID,
	)
	if err != nil {
		tx.Rollback(ctx)
		slog.Warn("Error updating recipient balance:", err)
		return errors.New(unavailable)
	}

	// Recording Transfer
	_, err = tx.Exec(ctx,
		"INSERT INTO Transfer (sen_id, rec_id, amount) VALUES ($1, $2, $3)",
		tn.SenderID, tn.RecipientID, tn.AmountToTransfer,
	)
	if err != nil {
		tx.Rollback(ctx)
		slog.Warn("Error recording transfer:", err)
		return errors.New(unavailable)
	}

	// Commit
	err = tx.Commit(ctx)
	if err != nil {
		slog.Warn("Error committing transaction:", err)
		return errors.New(unavailable)
	}

	slog.Info("Transfer successful", tn.SenderID, tn.RecipientID, tn.AmountToTransfer)
	return nil
}

func (db *CreditDB) CreateUser(username, password string) (*Person, error) {
	query := "INSERT INTO Person (username, password) VALUES ($1, $2) RETURNING id, created"
	var user Person
	err := db.conn.QueryRow(
		context.Background(), query, username, password, 0,
	).Scan(&user.ID, &user.Created)
	if err != nil {
		slog.Debug("Username isn't unique:", username)
		return nil, errors.New("user with these data already exist")
	}

	return &user, nil
}

func (db *CreditDB) GetUser(username string) (*Person, error) {
	query := `SELECT p.id, p.username, p.password, p.balance FROM Person AS p WHERE p.username = $1`
	row := db.conn.QueryRow(context.Background(), query, username)

	var person Person
	err := row.Scan(&person.ID, &person.Username, &person.Password, &person.Balance)
	if err != nil {
		return nil, errors.New("invalid username")
	}

	return &person, nil
}

func (db *CreditDB) CompletedTransferList(id uint64) ([]*Transfer, error) {
	query := `SELECT pr.username, tr.amount, tr.completed FROM Transfer AS tr
		JOIN Person AS pr ON pr.id = tr.rec_id 
		WHERE tr.sen_id = $1`

	rows, err := db.conn.Query(context.Background(), query, id)
	if err != nil {
		slog.Debug("Operations query cancelled")
		return nil, errors.New("invalid user_id")
	}
	defer rows.Close()

	var operations []*Transfer
	for rows.Next() {
		var operation Transfer
		err := rows.Scan(&operation.Receiver.Username, &operation.Amount, &operation.Completed)
		if err != nil {
			slog.Debug("Empty list of user operations. UserID:", id)
			return nil, err
		}
		operations = append(operations, &operation)
	}

	if rows.Err() != nil {
		return nil, errors.New(unavailable)
	}

	return operations, nil
}

func (db *CreditDB) FillBalance(id uint64, amount int32) error {

	ctx := context.TODO()

	tx, err := db.conn.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.RepeatableRead,
		AccessMode: pgx.ReadWrite,
	})
	if err != nil {
		slog.Error("Error starting transaction", err)
		return errors.New(unavailable)
	}

	// Updating User Balance
	_, err = tx.Exec(ctx,
		"UPDATE Person SET balance = balance + $1 WHERE id = $2",
		amount, id,
	)
	if err != nil {
		tx.Rollback(ctx)
		slog.Warn("Error recording transfer:", err)
		return errors.New(unavailable)
	}

	// Commit
	err = tx.Commit(ctx)
	if err != nil {
		slog.Warn("Error committing transaction:", err)
		return errors.New(unavailable)
	}

	slog.Info("Transfer successful", logger.Any(id, amount))
	return nil
}

func (db *CreditDB) GetBalance(id uint64) (int32, error) {
	query := `SELECT p.balance FROM Person AS p WHERE p.id = $1`
	row := db.conn.QueryRow(context.Background(), query, id)

	var balance int32
	err := row.Scan(&balance)
	if err != nil {
		return 0, errors.New("user doesn't exist")
	}

	return balance, nil
}
