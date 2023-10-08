package routing

import (
	"todo/pkg/cors"
	"todo/pkg/sessions"
	"todo/pkg/static"
)

func RouteBuilder() {
	Init()
	route := GetRouter()
	sessions.Init(route)
	static.LoadStatic(route)
	cors.UseCors(route)
	RegisterRoutes(route)
	Serve(route)

}
