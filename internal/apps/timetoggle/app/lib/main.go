package main

import (
	"github.com/Oppodelldog/toggleperfect/internal/apps"
	"github.com/Oppodelldog/toggleperfect/internal/apps/timetoggle/app"
	"github.com/Oppodelldog/toggleperfect/internal/display"
	"github.com/Oppodelldog/toggleperfect/internal/keys"
	"github.com/Oppodelldog/toggleperfect/internal/util"
	"log"
	"time"
)

func New(display display.UpdateChannel) apps.App {
	return &app.TimeToggle{Display: display}
}

func init() {

}

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

	projects := app.GetProjectsOverview()

	displayUpdate <- app.CreateStartScreen(projects)

	time.Sleep(time.Second)
}

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
