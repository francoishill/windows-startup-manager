package DefaultStartupAppsService

import (
	. "github.com/francoishill/golang-web-dry/errors/checkerror"
	"os"

	. "github.com/francoishill/windows-startup-manager/Domain/Entity/StartupApps"
)

func (s *service) findAppProcess(app *App) *os.Process {
	proc, err := os.FindProcess(app.CurrentProcessId)
	CheckError(err)
	return proc
}
