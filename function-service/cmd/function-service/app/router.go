package app

import (
	"fmt"

	"function-srv/pkg/messages"

	"github.com/julienschmidt/httprouter"
)

const (
	routePrefix                    = "/api/functions%s"
	setPathStartWithSlaceError     = "%s routes need to start with a /, got %q"
	setPathErrorEndsWithSlaceError = "%s routes can not and with a /, got %q"
)

func (app *Application) Router() *httprouter.Router {
	r := httprouter.New()

	r.GET(app.SetPath("/:id"), app.ShowFunction)

	return r
}

// checks if path is a valid route name and returns a string with the correct prefix
func (app *Application) SetPath(path ...string) string {
	switch len(path) {
	case 1:
		if []byte(path[0])[0] != '/' {
			app.Logger.Fatalf(setPathStartWithSlaceError, messages.Error(), path[0])
		}

		if []byte(path[0])[len(path[0])-1] == '/' {
			app.Logger.Fatalf(setPathErrorEndsWithSlaceError, messages.Error(), path[0])
		}
		return fmt.Sprintf(routePrefix, path[0])
	default:
		return fmt.Sprintf(routePrefix, "")
	}
}
