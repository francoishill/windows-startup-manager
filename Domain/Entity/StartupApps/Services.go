package StartupApps

type StartupAppsService interface {
	GetCurrentAppList() []*App
	ReloadAppsFromFile()

	IsPaused() bool
	PauseStarting()
	ResumeStarting()

	KillApp(appId int)
	StartApp(appId int)
	EnableApp(appId int)
	DisableApp(appId int)
	ClearStatusProgress(appId int)
}
