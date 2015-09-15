package Dashboard

import (
	"fmt"
	"net/http"
	"strings"

	. "github.com/francoishill/windows-startup-manager/Domain/Entity/StartupApps"
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
	globalAction := r.URL.Query().Get("global_action")
	if globalAction != "" {
		switch strings.ToUpper(globalAction) {
		case "RELOAD_APPS_FROM_CONFIG":
			ctx.StartupAppsService.ReloadAppsFromFile()
			break
		case "PAUSE_STARTING":
			ctx.StartupAppsService.PauseStarting()
			break
		case "RESUME_STARTING":
			ctx.StartupAppsService.ResumeStarting()
			break
		default:
			panic("Unsupported global action: " + globalAction)
		}
	}

	appAction := r.URL.Query().Get("app_action")
	if appAction != "" {
		appIdInt64 := ctx.HttpHelperService.GetRequiredUrlQueryValue_Int64(r, "appid")
		appId := int(appIdInt64)

		switch strings.ToUpper(appAction) {
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
			panic("Unsupported app action: " + appAction)
		}

		fmt.Println(r.URL.Path)
		http.Redirect(w, r, r.URL.Path, 302)
		return
	}

	data := &struct {
		IsPaused    bool
		CurrentApps []*App
	}{
		ctx.StartupAppsService.IsPaused(),
		ctx.StartupAppsService.GetCurrentAppList(),
	}
	ctx.HttpHelperService.RenderHtml(w, "Dashboard", data)
}
