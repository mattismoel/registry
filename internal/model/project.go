package model

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type ProjectError error

type Project struct {
	ID    int64
	Title string
	Path  string

	Added   time.Time
	Updated time.Time
}

func (p Project) Validate() ProjectError {
	if p.Title == "" {
		return ProjectError(fmt.Errorf("no title provided"))
	}

	if p.Path == "" {
		return ProjectError(fmt.Errorf("no path provided"))
	}

	pathInfo, err := os.Stat(p.Path)
	if err != nil {
		return ProjectError(fmt.Errorf("could not stat path"))
	}

	if !pathInfo.IsDir() {
		return ProjectError(fmt.Errorf("project path is not directory"))
	}

	if !filepath.IsAbs(p.Path) {
		return ProjectError(fmt.Errorf("project path is not absolute"))
	}

	return nil
}
