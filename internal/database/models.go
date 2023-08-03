package database

import "time"

// DB models
type Person struct {
	ID       int       `db:"id"`
	Username string    `db:"username"`
	Password string    `db:"password"`
	Balance  int       `db:"balance"`
	Created  time.Time `db:"created"`
}

type Token struct {
	UserID int    `db:"user_id"`
	Token  string `db:"token"`
}

type Transfer struct {
	ID         int       `db:"id"`
	SenderID   int       `db:"sen_id"`
	ReceiverID int       `db:"rec_id"`
	Amount     int       `db:"amount"`
	Completed  time.Time `db:"completed"`
}
