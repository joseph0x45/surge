package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/joseph0x45/surge/internal/models"
)

func (c *Conn) getUserByUsername(username string) (*models.User, error) {
	const query = "select * from users where username=?"
	user := &models.User{}
	err := c.db.Get(user, query, username)
	if err == nil {
		return user, nil
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return nil, fmt.Errorf("Error while getting user by username: %w", err)
}

func (c *Conn) getUserByID(id string) (*models.User, error) {
	const query = "select * from users where id=?"
	user := &models.User{}
	err := c.db.Get(user, query, id)
	if err == nil {
		return user, nil
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return nil, fmt.Errorf("Error while getting user by id: %w", err)
}

func (c *Conn) GetUser(by, value string) (*models.User, error) {
	var user *models.User
	var err error
	if by == "username" {
		user, err = c.getUserByUsername(value)
	} else {
		user, err = c.getUserByID(value)
	}
	return user, err
}

func (c *Conn) InsertUser(user *models.User) error {
	const query = `
    insert into users (
      id, username, password
    )
    values (
      :id, :username, :password
    );
  `
	if _, err := c.db.NamedExec(query, user); err != nil {
		return fmt.Errorf("Error while inserting user: %w", err)
	}
	return nil
}

func (c *Conn) UsernameExists(username string) (bool, error) {
	const query = "select exists(select 1 from users where username=?)"
	var exists bool
	err := c.db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("Error while checking for username existence: %w", err)
	}
	return exists, nil
}
