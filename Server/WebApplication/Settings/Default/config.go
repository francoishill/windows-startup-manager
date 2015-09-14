package Default

import (
	"github.com/howeyc/fsnotify"
	"github.com/ian-kent/go-log/log"
	"time"

	. "github.com/francoishill/golang-web-dry/errors/checkerror"

	"gopkg.in/sconf/ini.v0"
	"gopkg.in/sconf/sconf.v0"
)

var lastConfigModifiedTime time.Time

type config struct {
	Common struct {
		DevMode bool
		UseMock bool
	}
	Server struct {
		BackendUrl string
	}
	Apps struct {
		CurrentStartupAppsJsonFile string
	}
}

func (c *config) Validate() {
	// Do nothing for now
}

type iConfigWrapper interface {
	SetConfig(cfg *config)
	ConfigReloaded()
}

func WatchConfig(configWrapper iConfigWrapper, configPath string) {
	watcherForConfigFile, err := fsnotify.NewWatcher()
	CheckError(err)

	go func() {
		for {
			select {
			case e := <-watcherForConfigFile.Event:
				if lastConfigModifiedTime.Add(1 * time.Second).After(time.Now()) {
					continue
				}
				lastConfigModifiedTime = time.Now()

				// log.Warn("Config '%s' modified, but NOT RELOADING YET.", e.Name)
				log.Warn("Config '%s' modified, now reloading (most of the) config settings.", e.Name)
				configWrapper.SetConfig(loadFile(configPath))
				configWrapper.ConfigReloaded()
				break
			case err := <-watcherForConfigFile.Error:
				log.Error("Error: %+v", err)
				break
			}
		}
	}()

	err = watcherForConfigFile.Watch(configPath)
	CheckError(err)
}

func loadFile(configPath string) *config {
	cfg := &config{}
	err := sconf.Must(cfg).Read(ini.File(configPath))
	CheckError(err)
	cfg.Validate()
	return cfg
}

func LoadConfig(configPath string) *config {
	cfg := loadFile(configPath)
	return cfg
}
