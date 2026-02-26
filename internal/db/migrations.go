package db

import "github.com/joseph0x45/sad"

var migrations = []sad.Migration{
	{
		Version: 1,
		Name:    "create_recordings",
		SQL: `
      create table recording (
        id serial primary key,
        date_str text not null,
        seconds integer not null
      );
    `,
	},
}
