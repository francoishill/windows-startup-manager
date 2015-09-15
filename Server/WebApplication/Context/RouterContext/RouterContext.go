package RouterContext

import (
	"github.com/francoishill/windows-startup-manager/Domain/Entity/StartupApps"
	"github.com/francoishill/windows-startup-manager/Server/WebApplication/Errors"
	"github.com/francoishill/windows-startup-manager/Server/WebApplication/Errors/DefaultErrorService"
	"github.com/francoishill/windows-startup-manager/Server/WebApplication/HttpHelper"
	"github.com/francoishill/windows-startup-manager/Server/WebApplication/HttpHelper/DefaultHttpHelperService"
	"github.com/francoishill/windows-startup-manager/Server/WebApplication/Logger"
	DefaultLogger "github.com/francoishill/windows-startup-manager/Server/WebApplication/Logger/Default"
	"github.com/francoishill/windows-startup-manager/Server/WebApplication/Settings"
	"github.com/francoishill/windows-startup-manager/Server/WebApplication/StartupApps/DefaultStartupAppsService"
)

type RouterContext struct {
	Settings           Settings.Settings
	StartupAppsService StartupApps.StartupAppsService

	ErrorService      Errors.ErrorService
	HttpHelperService HttpHelperService.HttpHelperService

	Logger Logger.Logger
}

func (r *RouterContext) OnConfigReloaded() {
	r.loadRepos(r.Settings.UseMock())
}

func (r *RouterContext) setRepos(
	startupAppsService StartupApps.StartupAppsService,
	errorService Errors.ErrorService,
	httpHelperService HttpHelperService.HttpHelperService,
	logger Logger.Logger,
) {

	r.Logger = logger
	r.StartupAppsService = startupAppsService
	r.ErrorService = errorService
	r.HttpHelperService = httpHelperService

	logger.Debug("Router Context reloaded.")
}

func (r *RouterContext) loadRepos(useMock bool) {
	var logger Logger.Logger

	logFileName := "rolling-log.log"
	if useMock {
		logger = DefaultLogger.New(logFileName, "[MOCK]", r.Settings.IsDevMode())
	} else {
		logger = DefaultLogger.New(logFileName, "", r.Settings.IsDevMode())
	}

	startupAppsService := DefaultStartupAppsService.NewDefaultStartupAppsService(logger, r.Settings.CurrentStartupAppsJsonFile())
	startupAppsService.GetCurrentAppList()

	errorService := DefaultErrorService.New()
	httpHelperService := DefaultHttpHelperService.New(r.Settings.IsDevMode())

	r.setRepos(
		startupAppsService,
		errorService,
		httpHelperService,
		logger)
}

func New(settings Settings.Settings) *RouterContext {
	rc := &RouterContext{Settings: settings}
	rc.loadRepos(settings.UseMock())
	return rc
}
