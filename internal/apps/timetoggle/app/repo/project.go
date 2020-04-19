package repo

import (
	"encoding/json"
)

type Project struct {
	Description string
	ID          string
}

func AddProject(project Project) error {
	f, err := openProjectFileForWriting(project.ID)
	if err != nil {
		return err
	}
	defer closeWithPanic(f)

	return json.NewEncoder(f).Encode(project)
}

func GetProject(ID string) (Project, error) {
	f, err := openProjectFileForReading(ID)
	if err != nil {
		return Project{}, err
	}
	defer closeWithPanic(f)

	var project Project
	err = json.NewDecoder(f).Decode(&project)

	return project, err
}

func DeleteProject(ID string) error {
	return deleteProjectFile(ID)
}
