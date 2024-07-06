package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/theakhandpatel/NerdStore/internal/data"
)

func (app *application) createResourceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
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
