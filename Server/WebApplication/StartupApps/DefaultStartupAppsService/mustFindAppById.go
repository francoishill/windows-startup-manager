package DefaultStartupAppsService

import (
	"fmt"

	. "github.com/francoishill/windows-startup-manager/Domain/Entity/StartupApps"
)

func (s *service) mustFindAppById(appId int) *App {
	apps := s.GetCurrentAppList()
	for _, a := range apps {
		if a.TmpId == appId {
			return a
		}
	}

	panic(fmt.Sprintf("App with id %d is not found in list", appId))
}
