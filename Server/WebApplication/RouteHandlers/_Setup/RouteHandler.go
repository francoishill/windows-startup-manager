package _Setup

import (
	"github.com/codegangsta/negroni"
	"net/http"

	"github.com/francoishill/windows-startup-manager/Server/WebApplication/Context/RouterContext"
)

type HttpHandlerFunc func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)

func newNegroniMiddleware(ctx *RouterContext.RouterContext, middleWareFuncs []HttpHandlerFunc, controller RouteController) *negroni.Negroni {
	negroniHandlers := []negroni.Handler{}
	for ind, _ := range middleWareFuncs {
		negroniHandlers = append(negroniHandlers, negroni.HandlerFunc(middleWareFuncs[ind]))
	}
	negroniHandlers = append(negroniHandlers, negroni.Wrap(CreateHttpHandlerFunc(ctx, controller)))
	return negroni.New(negroniHandlers...)
}

type RouteDefinition struct {
	prefix           string
	middlewares      []HttpHandlerFunc
	routeDefinitions []iRoute
	subRoutes        []*RouteDefinition
}

func NewRouteDefinition_Prefix_Middles_Routes_Subroutes(prefix string, middlewares []HttpHandlerFunc, routeDefinitions []iRoute, subRoutes []*RouteDefinition) *RouteDefinition {
	return &RouteDefinition{
		prefix:           prefix,
		middlewares:      middlewares,
		routeDefinitions: routeDefinitions,
		subRoutes:        subRoutes,
	}
}

func NewRouteDefinition_Routes(routeDefinitions ...iRoute) *RouteDefinition {
	return NewRouteDefinition_Prefix_Middles_Routes_Subroutes("", nil, routeDefinitions, nil)
}

func NewRouteDefinition_Prefix_Routes(prefix string, routeDefinitions ...iRoute) *RouteDefinition {
	return NewRouteDefinition_Prefix_Middles_Routes_Subroutes(prefix, nil, routeDefinitions, nil)
}

func NewRouteDefinition_Prefix_Subroutes(prefix string, subRoutes []*RouteDefinition) *RouteDefinition {
	return NewRouteDefinition_Prefix_Middles_Routes_Subroutes(prefix, nil, nil, subRoutes)
}

func NewRouteDefinition_Middles_Routes(middlewares []HttpHandlerFunc, routeDefinitions ...iRoute) *RouteDefinition {
	return NewRouteDefinition_Prefix_Middles_Routes_Subroutes("", middlewares, routeDefinitions, nil)
}

func NewRouteDefinition_Prefix_Middles_Routes(prefix string, middlewares []HttpHandlerFunc, routeDefinitions ...iRoute) *RouteDefinition {
	return NewRouteDefinition_Prefix_Middles_Routes_Subroutes(prefix, middlewares, routeDefinitions, nil)
}
