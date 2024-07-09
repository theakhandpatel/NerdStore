package main

import (
	"fmt"
	"net/http"
	"time"

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

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showResourcesHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	resource := data.Resource{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Golang intro",
		Link:      "www.golang.com",
		Tags:      []string{"backend", "go"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"resource": resource}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
