package DefaultStartupAppsService

import (
	"fmt"
	. "github.com/francoishill/windows-startup-manager/Domain/Entity/StartupApps"
	"sync"
	"time"
)

var globalAppStatusLock *sync.RWMutex

func (s *service) setAppStatus(app *App, status string) {
	s.logger.Info("Setting status for app '%s': %s", app.Name, status)

	globalAppStatusLock.Lock()
	defer globalAppStatusLock.Unlock()

	prefixedStatus := fmt.Sprintf("[%s] %s", time.Now().Format("2006-01-02 15:04:05"), status)

	app.CurrentStatus = prefixedStatus
	app.StatusProgress = append(app.StatusProgress, prefixedStatus)
}

func init() {
	globalAppStatusLock = &sync.RWMutex{}
}
