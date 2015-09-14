package DefaultStartupAppsService

import (
	. "github.com/francoishill/windows-startup-manager/Domain/Entity/StartupApps"
	"os"
)

type wmicResultProcessSlice []*wmicResultProcess

func (w wmicResultProcessSlice) findFirstWithParentPid(parentPidToFind int) *wmicResultProcess {
	for _, rp := range w {
		if rp.ParentProcessId == parentPidToFind {
			return rp
		}
	}
	return nil
}

func (w wmicResultProcessSlice) findMostLikelyAlreadyRunningProcess(app *App) (proc *wmicResultProcess, parentIdIsThisApp bool) {
	if len(w) == 0 {
		return nil, false
	}

	if len(w) == 1 {
		return w[0], false
	}

	thisAppPid := os.Getpid()
	processWithThisAsParent := w.findFirstWithParentPid(thisAppPid)
	if processWithThisAsParent != nil {
		return processWithThisAsParent, true
	}

	for _, p := range w {
		if p.exeEquals(app.Exe) {
			return p, false
		}
	}

	return nil, false
}
