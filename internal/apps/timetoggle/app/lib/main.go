package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Oppodelldog/toggleperfect/internal/apps"
	"github.com/Oppodelldog/toggleperfect/internal/apps/timetoggle/app"
	"github.com/Oppodelldog/toggleperfect/internal/display"
	"github.com/Oppodelldog/toggleperfect/internal/keys"
	"github.com/Oppodelldog/toggleperfect/internal/util"
)

func New(display display.UpdateChannel) apps.App {
	return &app.TimeToggle{Display: display}
}

func init() {

}

//noinspection GoUnusedFunction
func main2() {
	ctx := util.NewInterruptContext()
	displayUpdate := apps.NewDevDisplayChannel(ctx)

	displayUpdate <- app.CreateProjectScreen(app.Project{
		Name:        "JIRA-ISSUE-19",
		Description: "ADD A PAGE",
		Capture:     "",
	})

	time.Sleep(time.Second)
}

func main() {
	ctx := util.NewInterruptContext()
	displayUpdate := apps.NewDevDisplayChannel(ctx)

	/*
		projects := []app.Project{
			{Name: "1", Capture: "A"},
			{Name: "2", Capture: "B"},
			{Name: "3", Capture: "C"},
			{Name: "4", Capture: "D"},
			{Name: "5", Capture: "E"},
			{Name: "6", Capture: "F"},
			{Name: "7", Capture: "G"},
			{Name: "8", Capture: "H"},
			{Name: "9", Capture: "I"},
			{Name: "10", Capture: "J"},
		}

		projectSummary := app.ProjectSummary{
			Date:       time.Now(),
			Projects:   projects,
			Pagination: app.Pagination{Page: 1, NumItems: len(projects), PerPage: 5},
		}
	*/

	projectSummary := app.GetProjectsOverview(-2)
	projectSummary.Projects = append(projectSummary.Projects, app.Project{
		Name:        "TEST",
		Description: "",
		Capture:     "sfmnik",
	})
	projectSummary.Pagination.NumItems = len(projectSummary.Projects)
	fmt.Print(projectSummary.Pagination.NextPage())
	displayUpdate <- app.CreateStartScreen(projectSummary)

	time.Sleep(time.Second)
}

//noinspection GoUnusedFunction
func main1() {
	log.Print("** TimeToggle Standalone **")
	ctx := util.NewInterruptContext()

	displayUpdate := apps.NewDevDisplayChannel(ctx)
	timeToggle := New(displayUpdate)
	timeToggle.Init()
	timeToggle.Activate()

	time.AfterFunc(time.Second, func() {
		timeToggle.HandleEvent(keys.Event{
			State: keys.Clicked,
			Key:   keys.Key1,
		})
	})

	time.AfterFunc(time.Second*10, func() {
		timeToggle.HandleEvent(keys.Event{
			State: keys.Clicked,
			Key:   keys.Key3,
		})
	})

	<-ctx.Done()
	timeToggle.Deactivate()
	timeToggle.Dispose()
}
