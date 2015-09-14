package Default

import (
	. "github.com/francoishill/windows-startup-manager/Server/WebApplication/Settings"
	"sync"
)

func NewDefaultSettings(configPath string) Settings {
	cfg := LoadConfig(configPath)
	dCfg := &defaultSettings{}
	dCfg.config = cfg
	WatchConfig(dCfg, configPath)
	return dCfg
}

type defaultSettings struct {
	*config

	reloadEventNotifyLock sync.Mutex
	subscribersOnReload   []ConfigReloadHandler
}

func (d *defaultSettings) SetConfig(cfg *config) { d.config = cfg }

func (d *defaultSettings) IsDevMode() bool          { return d.config.Common.DevMode }
func (d *defaultSettings) UseMock() bool            { return d.config.Common.UseMock }
func (d *defaultSettings) ServerBackendUrl() string { return d.config.Server.BackendUrl }
func (d *defaultSettings) CurrentStartupAppsJsonFile() string {
	return d.config.Apps.CurrentStartupAppsJsonFile
}
