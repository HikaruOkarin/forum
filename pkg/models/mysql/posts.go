package mysql

import (
	"database/sql"
	"forum/pkg/models"
)

type PostModel struct {
	DB *sql.DB
}

func (m *PostModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title,content,created,expires)
	VALUES(?,?,UTC_TIMESTAMP(),DATE_ADD(UTC_TIMESTAMP(),INTERVAL ? DAY))`
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *PostModel) Get(id int) (*models.Post, error) {
	return nil, nil
}

func (m *PostModel) Latest() ([]*models.Post, error) {
	return nil, nil
}
