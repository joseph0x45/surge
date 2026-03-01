package db

import "github.com/joseph0x45/sad"

var migrations = []sad.Migration{
	{
		Version: 1,
		Name:    "create_users",
		SQL: `
      create table users (
        id text not null primary key,
        username text not null unique,
        password text not null,
        time_limit integer not null
      );
    `,
	},
	{
		Version: 2,
		Name:    "create_sessions",
		SQL: `
      create table sessions (
        id text not null primary key,
        user_id text not null references users(id)
      );
    `,
	},
	{
		Version: 3,
		Name:    "create_logs",
		SQL: `
      create table logs (
        id text not null primary key,
        user_id text not null references users(id),
        date_str text not null,
        elapsed integer not null,
        created_at integer not null,
        unique(user_id, date_str)
      );
    `,
	},
}
