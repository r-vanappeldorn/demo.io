package app

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type createFunctionReq struct {
	Title  string `json:"title" validate:"required"`
	Salary string `json:"salary" validate:"required"`
}

func (app *Application) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var req createFunctionReq

	if err := DecodeJSON(r.Body, &req); err != nil {
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "success")
}
