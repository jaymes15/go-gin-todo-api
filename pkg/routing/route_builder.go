package routing

import (
	"todo/pkg/sessions"
	"todo/pkg/static"
)

func RouteBuilder() {
	Init()
	route := GetRouter()
	sessions.Init(route)
	static.LoadStatic(route)
	RegisterRoutes(route)
	Serve(route)

}
