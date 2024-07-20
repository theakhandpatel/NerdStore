package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/resources", app.createResourceHandler)
	router.HandlerFunc(http.MethodGet, "/v1/resources/:id", app.showResourceHandler)
	router.HandlerFunc(http.MethodPut, "/v1/resources/:id", app.updateResourceHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/resources/:id", app.deleteResourceHandler)

	return app.recoverPanic(router)
}
