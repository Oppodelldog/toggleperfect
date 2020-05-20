// This file is safe to edit. Once it exists it will not be overwritten

package server

import (
	"crypto/tls"
	"net/http"
	"strings"

	"github.com/Oppodelldog/toggleperfect/internal/apps/timetoggle/apidocs"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/Oppodelldog/toggleperfect/internal/apps/timetoggle/api/server/api"
	"github.com/Oppodelldog/toggleperfect/internal/apps/timetoggle/api/server/api/project"
)

//go:generate swagger generate server --target ../../../timetoggle --name Timetoggle --spec ../../swagger.yml --api-package api --model-package api/model --server-package api/server --exclude-main

//noinspection ALL
func configureFlags(api *api.TimetoggleAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *api.TimetoggleAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()
	api.XMLConsumer = runtime.XMLConsumer()

	api.JSONProducer = runtime.JSONProducer()
	api.XMLProducer = runtime.XMLProducer()

	if api.ProjectAddProjectHandler == nil {
		api.ProjectAddProjectHandler = project.AddProjectHandlerFunc(func(params project.AddProjectParams) middleware.Responder {
			return middleware.NotImplemented("operation project.AddProject has not yet been implemented")
		})
	}
	if api.ProjectDeleteProjectHandler == nil {
		api.ProjectDeleteProjectHandler = project.DeleteProjectHandlerFunc(func(params project.DeleteProjectParams) middleware.Responder {
			return middleware.NotImplemented("operation project.DeleteProject has not yet been implemented")
		})
	}
	if api.ProjectGetProjectByIDHandler == nil {
		api.ProjectGetProjectByIDHandler = project.GetProjectByIDHandlerFunc(func(params project.GetProjectByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation project.GetProjectByID has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
//noinspection ALL
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
//noinspection ALL
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return addApiDocs(handler)
}

func addApiDocs(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isApiDocsRoute(r) {
			if isApiDocsRoot(r) {
				redirectToIndexPage(w)
				return
			}
			handleApiDocsFileServer(w, r)
		} else {
			handler.ServeHTTP(w, r)
		}
	})
}

func handleApiDocsFileServer(w http.ResponseWriter, r *http.Request) {
	apidocs.NewHandler().ServeHTTP(w, r)
}

func redirectToIndexPage(w http.ResponseWriter) {
	w.Header().Set("Location", "/apidocs/ui.html")
	w.WriteHeader(http.StatusFound)
}

func isApiDocsRoot(r *http.Request) bool {
	return r.RequestURI == "/apidocs" || r.RequestURI == "/apidocs/"
}

func isApiDocsRoute(r *http.Request) bool {
	return strings.HasPrefix(r.RequestURI, "/apidocs")
}
