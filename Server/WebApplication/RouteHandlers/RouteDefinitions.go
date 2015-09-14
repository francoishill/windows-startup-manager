package RouteHandlers

import (
	. "github.com/francoishill/windows-startup-manager/Server/WebApplication/Context/RouterContext"
	"github.com/francoishill/windows-startup-manager/Server/WebApplication/RouteHandlers/Dashboard"
	. "github.com/francoishill/windows-startup-manager/Server/WebApplication/RouteHandlers/_Setup"
)

func dashboardRouteDefinitions(ctx *RouterContext) []*RouteDefinition {
	return []*RouteDefinition{
		NewRouteDefinition_Routes(Dashboard.NewRoute()),
	}
}

func GetRouteDefinitions(ctx *RouterContext) []*RouteDefinition {
	routeDefinitions := []*RouteDefinition{}
	routeDefinitions = append(routeDefinitions, dashboardRouteDefinitions(ctx)...)
	return routeDefinitions
}
