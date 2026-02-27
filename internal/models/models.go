package models

import "html/template"

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

type PageData struct {
	Title   string
	Content template.HTML
}
