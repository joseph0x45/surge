package db

import (
	"fmt"

	"github.com/joseph0x45/surge/internal/models"
)

func (c *Conn) UpdateLogs(log *models.Log) error {
	const query = `
    insert into logs (
      id, user_id, date_str, elapsed, created_at
    )values (
      :id, :user_id, :date_str, :elapsed, :created_at
    )
    on conflict(user_id, date_str) do update set elapsed = excluded.elapsed;
  `
	if _, err := c.db.NamedExec(query, log); err != nil {
		return fmt.Errorf("Error while inserting log: %w", err)
	}
	return nil
}

func (c *Conn) GetUserLogs(userID string) ([]models.Log, error) {
	logs := []models.Log{}
	const query = `
    select * from logs
    where user_id = ?
    order by created_at desc;
  `
	if err := c.db.Select(&logs, query, userID); err != nil {
		return nil, fmt.Errorf("Error while getting user logs: %w", err)
	}
	return logs, nil
}
