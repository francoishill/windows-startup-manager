package Settings

type ConfigReloadHandler interface {
	OnConfigReloaded()
}

type reloaderSettings interface {
	ConfigReloaded()
	SubscribeConfigReloaded(handler ConfigReloadHandler)
	UnsubscribeConfigReloaded(handler ConfigReloadHandler)
}

type Settings interface {
	reloaderSettings

	IsDevMode() bool
	UseMock() bool
	ServerBackendUrl() string
	CurrentStartupAppsJsonFile() string
}
