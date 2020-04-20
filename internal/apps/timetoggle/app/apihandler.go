package app

import (
	"github.com/Oppodelldog/toggleperfect/internal/apps/timetoggle/api/model"
	"github.com/Oppodelldog/toggleperfect/internal/apps/timetoggle/api/server/api/capture"
	"github.com/Oppodelldog/toggleperfect/internal/apps/timetoggle/api/server/api/project"
	"github.com/Oppodelldog/toggleperfect/internal/apps/timetoggle/app/repo"
	"github.com/go-openapi/runtime/middleware"
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
