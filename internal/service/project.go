package service

import (
	"fmt"

	"github.com/mattismoel/registry/internal/model"
	"github.com/mattismoel/registry/internal/storage"
)

type ProjectService struct {
	store storage.Store
}

func NewProjectService(store storage.Store) *ProjectService {
	return &ProjectService{store: store}
}

func (ps ProjectService) All() ([]model.Project, error) {
	projects, err := ps.store.AllProjects()
	if err != nil {
		return nil, fmt.Errorf("could not get projects: %v", err)
	}

	return projects, nil
}

func (ps ProjectService) Add(project model.Project) error {
	err := project.Validate()
	if err != nil {
		return err
	}

	err = ps.store.AddProject(project)
	if err != nil {
		return fmt.Errorf("could not add project: %v", err)
	}

	return nil
}

func (ps ProjectService) ByID(id int64) (model.Project, error) {
	project, err := ps.store.ProjectByID(id)
	if err != nil {
		return model.Project{}, fmt.Errorf("could not get project by id %d: %v", id, err)
	}

	return project, nil
}

func (ps ProjectService) ByTitle(title string) (model.Project, error) {
	project, err := ps.store.ProjectByTitle(title)
	if err != nil {
		return model.Project{}, fmt.Errorf("could not get project by title %q: %v", title, err)
	}

	return project, nil
}
