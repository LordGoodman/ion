package main

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
)

func main() {
	app := ion.New()
	app.Get("/", func(ctx context.Context) {
		ctx.HTML("<b>Hello!</b>")
	})
	// [...]

	// Good when you have two configurations, one for development and a different one for production use.
	app.Run(ion.Addr(":8080"), ion.WithConfiguration(ion.YAML("./configs/ion.yml")))

	// or before run:
	// app.Configure(ion.WithConfiguration(ion.YAML("./configs/ion.yml")))
	// app.Run(ion.Addr(":8080"))
}
