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
	smt := `SELECT id, title,content, created,expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`
	row := m.DB.QueryRow(smt, id)

	s := &models.Post{}
	err := row.Scan(&s.Id, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return s, nil
}

func (m *PostModel) Latest() ([]*models.Post, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	snippets := []*models.Post{}
	for rows.Next() {
		s := &models.Post{}
		err = rows.Scan(&s.Id, &s.Content, &s.Title, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		} else {
			snippets = append(snippets, s)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return snippets, nil
}
