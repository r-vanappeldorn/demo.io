package app

import "fmt"

const (
	prefix                         = "api/auth%s"
	setPathStartWithSlaceError     = "%s routes need to start with a /, got %q"
	setPathErrorEndsWithSlaceError = "%s routes can not and with a /, got %q"
)

// checks if path is a valid route name and returns a string with the correct prefix
func (app *Application) SetPath(path ...string) string {
	switch len(path) {
	case 1:
		if []byte(path[0])[0] != '/' {
			app.Logger.Fatalf(setPathStartWithSlaceError, "[ERROR]", path[0])
		}

		if []byte(path[0])[len(path[0])-1] == '/' {
			app.Logger.Fatalf(setPathErrorEndsWithSlaceError, "[ERROR]", path[0])
		}
		return fmt.Sprintf(prefix, path[0])
	default:
		return fmt.Sprintf(prefix, "")
	}
}
