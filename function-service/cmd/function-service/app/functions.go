package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) ShowFunction(
	w http.ResponseWriter,
	r *http.Request,
	params httprouter.Params,
) {
	http.NotFound(w, r)
	// id := params.ByName("id")
	// fmt.Fprint(w, "function id: "+id)
}
