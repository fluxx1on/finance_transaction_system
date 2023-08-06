package database

import (
	"context"
	"fmt"

	"github.com/fluxx1on/finance_transaction_system/internal/rpc"
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

// client
func (db *CreditDB) userByToken(token string) (*Person, error) {
	var sender Person

	query := "SELECT p.id, p.balance FROM Token AS t JOIN Person AS p ON p.id = t.user_id WHERE t.token = $1"
	err := db.conn.QueryRow(context.Background(), query, token).Scan(&sender.ID, &sender.Balance)
	if err != nil {
		return nil, err
	}
	// ID, Balance returned
	return &sender, nil
}

// recipient
func (db *CreditDB) userByName(username string) (*Person, error) {
	var recipient Person

	query := "SELECT p.id FROM Person AS p WHERE p.username = $1"
	err := db.conn.QueryRow(context.Background(), query, username).Scan(&recipient.ID)
	if err != nil {
		return nil, err
	}
	// Only ID returned
	return &recipient, nil
}

// PreTransfer get transfer sides. Return senderID, recipientID, error(for responseMessage)
func (db *CreditDB) PreTransfer(token, recipientUsername string, amountToTransfer int) (int, int, error) {
	var sender, recipient *Person

	sender, err := db.userByToken(token)
	if err != nil {
		slog.Debug("Error scanning token", err)
		return 0, 0, fmt.Errorf("wrong token")
	}

	recipient, err = db.userByName(recipientUsername)
	if err != nil {
		slog.Debug("Error searching by username", err)
		return 0, 0, fmt.Errorf("no one user with name: %s", recipientUsername)
	}

	// Checking that we have enough sum on senderBalance
	if sender.Balance < amountToTransfer {
		slog.Debug("Insufficient funds")
		return 0, 0, fmt.Errorf("insufficient funds")
	}

	return sender.ID, recipient.ID, nil
}

// TransferTransaction start up transaction between transfer sides
func (db *CreditDB) TransferTransaction(ctx context.Context, tn rpc.TransactionInfo) error {

	tx, err := db.conn.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.RepeatableRead,
		AccessMode: pgx.ReadWrite,
	})
	if err != nil {
		slog.Error("Error starting transaction", err)
		return fmt.Errorf(unavailable)
	}

	// Updating sender and recipient Balances
	_, err = tx.Exec(context.Background(),
		"UPDATE Person SET balance = balance - $1 WHERE id = $2",
		tn.AmountToTransfer, tn.SenderID,
	)
	if err != nil {
		tx.Rollback(ctx)
		slog.Warn("Error updating sender balance:", err)
		return fmt.Errorf(unavailable)
	}

	_, err = tx.Exec(context.Background(),
		"UPDATE Person SET balance = balance + $1 WHERE id = $2",
		tn.AmountToTransfer, tn.RecipientID,
	)
	if err != nil {
		tx.Rollback(ctx)
		slog.Warn("Error updating recipient balance:", err)
		return fmt.Errorf(unavailable)
	}

	// Recording Transfer
	_, err = tx.Exec(context.Background(),
		"INSERT INTO Transfer (sen_id, rec_id, amount) VALUES ($1, $2, $3)",
		tn.SenderID, tn.RecipientID, tn.AmountToTransfer,
	)
	if err != nil {
		tx.Rollback(ctx)
		slog.Warn("Error recording transfer:", err)
		return fmt.Errorf(unavailable)
	}

	// Commit
	err = tx.Commit(context.Background())
	if err != nil {
		slog.Warn("Error committing transaction:", err)
		return fmt.Errorf(unavailable)
	}

	slog.Info("Transfer successful", tn.SenderID, tn.RecipientID, tn.AmountToTransfer)
	return nil
}

func (db *CreditDB) CreateUser(username, password string) (*Person, error) {

	// TODO

	return nil, nil
}

func (db *CreditDB) UserByAuth(username string) (*Person, error) {
	query := "SELECT p.id, p.username, p.password, p.balance FROM Person AS p WHERE p.username = $1"
	row := db.conn.QueryRow(context.Background(), query, username)

	var person Person
	err := row.Scan(&person.ID, &person.Username, &person.Password, &person.Balance)
	if err != nil {
		return nil, err
	}

	return &person, nil
}

// Token find a token by Person.ID. Called After GetUserByAuth to Login
func (db *CreditDB) Token(id int) (*Token, error) {
	query := "SELECT t.token FROM Token AS t WHERE t.user_id = $1"
	row := db.conn.QueryRow(context.Background(), query, id)

	var token Token
	err := row.Scan(&token.Token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (db *CreditDB) CompletedTransferList(token string) ([]*Transfer, error) {
	query := "SELECT tr.rec_id, tr.amount, tr.completed FROM Transfer AS tr JOIN Token AS t ON t.user_id = tr.recipient WHERE t.token = $1"
	rows, err := db.conn.Query(context.Background(), query, token)
	if err != nil {
		slog.Debug("Operations query cancelled")
		return nil, err
	}

	var operations []*Transfer
	for rows.Next() {
		var operation Transfer
		err := rows.Scan(&operation.ReceiverID, &operation)
		if err != nil {
			return nil, err
		}
		operations = append(operations, &operation)
	}

	return operations, nil
}

func (db *CreditDB) Close() {

}
