package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/theakhandpatel/NerdStore/internal/data"
	"github.com/theakhandpatel/NerdStore/internal/validator"
)

func (app *application) createResourceHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string   `json:"title"`
		Link  string   `json:"link"`
		Tags  []string `json:"tags"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	resource := &data.Resource{
		Title: input.Title,
		Link:  input.Link,
		Tags:  input.Tags,
	}

	v := validator.New()

	if data.ValidateResource(v, resource); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Resources.Insert(resource)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/resources/%d", resource.ID))
	err = app.writeJSON(w, http.StatusCreated, envelope{"resource": resource}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showResourcesHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	resource, err := app.models.Resources.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"resource": resource}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
