package repo

import "time"

// DB models
type Person struct {
	ID       uint64    `db:"id" json:"id"`
	Username string    `db:"username" json:"username"`
	Password string    `db:"password" json:"-"`
	Balance  int32     `db:"balance" json:"balance"`
	Created  time.Time `db:"created" json:"created"`
}

type Transfer struct {
	ID        uint64    `db:"id"`
	Sender    Person    `db:"sen_id"`
	Receiver  Person    `db:"rec_id"`
	Amount    int32     `db:"amount"`
	Completed time.Time `db:"completed"`
}
