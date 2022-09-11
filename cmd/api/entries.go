//Filename: cmd/api/entries.go

package main

import (
	"fmt"
	"net/http"
)

//createEntryHandler for the "POST /v1/entries" endpoint
func (app *application) createEntryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is a test...")

}

//showEntryHandler for the "GET /v1/entries/:id" endpoint
func (app *application) showEntryHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	//Display the entries id
	fmt.Fprintf(w, "show the details for entries %d\n", id)
}
