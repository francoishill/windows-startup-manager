package _Setup

import (
	"github.com/gorilla/mux"

	"github.com/francoishill/windows-startup-manager/Server/WebApplication/Context/RouterContext"
)

func newSliceFromHttpHandlerFuncs(inSlice []HttpHandlerFunc) []HttpHandlerFunc {
	newSlice := []HttpHandlerFunc{}
	newSlice = append(newSlice, inSlice...)
	return newSlice
}

func setupRouters(ctx *RouterContext.RouterContext, router *mux.Router, parentMiddleWare []HttpHandlerFunc, routeDefinitions []*RouteDefinition) {
	if len(routeDefinitions) == 0 {
		return
	}

	for _, rh := range routeDefinitions {
		var routerToUse *mux.Router
		if rh.prefix != "" {
			routerToUse = router.PathPrefix(rh.prefix).Subrouter()
		} else {
			routerToUse = router
		}

		combinedMiddleWareHandlers := []HttpHandlerFunc{}
		combinedMiddleWareHandlers = append(combinedMiddleWareHandlers, parentMiddleWare...)
		combinedMiddleWareHandlers = append(combinedMiddleWareHandlers, rh.middlewares...)

		for _, singleRouteDefinition := range rh.routeDefinitions {
			for method, h := range GetIRouteControllers(singleRouteDefinition) {
				muxRoute := routerToUse.Handle(singleRouteDefinition.GetPathPart(), newNegroniMiddleware(ctx, combinedMiddleWareHandlers, h))
				muxRoute.Methods(method)
			}
		}

		setupRouters(ctx, routerToUse, combinedMiddleWareHandlers, rh.subRoutes)
	}
}

func RegisterRouteDefinitions(ctx *RouterContext.RouterContext, router *mux.Router, baseMiddleWare []HttpHandlerFunc, definitions []*RouteDefinition) {
	//Routers
	setupRouters(ctx, router, baseMiddleWare, definitions)
}
