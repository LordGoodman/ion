package main

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
)

func main() {
	app := ion.New()

	none := app.None("/invisible/{username}", func(ctx context.Context) {
		ctx.Writef("Hello %s with method: %s", ctx.Values().GetString("username"), ctx.Method())

		if from := ctx.Values().GetString("from"); from != "" {
			ctx.Writef("\nI see that you're coming from %s", from)
		}
	})

	app.Get("/change", func(ctx context.Context) {

		if none.IsOnline() {
			none.Method = ion.MethodNone
		} else {
			none.Method = ion.MethodGet
		}

		// refresh re-builds the router at serve-time in order to be notified for its new routes.
		app.RefreshRouter()
	})

	app.Get("/execute", func(ctx context.Context) {
		// same as navigating to "http://localhost:8080/invisible/ion" when /change has being invoked and route state changed
		// from "offline" to "online"
		ctx.Values().Set("from", "/execute") // values and session can be shared when calling Exec from a "foreign" context.
		ctx.Exec("GET", "/invisible/ion")
	})

	app.Run(ion.Addr(":8080"))
}
