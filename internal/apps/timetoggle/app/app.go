package app

import (
	"context"
	"fmt"
	"time"

	"github.com/Oppodelldog/toggleperfect/internal/log"

	"github.com/Oppodelldog/toggleperfect/internal/apps/timetoggle/app/repo"
	"github.com/Oppodelldog/toggleperfect/internal/display"
	"github.com/Oppodelldog/toggleperfect/internal/keys"
)

type TimeToggle struct {
	Display            display.UpdateChannel
	serverCtx          context.Context
	cancelServer       func()
	activeProject      int
	projects           []Project
	capturing          bool
	captureTicker      CaptureTicker
	updateScreenTicker *time.Ticker
	projectSummary     ProjectSummary
	summaryOffset      int
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
	a.updateScreenTicker = time.NewTicker(time.Minute * 10)
	a.projects = loadProjects()
	a.activeProject = 0
	a.summaryOffset = 0
	a.loadProjectSummary()
	a.Display <- CreateStartScreen(a.projectSummary)
}

func (a TimeToggle) Deactivate() {
	a.updateScreenTicker.Stop()
}

func (a *TimeToggle) HandleEvent(event keys.Event) bool {
	if !a.capturing {
		if event.State == keys.Clicked && a.hasProjects() {
			if event.Key == keys.Key1 {
				a.capturing = true
				a.startCapture(a.currentProject())
				a.Display <- CreateProjectScreen(a.currentProject())
				return true
			}
			if event.Key == keys.Key2 {
				a.summaryOffset--
				a.loadProjectSummary()
				a.Display <- CreateStartScreen(a.projectSummary)
				return true
			}
			if event.Key == keys.Key3 {
				a.projectSummary.Pagination.Page = a.projectSummary.Pagination.NextPage()
				a.Display <- CreateStartScreen(a.projectSummary)
				return true
			}
		}
		if event.State == keys.PressedReleased && a.hasProjects() {
			if event.Key == keys.Key2 {
				a.summaryOffset = 0
				a.loadProjectSummary()
				a.Display <- CreateStartScreen(a.projectSummary)
				return true
			}
			if event.Key == keys.Key3 {
				a.projectSummary.Pagination.Page = 1
				a.Display <- CreateStartScreen(a.projectSummary)
				return true
			}
		}
	} else {
		if event.State == keys.Clicked {
			if event.Key == keys.Key1 {
				a.capturing = false
				a.stopCapture()
				a.Display <- CreateStartScreen(a.projectSummary)
			}
			if event.Key == keys.Key2 && a.hasProjects() {
				projectID := a.currentProject().Name
				toggleProjectClosed(projectID)
				a.updateProjects()
			}
			if event.Key == keys.Key3 && a.hasProjects() {
				a.stopCapture()
				nextProject := a.nextProject()
				a.startCapture(nextProject)
				a.Display <- CreateProjectScreen(nextProject)
			}
			if event.Key == keys.Key4 && a.hasProjects() {
				a.stopCapture()
				previousProject := a.previousProject()
				a.startCapture(previousProject)
				a.Display <- CreateProjectScreen(previousProject)
			}
		}
		return true
	}
	return false
}

func (a *TimeToggle) loadProjectSummary() {
	a.projectSummary = GetProjectsOverview(a.summaryOffset)
	if len(a.projectSummary.Projects) == 0 {
		a.projectSummary.Projects = a.projects
	}
}

func (a *TimeToggle) updateProjects() {
	projectID := a.currentProject().Name
	a.projects = loadProjects()
	a.selectProject(projectID)
	a.Display <- CreateProjectScreen(a.currentProject())
}

func toggleProjectClosed(ID string) {
	p, err := repo.GetProject(ID)
	if err != nil {
		log.Printf("error loading project for toggle closed flag: %v", err)
	} else {
		p.Closed = !p.Closed
		err := repo.SaveProject(p)
		if err != nil {
			log.Printf("error saving project for toggle closed flag: %v", err)
		}
	}
}

func (a *TimeToggle) startCapture(previousProject Project) {
	previousProject.startCapture()
	a.captureTicker = NewCaptureTicker(a.serverCtx, a.currentProject().Name)
	a.captureTicker.Start()
}

func (a *TimeToggle) stopCapture() {
	a.currentProject().stopCapture()
	a.captureTicker.Stop()
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

func (a *TimeToggle) selectProject(ID string) {
	for i := 0; i < len(a.projects); i++ {
		if a.projects[i].Name == ID {
			a.activeProject = i
			return
		}
	}
	log.Printf("could not find project to select: %s", ID)
}

func loadProjects() []Project {
	var projects []Project
	list, err := repo.GetProjectList()
	if err != nil {
		err = fmt.Errorf("error loading projects: %v", err)
		panic(err)
	}
	for _, prj := range list {
		if prj.Closed {
			continue
		}

		projects = append(projects, Project{
			Name:        prj.ID,
			Description: prj.Description,
			Capture:     "",
			Closed:      prj.Closed,
		})
	}

	return projects
}

func GetProjectsOverview(offset int) ProjectSummary {
	date := time.Now().AddDate(0, 0, offset)
	captures, err := repo.GetReportCaptures(repo.TimeSpanDay(date))
	if err != nil {
		panic(err)
	}

	var projects []Project
	for _, project := range captures.Projects {
		if project.TimeWorked < 60 {
			continue
		}
		projects = append(projects, Project{
			Name:    project.ID,
			Capture: fmtDuration(time.Duration(project.TimeWorked) * time.Second),
		})
	}

	return ProjectSummary{
		Date:     date,
		Projects: projects,
		Pagination: Pagination{
			Page:     1,
			PerPage:  5,
			NumItems: len(projects),
		},
	}
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%02d:%02d", h, m)
}
