package storage

import "github.com/mattismoel/registry/internal/model"

type Store interface {
	AddProject(model.Project) error
	AllProjects() ([]model.Project, error)
	ProjectByID(int64) (model.Project, error)
	ProjectByTitle(string) (model.Project, error)
	RemoveProjectByID(int64) error
}
