package app

import (
	"github.com/Oppodelldog/toggleperfect/internal/apps/timetoggle/api/model"
	"github.com/Oppodelldog/toggleperfect/internal/apps/timetoggle/api/server/api/project"
	"github.com/go-openapi/runtime/middleware"
)

type AddProjectHandler struct {
}

func (a AddProjectHandler) Handle(params project.AddProjectParams) middleware.Responder {
	//TODO: implement

	return &project.AddProjectNoContent{}
}

type UpdateProjectHandler struct {
}

func (u UpdateProjectHandler) Handle(params project.UpdateProjectParams) middleware.Responder {
	//TODO: implement

	return &project.UpdateProjectNoContent{}
}

type DeleteProjectHandler struct {
}

func (d DeleteProjectHandler) Handle(params project.DeleteProjectParams) middleware.Responder {

	//TODO: implement

	return &project.DeleteProjectNoContent{}
}

type GetProjectHandler struct {
}

func (g GetProjectHandler) Handle(params project.GetProjectByIDParams) middleware.Responder {
	//TODO: implement

	return &project.GetProjectByIDOK{
		Payload: &model.Project{
			Description: "My very first project!",
			ID:          "dsgmipgdnpgdg",
		},
	}
}
