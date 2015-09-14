package DefaultStartupAppsService

import (
	"sync"
)

var newIdLock *sync.RWMutex
var _curAppId int = 0

func generateNewAppId() int {
	newIdLock.Lock()
	defer newIdLock.Unlock()

	_curAppId++
	return _curAppId
}

func init() {
	newIdLock = &sync.RWMutex{}
}
