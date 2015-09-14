package DefaultStartupAppsService

import (
	"fmt"
	. "github.com/francoishill/windows-startup-manager/Domain/Entity/StartupApps"
	"os/exec"
	"path/filepath"
	"strings"
)

func (s *service) findRunningProcesses(app *App) wmicResultProcessSlice {
	exeNameOnly := filepath.Base(app.Exe)

	cmd := exec.Command(
		"wmic",
		"process",
		"where",
		fmt.Sprintf("name='%s'", exeNameOnly),
		"get",
		"Caption,CommandLine,Description,ExecutablePath,Name,ParentProcessId,ProcessId",
		"/format:list")
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(fmt.Sprintf("Unable to run wmic for name '%s'. Error was: %s. CombinedOutput was: %s", exeNameOnly, err.Error(), string(output)))
	}

	return s.parseWmicOutput(strings.Split(string(output), "\n"))
}
