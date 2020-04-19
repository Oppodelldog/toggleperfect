package repo

import (
	"fmt"
	"github.com/Oppodelldog/toggleperfect/internal/util"
	"log"
	"os"
	"path"
)

var repoDir, projectsDir, capturesDir string

func init() {
	var hasRepoDir bool
	repoDir, hasRepoDir = os.LookupEnv("TOGGLE_PERFECT_REPO_DIR")

	if !hasRepoDir {
		repoDir = path.Dir(util.GetExecutableDir())
	}

	repoDir = path.Join(repoDir, ".data", "repo")
	log.Printf("repo directory is set to: %v", repoDir)
	err := os.MkdirAll(repoDir, 0777)
	if err != nil {
		panic(err)
	}

	projectsDir = path.Join(repoDir, "projects")
	log.Printf("repo projects directory is set to: %v", projectsDir)
	err = os.MkdirAll(projectsDir, 0777)
	if err != nil {
		panic(err)
	}

	capturesDir = path.Join(repoDir, "captures")
	log.Printf("repo captures directory is set to: %v", capturesDir)
	err = os.MkdirAll(capturesDir, 0777)
	if err != nil {
		panic(err)
	}
}

func openProjectFileForWriting(ID string) (*os.File, error) {
	return os.OpenFile(getProjectFilePath(ID), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0655)
}

func openProjectFileForReading(ID string) (*os.File, error) {
	return os.OpenFile(getProjectFilePath(ID), os.O_RDONLY, 0655)
}

func deleteProjectFile(ID string) error {
	return os.Remove(getProjectFilePath(ID))
}

func getProjectFilePath(ID string) string {
	projectFileName := fmt.Sprintf("%v.json", ID)
	projectFilePath := path.Join(projectsDir, projectFileName)
	return projectFilePath
}

func openCaptureFileForWriting(ID string) (*os.File, error) {
	return os.OpenFile(getCaptureFilepath(ID), os.O_CREATE|os.O_RDWR, 0655)
}

func getCaptureFilepath(ID string) string {
	captureFilename := fmt.Sprintf("%v.json", ID)
	captureFilePath := path.Join(capturesDir, captureFilename)
	return captureFilePath
}
func closeWithPanic(f *os.File) {
	err := f.Close()
	if err != nil {
		panic(err)
	}
}
