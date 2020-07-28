package design

import (
	. "goa.design/goa/v3/dsl"
)

var Resource = ResultType("application/vnd.hub.resource", func() {
	Description("The resource type describes resource information.")
	TypeName("Resource")

	Attribute("id", UInt, "ID is the unique id of the resource", func() {
		Example("id", 1)
	})
	Attribute("name", String, "Name of resource", func() {
		Example("name", "apple")
	})
	Attribute("version", ArrayOf("ResVersion"))

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("version")
	})

	Required("id", "name", "version")
})

var ResVersion = ResultType("application/vnd.hub.version", func() {
	Description("The Version result type describes resource's version information.")
	TypeName("ResVersion")

	Attribute("id", UInt, "ID is the unique id of resource's version", func() {
		Example("id", 1)
	})
	Attribute("version", String, "Version of resource", func() {
		Example("version", "0.1")
	})
	Attribute("resource", Resource, "Resource to which the version belongs", func() {
		View("info")
		Example("resource", func() {
			Value(Val{
				"id":   1,
				"name": "apple",
			})
		})
	})

	View("urls", func() {
		Attribute("id")
		Attribute("version")
	})

	View("default", func() {
		Attribute("id")
		Attribute("version")
		Attribute("resource")
	})

	Required("id", "version", "resource")
})
