package repo

import "time"

// DB models
type Person struct {
	ID       int       `db:"id" json:"id"`
	Username string    `db:"username" json:"username"`
	Password string    `db:"password" json:"-"`
	Balance  int       `db:"balance" json:"balance"`
	Created  time.Time `db:"created" json:"created"`
}

type Transfer struct {
	ID         int       `db:"id"`
	SenderID   int       `db:"sen_id"`
	ReceiverID int       `db:"rec_id"`
	Amount     int       `db:"amount"`
	Completed  time.Time `db:"completed"`
}
