package repo

type Project struct {
	Description string
	ID          string
}

func AddProject(project Project) error {
	var err error

	return err
}
func UpdateProject(project Project) error {
	var err error

	return err
}
func GetProject(ID string) (Project, error) {
	var err error

	return Project{
		Description: "My Project Nr " + ID,
		ID:          ID,
	}, err
}
func DeleteProject(ID string) error {
	var err error

	return err
}
