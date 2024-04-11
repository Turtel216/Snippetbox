package models

import (
  "database/sql"
  "time"
  "errors"
)

type Snippet struct {
  ID      int
  Title   string
  Content string
  Created time.Time
  Expires time.Time
}

type SnippetModel struct {
  DB *sql.DB
}

func (model *SnippetModel) Insert(title string, content string, expires int) (int, error) {

  stmt :=  `INSERT INTO snippets (title, content, created, expires)
  VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

  result, err := model.DB.Exec(stmt, title, content, expires)
  if err != nil {
    return 0, err
  }

  id, err := result.LastInsertId()
  if err != nil {
    return 0, err 
  }

  return int(id), nil
}

func (model *SnippetModel) Get(id int) (*Snippet, error) {

  stmt := `SELECT id, title, content, created, expires FROM snippets
  WHERE expires > UTC_TIMESTAMP() AND id = ?`

  row := model.DB.QueryRow(stmt, id)

  snippet := &Snippet{}

  err := row.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)

  if err != nil {

    if errors.Is(err, sql.ErrNoRows) {
      return nil, ErrNoRecord
    } else {
      return nil, err
    }
  }

  return snippet, nil
}

func (model *SnippetModel) Latest() ([]*Snippet, error) {

  stmt := `SELECT id, title, conten, created, expires FROM snippets
  WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`
  
  rows, err := model.DB.Query(stmt)
  if err != nil {
    return nil, err
  }

  defer rows.Close()

  snippets := []*Snippet{}

  for rows.Next() {
    snippet := &Snippet{}

    err = rows.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)
    if err != nil {
      return nil, err
    }

    snippets = append(snippets, snippet)
  }

  if err = rows.Err(); err != nil {
    return nil, err
  }

  return snippets, nil


}
