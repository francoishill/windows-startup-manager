package Dashboard

import (
	"fmt"
	"net/http"
	"strings"

	. "github.com/francoishill/windows-startup-manager/Server/WebApplication/Context/RouterContext"
)

func NewRoute() *route {
	return &route{}
}

type route struct{}

func (rt *route) GetPathPart() string {
	return "/dashboard"
}

func (rt *route) Get(w http.ResponseWriter, r *http.Request, ctx *RouterContext) {
	action := r.URL.Query().Get("action")
	if action != "" {
		appIdInt64 := ctx.HttpHelperService.GetRequiredUrlQueryValue_Int64(r, "appid")
		appId := int(appIdInt64)

		switch strings.ToUpper(action) {
		case "KILL":
			ctx.StartupAppsService.KillApp(appId)
			break
		case "RESTART":
			ctx.StartupAppsService.KillApp(appId)
			ctx.StartupAppsService.StartApp(appId)
			break
		case "START":
			ctx.StartupAppsService.StartApp(appId)
			break
		case "ENABLE":
			ctx.StartupAppsService.EnableApp(appId)
			break
		case "DISABLE":
			ctx.StartupAppsService.DisableApp(appId)
			break
		case "CLEAR_STATUS_PROGRESS":
			ctx.StartupAppsService.ClearStatusProgress(appId)
			break
		default:
			panic("Unsupported action: " + action)
		}

		fmt.Println(r.URL.Path)
		http.Redirect(w, r, r.URL.Path, 302)
		return
	}

	currentApps := ctx.StartupAppsService.GetCurrentAppList()
	ctx.HttpHelperService.RenderHtml(w, "Dashboard", currentApps)
}
