package models

type User struct {
	ID        string `db:"id"`
	Username  string `db:"username"`
	Password  string `db:"password"`
	Threshold int    `db:"threshold"`
}

type Session struct {
	ID     string `db:"id"`
	UserID string `db:"user_id"`
}

type Log struct {
	ID        string `db:"id"`
	UserID    string `db:"user_id"`
	DateStr   string `db:"date_str"`
	Elapsed   int  `db:"elapsed"`
	CreatedAt int64  `db:"created_at"`
}
