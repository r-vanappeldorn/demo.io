package app

import "account-service/pkg/middlewares"

func (app *Application) Router() {
	app.Engine.Use(middlewares.ContentType)
	app.Engine.Use(middlewares.CORSMiddleware)

	app.Engine.POST(app.SetPath("/signup"), app.Routes.SignUp)
	app.Engine.POST(app.SetPath("/signin"), app.Routes.Signin)
}
