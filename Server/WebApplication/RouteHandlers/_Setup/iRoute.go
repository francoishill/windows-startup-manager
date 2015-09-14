package _Setup

import (
	"net/http"

	"github.com/francoishill/windows-startup-manager/Server/WebApplication/Context/RouterContext"
)

type iRoute interface {
	GetPathPart() string
}

type iOptionsHandler interface {
	Options(w http.ResponseWriter, r *http.Request, ctx *RouterContext.RouterContext)
}

type iGetHandler interface {
	Get(w http.ResponseWriter, r *http.Request, ctx *RouterContext.RouterContext)
}

type iHeadHandler interface {
	Head(w http.ResponseWriter, r *http.Request, ctx *RouterContext.RouterContext)
}

type iPostHandler interface {
	Post(w http.ResponseWriter, r *http.Request, ctx *RouterContext.RouterContext)
}

type iPutHandler interface {
	Put(w http.ResponseWriter, r *http.Request, ctx *RouterContext.RouterContext)
}

type iDeleteHandler interface {
	Delete(w http.ResponseWriter, r *http.Request, ctx *RouterContext.RouterContext)
}

func GetIRouteControllers(route iRoute) map[string]RouteController {
	m := make(map[string]RouteController)

	if o, ok := route.(iOptionsHandler); ok {
		m["OPTIONS"] = o.Options
	}

	if g, ok := route.(iGetHandler); ok {
		m["GET"] = g.Get
	}

	if h, ok := route.(iHeadHandler); ok {
		m["HEAD"] = h.Head
	}

	if h, ok := route.(iPostHandler); ok {
		m["POST"] = h.Post
	}

	if h, ok := route.(iPutHandler); ok {
		m["PUT"] = h.Put
	}

	if h, ok := route.(iDeleteHandler); ok {
		m["DELETE"] = h.Delete
	}

	return m
}
