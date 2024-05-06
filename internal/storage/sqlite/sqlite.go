package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/mattismoel/registry/internal/model"
	_ "github.com/mattn/go-sqlite3"
)

type sqliteStore struct {
	db *sql.DB
}

func New(dbPath string) (*sqliteStore, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("could not open database at path %q: %v", dbPath, err)
	}

	return &sqliteStore{db: db}, nil
}

func (s sqliteStore) Initialise() error {
	err := s.initialiseProjectsTable()
	if err != nil {
		return err
	}
	return nil
}

func (s sqliteStore) initialiseProjectsTable() error {
	query := `
  CREATE TABLE IF NOT EXISTS projects (
    id              INTEGER     NOT NULL PRIMARY KEY,
    title           TEXT        NOT NULL,
    path            TEXT        NOT NULL,
    date_added      DATETIME    DEFAULT CURRENT_TIMESTAMP,
    date_updated    DATETIME    DEFAULT CURRENT_TIMESTAMP
  )`
	_, err := s.db.Exec(query)
	if err != nil {
		return fmt.Errorf("could not create projects table: %v", err)
	}

	return nil
}

func (s sqliteStore) AddProject(proj model.Project) error {
	query := "INSERT INTO projects (title, path) VALUES (?, ?)"
	_, err := s.db.Exec(query, proj.Title, proj.Path)
	if err != nil {
		return fmt.Errorf("could not insert project: %v", err)
	}

	return nil
}

func (s sqliteStore) AllProjects() ([]model.Project, error) {
	query := `
  SELECT
    id,
    title,
    path,
    date_added,
    date_updated
  FROM
    projects
  ORDER BY
    date_added DESC`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("could not get projects from database: %v", err)
	}

	defer rows.Close()

	var projects []model.Project
	for rows.Next() {
		var p model.Project
		err := rows.Scan(&p.ID, &p.Title, &p.Path, &p.Added, &p.Updated)
		if err != nil {
			return nil, fmt.Errorf("could not scan into project struct: %v", err)
		}
		projects = append(projects, p)
	}

	return projects, nil
}

func (s sqliteStore) ProjectByID(id int64) (model.Project, error) {
	query := `
  SELECT
    id,
    title,
    path,
    date_added,
    date_updated
  FROM
    projects
  WHERE
    id = ?`

	var p model.Project
	err := s.db.QueryRow(query, id).Scan(
		&p.ID,
		&p.Title,
		&p.Path,
		&p.Added,
		&p.Updated,
	)
	if err != nil {
		return model.Project{}, fmt.Errorf("could not get project from database: %v", err)
	}

	return p, nil
}

func (s sqliteStore) ProjectByTitle(title string) (model.Project, error) {
	query := `
  SELECT
    id,
    title,
    path,
    date_added,
    date_updated
  FROM
    projects
  WHERE
    title = ?`

	var p model.Project
	err := s.db.QueryRow(query, title).Scan(
		&p.ID,
		&p.Title,
		&p.Path,
		&p.Added,
		&p.Updated,
	)
	if err != nil {
		return model.Project{}, fmt.Errorf("could not get project from database: %v", err)
	}

	return p, nil
}

func (s sqliteStore) RemoveProjectByID(int64) error {
	return nil
}
