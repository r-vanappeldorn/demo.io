package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mholt/binding"
)

type createFunctionReq struct {
	Title  string `json:"title" binding:"required"`
	Salary string `json:"salary" binding:"required, number"`
}

func (app *Application) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	createReq := new(createFunctionReq)
	err := binding.Bind(r, createReq )
}