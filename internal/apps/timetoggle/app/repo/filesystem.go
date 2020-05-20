package repo

import (
	"fmt"
	"github.com/Oppodelldog/toggleperfect/internal/util"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

var repoDir, projectsDir, capturesDir string

const dirPerm = 0777
const filePerm = 0655

func init() {
	var hasRepoDir bool
	repoDir, hasRepoDir = os.LookupEnv("TOGGLE_PERFECT_REPO_DIR")

	if !hasRepoDir {
		repoDir = path.Dir(util.GetExecutableDir())
	}

	repoDir = path.Join(repoDir, ".data", "repo")
	log.Printf("repo directory is set to: %v", repoDir)

	err := os.MkdirAll(repoDir, dirPerm)
	if err != nil {
		panic(err)
	}

	projectsDir = path.Join(repoDir, "projects")
	log.Printf("repo projects directory is set to: %v", projectsDir)
	err = os.MkdirAll(projectsDir, dirPerm)
	if err != nil {
		panic(err)
	}

	capturesDir = path.Join(repoDir, "captures")
	log.Printf("repo captures directory is set to: %v", capturesDir)
	err = os.MkdirAll(capturesDir, dirPerm)
	if err != nil {
		panic(err)
	}
}

func openProjectFileForReadingWriting(fileDir string, ID string) (*os.File, error) {
	return os.OpenFile(getProjectFilePath(ID), os.O_CREATE|os.O_TRUNC|os.O_RDWR, filePerm)
}

func openProjectFileForReading(fileDir string, ID string) (*os.File, error) {
	return os.OpenFile(getProjectFilePath(ID), os.O_RDONLY, filePerm)
}

func deleteProjectFile(ID string) error {
	return os.Remove(getProjectFilePath(ID))
}

func getProjectFilePath(ID string) string {
	projectFileName := fmt.Sprintf("%v.json", ID)
	projectFilePath := path.Join(projectsDir, projectFileName)
	return projectFilePath
}

func openCaptureFileForWritingForCurrentMonth(fileDir string, ID string) (*os.File, error) {
	storageFilePath := getCaptureFilepathForCurrentMonth(ID)
	err := os.MkdirAll(path.Dir(storageFilePath), dirPerm)
	if err != nil {
		panic(err)
	}
	return os.OpenFile(storageFilePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, filePerm)
}

func openCaptureFileForReadingForCurrentMonth(fileDir string, ID string) (*os.File, error) {
	return os.OpenFile(path.Join(fileDir, getCaptureFilepathForCurrentMonth(ID)), os.O_RDONLY, filePerm)
}

func openCaptureFileForReadingFromDir(fileDir string, ID string) (*os.File, error) {
	return os.OpenFile(path.Join(fileDir, getCaptureFileName(ID)), os.O_RDONLY, filePerm)
}

func getCaptureFilepathForCurrentMonth(ID string) string {
	captureTimeDir := time.Now().Format("200601")
	captureFilename := getCaptureFileName(ID)
	captureFilePath := path.Join(capturesDir, captureTimeDir, captureFilename)
	return captureFilePath
}

func getCaptureFileName(ID string) string {
	captureFilename := fmt.Sprintf("%v.json", ID)
	return captureFilename
}

func closeWithPanic(f *os.File) {
	err := f.Close()
	if err != nil {
		panic(err)
	}
}

type openStorageFile func(fileDir string, ID string) (*os.File, error)

func getStorageFiles(walkDir string, openFileFunc openStorageFile) ([]*os.File, error) {
	files := []*os.File{}
	err := filepath.Walk(walkDir, func(filePath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		suffix := ".json"
		if strings.HasSuffix(info.Name(), suffix) {
			id := info.Name()[:len(info.Name())-len(suffix)]
			f, err := openFileFunc(path.Dir(filePath), id)
			if err != nil {
				return err
			}
			files = append(files, f)
		}

		return nil
	})

	if err != nil {
		for _, f := range files {
			err := f.Close()
			if err != nil {
				log.Printf("error in error cleanup while closing file: %v", err)
			}
		}
		return nil, err
	}

	return files, nil
}
