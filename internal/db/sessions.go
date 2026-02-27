package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/joseph0x45/surge/internal/models"
)

func (c *Conn) InsertSession(session *models.Session) error {
	const query = `
    insert into sessions (
      id, user_id
    )
    values (
      :id, :user_id
    );
  `
	if _, err := c.db.NamedExec(query, session); err != nil {
		return fmt.Errorf("Error while inserting session: %w", err)
	}
	return nil
}

func (c *Conn) GetSession(id string) (*models.Session, error) {
	const query = "select * from sessions where id=?"
	session := &models.Session{}
	err := c.db.Get(session, query, id)
	if err == nil {
		return session, nil
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return nil, fmt.Errorf("Error while getting session: %w", err)
}

func (c *Conn) DeleteSession(id string) error {
	const query = "delete from sessions where id=?"
	if _, err := c.db.Exec(query, id); err != nil {
		return fmt.Errorf("Error while deleting sessio: %w", err)
	}
	return nil
}
