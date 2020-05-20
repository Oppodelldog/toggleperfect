package app

import (
	"github.com/Oppodelldog/toggleperfect/internal/apps/timetoggle/api/model"
	"github.com/Oppodelldog/toggleperfect/internal/apps/timetoggle/api/server/api/capture"
	"github.com/Oppodelldog/toggleperfect/internal/apps/timetoggle/api/server/api/project"
	"github.com/Oppodelldog/toggleperfect/internal/apps/timetoggle/api/server/api/reports"
	"github.com/Oppodelldog/toggleperfect/internal/apps/timetoggle/app/repo"
	"github.com/go-openapi/runtime/middleware"
	"time"
)

type GetProjectListHandler struct{}

func (g GetProjectListHandler) Handle(params project.GetProjectListParams) middleware.Responder {
	payloadProjects := []*model.Project{}

	projects, err := repo.GetProjectList()
	if err != nil {
		return &project.GetProjectListInternalServerError{Payload: &model.ServerError{
			Description: err.Error(),
		}}
	}
	for _, prj := range projects {
		payloadProjects = append(payloadProjects, projectToPayload(prj))
	}

	return &project.GetProjectListOK{Payload: &model.Projects{Projects: payloadProjects}}
}

type AddProjectHandler struct {
}

func (a AddProjectHandler) Handle(params project.AddProjectParams) middleware.Responder {
	err := repo.AddProject(projectFromPayload(params.Body))
	if err != nil {
		return &project.AddProjectInternalServerError{Payload: &model.ServerError{Description: err.Error()}}
	}
	return &project.AddProjectNoContent{}
}

type DeleteProjectHandler struct {
}

func (d DeleteProjectHandler) Handle(params project.DeleteProjectParams) middleware.Responder {
	err := repo.DeleteProject(params.ProjectID)
	if err != nil {
		return &project.DeleteProjectNotFound{}
	}
	return &project.DeleteProjectNoContent{}
}

type GetProjectHandler struct {
}

func (g GetProjectHandler) Handle(params project.GetProjectByIDParams) middleware.Responder {
	prj, err := repo.GetProject(params.ProjectID)
	if err != nil {
		return &project.GetProjectByIDNotFound{}
	}

	return &project.GetProjectByIDOK{Payload: projectToPayload(prj)}
}

func projectFromPayload(payload *model.Project) repo.Project {
	return repo.Project(*payload)
}

func projectToPayload(prj repo.Project) *model.Project {
	p := model.Project(prj)
	return &p
}

type GetCaptureListHandler struct{}

func (g GetCaptureListHandler) Handle(params capture.GetCaptureListParams) middleware.Responder {
	payload := &model.Captures{}

	captures, err := repo.GetCaptures()
	if err != nil {
		return &capture.GetCaptureListInternalServerError{Payload: &model.ServerError{
			Description: err.Error(),
		}}
	}
	for _, captureEntry := range captures {
		payload.Captures = append(payload.Captures, captureToPayload(captureEntry))
	}

	return &capture.GetCaptureListOK{Payload: payload}
}

func captureToPayload(c repo.CaptureFile) *model.ProjectCaptures {
	payload := model.ProjectCaptures(c)
	return &payload
}

type AddCaptureStartHandler struct{}

func (a AddCaptureStartHandler) Handle(params capture.AddStartCaptureParams) middleware.Responder {
	err := repo.AddStart(captureFromPayload(params.Body))
	if err != nil {
		return &capture.AddStartCaptureInternalServerError{Payload: &model.ServerError{Description: err.Error()}}
	}

	return &capture.AddStartCaptureNoContent{}
}

func captureFromPayload(payload *model.Capture) repo.Capture {
	return repo.Capture(*payload)
}

type AddCaptureStopHandler struct{}

func (a AddCaptureStopHandler) Handle(params capture.AddStopCaptureParams) middleware.Responder {
	err := repo.AddStop(captureFromPayload(params.Body))
	if err != nil {
		return &capture.AddStopCaptureInternalServerError{Payload: &model.ServerError{Description: err.Error()}}
	}

	return &capture.AddStartCaptureNoContent{}
}

type SetCaptureLatestStopHandler struct{}

func (a SetCaptureLatestStopHandler) Handle(params capture.SetLatestStopParams) middleware.Responder {
	err := repo.SetLatestStop(captureFromPayload(params.Body))
	if err != nil {
		return &capture.SetLatestActivityInternalServerError{Payload: &model.ServerError{Description: err.Error()}}
	}

	return &capture.SetLatestActivityNoContent{}
}

type GetReportCapturesTodayHandler struct{}

func (g GetReportCapturesTodayHandler) Handle(params reports.GetReportCapturesTodayParams) middleware.Responder {
	now := time.Now()
	minTime := time.Date(now.Year(), now.Month(), now.Day(), 6, 0, 0, 0, time.UTC)
	maxTime := minTime.Add(time.Hour * 16).Add(time.Nanosecond * -1)

	return getReportCaptures(minTime, maxTime)
}

type GetReportCapturesCurrentMonthHandler struct {
}

func (g GetReportCapturesCurrentMonthHandler) Handle(params reports.GetReportCapturesCurrentMonthParams) middleware.Responder {
	now := time.Now()
	minTime := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	maxTime := minTime.AddDate(0, 1, 0).Add(time.Nanosecond * -1)

	return getReportCaptures(minTime, maxTime)
}

func getReportCaptures(minTime time.Time, maxTime time.Time) middleware.Responder {
	captures, err := repo.GetCaptures()
	if err != nil {
		return &reports.GetReportCapturesTodayInternalServerError{Payload: &model.ServerError{
			Description: err.Error(),
		}}
	}

	payload := &model.ReportCapturesList{}
	for _, c := range captures {
		var secondsWorked int
		var numberOfTimesWorked int64
		for i, start := range c.Starts {
			if minTime.Unix() > start || maxTime.Unix() < start {
				continue
			}

			if len(c.Stops) > i {
				secondsWorked += int(c.Stops[i] - start)
				numberOfTimesWorked++
			}
		}

		if numberOfTimesWorked > 0 {
			payload.Projects = append(payload.Projects, &model.ReportCapturesCapture{
				ID:                  c.ID,
				TimeWorked:          int64(secondsWorked),
				TimeWorkedDisplay:   (time.Duration(secondsWorked) * time.Second).String(),
				NumberOfTimesWorked: numberOfTimesWorked,
			})
		}
	}

	return &reports.GetReportCapturesTodayOK{Payload: payload}
}
