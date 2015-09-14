package DefaultStartupAppsService

import (
	"fmt"
	. "github.com/francoishill/golang-web-dry/errors/checkerror"
	"os"
	"os/exec"

	. "github.com/francoishill/windows-startup-manager/Domain/Entity/StartupApps"
)

func (s *service) startAndWaitForApp(app *App, triggerWaitGroupDone bool) {
	defer func() {
		if r := recover(); r != nil {
			s.setAppStatus(app, fmt.Sprintf("Error: %+v", r))
		}
	}()

	var err error

	app.IsRunning = false
	app.HasError = false
	app.CurrentProcessId = -1

	possibleRunningProcesses := s.findRunningProcesses(app)
	mostLikelyAlreadyRunningProcess, parentIdIsThisApp := possibleRunningProcesses.findMostLikelyAlreadyRunningProcess(app)
	if mostLikelyAlreadyRunningProcess != nil {
		if parentIdIsThisApp {
			s.setAppStatus(app, fmt.Sprintf("Process is already running with Pid %d (started by this application)", mostLikelyAlreadyRunningProcess.ProcessId))
		} else {
			s.setAppStatus(app, fmt.Sprintf("Process is already running with Pid %d", mostLikelyAlreadyRunningProcess.ProcessId))
		}
		proc, err := os.FindProcess(mostLikelyAlreadyRunningProcess.ProcessId)
		CheckError(err)

		app.IsRunning = true
		app.CurrentProcessId = mostLikelyAlreadyRunningProcess.ProcessId
		s.setAppStatus(app, "Running")
		defer func() { app.IsRunning = false }()

		if triggerWaitGroupDone {
			s.wg.Done()
		}

		_, err = proc.Wait()
		if err != nil {
			app.HasError = true
			s.setAppStatus(app, fmt.Sprintf("Error in waiting for (already started) app %s, error: %s", app.Name, err.Error()))
			return
		} else {
			app.HasError = false
			s.setAppStatus(app, "Application completed (exited)")
		}
	} else {
		if app.Disabled {
			s.setAppStatus(app, "Skipping start, app is disabled")
			if triggerWaitGroupDone {
				s.wg.Done()
			}
			return
		}

		cmd := exec.Command(app.Exe, app.Args...)

		s.setAppStatus(app, "Attempting to start")
		err = cmd.Start()
		CheckError(err)

		app.IsRunning = true
		if cmd.Process != nil {
			app.CurrentProcessId = cmd.Process.Pid
		}

		s.setAppStatus(app, "Running")
		defer func() { app.IsRunning = false }()

		if triggerWaitGroupDone {
			s.wg.Done()
		}

		err = cmd.Wait()
		if err != nil {
			app.HasError = true
			s.setAppStatus(app, fmt.Sprintf("Error in waiting for app %s, error: %s", app.Name, err.Error()))
			return
		} else {
			app.HasError = false
			s.setAppStatus(app, "Application completed (exited)")
		}
	}
}
