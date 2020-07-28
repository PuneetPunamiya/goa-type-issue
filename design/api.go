package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("hub", func() {
	Title("Goa Framework")
	Description("HTTP services for managing Goa Framework")
	Version("0.1")
	Meta("swagger:example", "false")
	Server("GOA", func() {
		Host("production", func() {
			URI("http://api.goa-type.dev")
		})

		Services("resource")
	})

	cors.Origin("*", func() {
		cors.Headers("Content-Type")
		cors.Methods("GET")
	})
})
