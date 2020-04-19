package main

import (
	"gitlab.com/Oppodelldog/toggleperfect/internal/apps"
	"gitlab.com/Oppodelldog/toggleperfect/internal/apps/timetoggle/app"
	"gitlab.com/Oppodelldog/toggleperfect/internal/display"
)

func New(display display.UpdateChannel) apps.App {
	return &app.App{Display: display}
}
