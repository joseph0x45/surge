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
        password text not null
      );
    `,
	},
}
