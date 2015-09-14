package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	. "github.com/francoishill/golang-web-dry/errors/checkerror"
	. "github.com/francoishill/golang-web-dry/middleware/recoverymiddleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/thoas/stats"
	"gopkg.in/tylerb/graceful.v1"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/francoishill/windows-startup-manager/Server/WebApplication/Context/RouterContext"
	. "github.com/francoishill/windows-startup-manager/Server/WebApplication/Logger"
	"github.com/francoishill/windows-startup-manager/Server/WebApplication/RouteHandlers"
	. "github.com/francoishill/windows-startup-manager/Server/WebApplication/RouteHandlers/_Setup"
	. "github.com/francoishill/windows-startup-manager/Server/WebApplication/Settings/Default"
)

func setRouteDefinitions(ctx *RouterContext.RouterContext, router *mux.Router) {
	baseRouterMiddleware := []HttpHandlerFunc{}

	apiV1Router := router.PathPrefix("/api/v1").Subrouter()

	routeDefinitions := RouteHandlers.GetRouteDefinitions(ctx)
	RegisterRouteDefinitions(ctx, apiV1Router, baseRouterMiddleware, routeDefinitions)

	ctx.Settings.SubscribeConfigReloaded(ctx)
}

type tmpRouterRecoveryWrapper struct {
	logger Logger
}

func (t *tmpRouterRecoveryWrapper) onRouterRecoveryError(errDetails *RecoveredErrorDetails) {
	t.logger.Error("ERROR: %s\nStack:\n%s", errDetails.Error, errDetails.StackTrace)
}

func getNegroniHandlers(ctx *RouterContext.RouterContext, router *mux.Router) []negroni.Handler {
	tmpArray := []negroni.Handler{}

	// fullRecoveryStackMessage := GlobalConfigSettings.Common.DevMode

	routerRecoveryWrapper := &tmpRouterRecoveryWrapper{ctx.Logger}

	// tmpArray = append(tmpArray, gzip.Gzip(gzip.DefaultCompression))
	tmpArray = append(tmpArray, NewRecovery(routerRecoveryWrapper.onRouterRecoveryError)) //recovery.JSONRecovery(fullRecoveryStackMessage))
	tmpArray = append(tmpArray, negroni.NewLogger())

	if ctx.Settings.IsDevMode() {
		middleware := stats.New()
		router.HandleFunc("/stats.json", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			stats := middleware.Data()

			b, _ := json.Marshal(stats)

			w.Write(b)
		})
		tmpArray = append(tmpArray, middleware)
	}

	return tmpArray
}

func main() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("-------------------------STARTUP OF SERVER-------------------------")

	if len(os.Args) < 2 {
		panic("Cannot startup, the first argument must be the config path")
	}

	settings := NewDefaultSettings(os.Args[1])
	ctx := RouterContext.New(settings)

	router := mux.NewRouter()
	setRouteDefinitions(ctx, router)

	n := negroni.New(getNegroniHandlers(ctx, router)...) //negroni.Classic()
	n.UseHandler(context.ClearHandler(router))

	ctx.Logger.Info("Now serving on %s", settings.ServerBackendUrl())

	backendUrl, err := url.Parse(settings.ServerBackendUrl())
	CheckError(err)
	hostWithPossiblePortOnly := backendUrl.Host
	if settings.IsDevMode() {
		graceful.Run(hostWithPossiblePortOnly, 0, n)
	} else {
		graceful.Run(hostWithPossiblePortOnly, 5*time.Second, n) //Graceful shutdown to allow 5 seconds to close connections
	}
}
