package _Setup

import (
	"net/http"

	"github.com/francoishill/windows-startup-manager/Server/WebApplication/Context/RouterContext"
)

type RouteController func(w http.ResponseWriter, r *http.Request, ctx *RouterContext.RouterContext)

func CreateHttpHandlerFunc(ctx *RouterContext.RouterContext, fn RouteController) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, ctx)
	})
}
