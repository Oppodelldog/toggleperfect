package app

import (
	"context"
	"github.com/Oppodelldog/toggleperfect/internal/apps/timetoggle/app/repo"
	"github.com/Oppodelldog/toggleperfect/internal/display"
	"github.com/Oppodelldog/toggleperfect/internal/keys"
)

type TimeToggle struct {
	Display       display.UpdateChannel
	serverCtx     context.Context
	cancelServer  func()
	activeProject int
	projects      []Project
	capturing     bool
}

func (a *TimeToggle) Init() {
	a.serverCtx, a.cancelServer = context.WithCancel(context.Background())
	StartApiServer(a.serverCtx)
	a.projects = loadProjects()
}

func (a TimeToggle) Dispose() {
	a.cancelServer()
}

func (a *TimeToggle) Activate() {
	a.projects = loadProjects()
	a.activeProject = 0
	a.Display <- CreateStartScreen(len(a.projects))
}

func (a TimeToggle) Deactivate() {
}

func (a *TimeToggle) HandleEvent(event keys.Event) bool {
	if !a.capturing {
		if event.State == keys.Clicked {
			if event.Key == keys.Key1 && a.hasProjects() {
				a.capturing = true
				a.currentProject().startCapture()
				a.Display <- CreateProjectScreen(a.currentProject())
				return true
			}
		}
	} else {
		if event.State == keys.Clicked {
			if event.Key == keys.Key1 {
				a.capturing = false
				a.currentProject().stopCapture()
				a.Display <- CreateStartScreen(len(a.projects))
			}
			if event.Key == keys.Key3 && a.hasProjects() {
				a.currentProject().stopCapture()
				nextProject := a.nextProject()
				nextProject.startCapture()
				a.Display <- CreateProjectScreen(nextProject)
			}
			if event.Key == keys.Key4 && a.hasProjects() {
				a.currentProject().stopCapture()
				previousProject := a.previousProject()
				previousProject.startCapture()
				a.Display <- CreateProjectScreen(previousProject)
			}
		}
		return true
	}
	return false
}

func (a *TimeToggle) nextProject() Project {
	a.activeProject++
	if a.activeProject == len(a.projects) {
		a.activeProject = 0
	}

	return a.currentProject()
}

func (a *TimeToggle) previousProject() Project {
	a.activeProject--
	if a.activeProject == -1 {
		a.activeProject = len(a.projects) - 1
	}

	return a.currentProject()
}

func (a TimeToggle) currentProject() Project {
	return a.projects[a.activeProject]
}

func (a TimeToggle) hasProjects() bool {
	return len(a.projects) > 0
}

func loadProjects() []Project {
	var projects []Project
	list, err := repo.GetProjectList()
	if err != nil {
		panic(err)
	}
	for _, prj := range list {
		projects = append(projects, Project{
			Name:        prj.ID,
			Description: prj.Description,
			Capture:     "",
		})
	}

	return projects
}
