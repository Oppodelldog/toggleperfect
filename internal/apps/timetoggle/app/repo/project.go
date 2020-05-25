package repo

import (
	"encoding/json"
	"fmt"
	"log"
)

type Project struct {
	Closed      bool
	Description string
	ID          string
}

func SaveProject(project Project) error {
	return saveProject(project)
}

func GetProject(ID string) (Project, error) {
	return openProject(ID)
}

func DeleteProject(ID string) error {
	return deleteProjectFile(ID)
}

func GetProjectList() ([]Project, error) {
	files, err := getStorageFiles(projectsDir, openProjectFileForReading)
	if err != nil {
		return nil, fmt.Errorf("error finding storage files: %v", err)
	}

	defer func() {
		for _, f := range files {
			err := f.Close()
			if err != nil {
				log.Printf("error in error cleanup while closing file: %v", err)
			}
		}
	}()

	var projects []Project
	for _, f := range files {
		var project Project
		err = json.NewDecoder(f).Decode(&project)
		if err != nil {
			return nil, fmt.Errorf("error decoding storage file '%s': %v", f.Name(), err)
		}

		projects = append(projects, project)
	}

	return projects, nil
}

func openProject(ID string) (Project, error) {
	f, err := openProjectFileForReading("", ID)
	if err != nil {
		return Project{}, err
	}
	defer closeWithPanic(f)

	var project Project
	err = json.NewDecoder(f).Decode(&project)

	return project, err
}

func saveProject(project Project) error {
	f, err := openProjectFileForReadingWriting("", project.ID)
	if err != nil {
		return err
	}
	defer closeWithPanic(f)

	return json.NewEncoder(f).Encode(project)
}
