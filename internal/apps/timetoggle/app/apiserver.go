package app

import (
	"context"
	"os"

	"github.com/Oppodelldog/toggleperfect/internal/log"

	"github.com/Oppodelldog/toggleperfect/internal/apps/timetoggle/api/server"
	"github.com/Oppodelldog/toggleperfect/internal/apps/timetoggle/api/server/api"
	"github.com/go-openapi/loads"
	"github.com/jessevdk/go-flags"
)

func StartApiServer(ctx context.Context) {

	swaggerSpec, err := loads.Embedded(server.SwaggerJSON, server.FlatSwaggerJSON)
	if err != nil {
		log.Fatal(err)
	}

	swaggerApi := api.NewTimetoggleAPI(swaggerSpec)
	apiServer := server.NewServer(swaggerApi)

	swaggerApi.ProjectAddProjectHandler = AddProjectHandler{}
	swaggerApi.ProjectDeleteProjectHandler = DeleteProjectHandler{}
	swaggerApi.ProjectGetProjectByIDHandler = GetProjectHandler{}
	swaggerApi.ProjectGetProjectListHandler = GetProjectListHandler{}
	swaggerApi.CaptureAddStartCaptureHandler = AddCaptureStartHandler{}
	swaggerApi.CaptureAddStopCaptureHandler = AddCaptureStopHandler{}
	swaggerApi.CaptureSetLatestStopHandler = SetCaptureLatestStopHandler{}
	swaggerApi.CaptureGetCaptureListHandler = GetCaptureListHandler{}
	swaggerApi.ReportsGetReportCapturesTodayHandler = GetReportCapturesTodayHandler{}
	swaggerApi.ReportsGetReportCapturesCurrentMonthHandler = GetReportCapturesCurrentMonthHandler{}

	parser := flags.NewParser(apiServer, flags.Default)
	parser.ShortDescription = "Timetoggle API"
	parser.LongDescription = "Swagger definition for time toggle app in toggleperfect"
	apiServer.ConfigureFlags()

	for _, optsGroup := range swaggerApi.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatal(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	apiServer.ConfigureAPI()

	go func() {
		log.Print("Starting api server")
		apiServer.Port = 8001
		apiServer.Host = ""
		if err := apiServer.Serve(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		<-ctx.Done()
		log.Print("Shutting down api server")
		err := apiServer.Shutdown()
		if err != nil {
			log.Printf("error shuttiong down api server: %v", err)
		}
	}()
}
