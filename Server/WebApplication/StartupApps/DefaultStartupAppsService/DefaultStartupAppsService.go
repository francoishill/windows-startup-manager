package DefaultStartupAppsService

import (
	. "github.com/francoishill/windows-startup-manager/Domain/Entity/StartupApps"
	. "github.com/francoishill/windows-startup-manager/Server/WebApplication/Logger"
	"sync"
	"time"
)

type service struct {
	logger       Logger
	jsonFilePath string
	lock         *sync.RWMutex
	currentApps  []*App
	wg           *sync.WaitGroup

	isPaused     bool
	isPausedLock *sync.RWMutex
}

func (s *service) setPaused(paused bool) {
	s.isPausedLock.Lock()
	defer s.isPausedLock.Unlock()
	s.isPaused = paused
}

func (s *service) startAllApps() {
	s.wg = &sync.WaitGroup{}
	s.wg.Add(len(s.currentApps))

	for _, app := range s.currentApps {
		for s.isPaused {
		}

		go s.startAndWaitForApp(app, true)
		if !app.Disabled {
			time.Sleep(3 * time.Second)
		}
	}

	s.wg.Wait()
}

func (s *service) GetCurrentAppList() []*App {
	if s.currentApps == nil {
		currentStartupApps := s.loadCurrentStartupApps()

		s.currentApps = currentStartupApps
		go s.startAllApps()
	}
	return s.currentApps
}

func (s *service) ReloadAppsFromFile() {
	s.currentApps = nil
	s.GetCurrentAppList()
}

func (s *service) IsPaused() bool {
	return s.isPaused
}

func (s *service) PauseStarting() {
	s.setPaused(true)
}
func (s *service) ResumeStarting() {
	s.setPaused(false)
}

func (s *service) KillApp(appId int) {
	app := s.mustFindAppById(appId)
	proc := s.findAppProcess(app)

	s.setAppStatus(app, "Killing")
	err := proc.Kill()
	if err != nil {
		s.setAppStatus(app, "Unable to kill, error: "+err.Error())
	}
}

func (s *service) StartApp(appId int) {
	app := s.mustFindAppById(appId)
	go s.startAndWaitForApp(app, false)
}

func (s *service) EnableApp(appId int) {
	app := s.mustFindAppById(appId)
	if app.Disabled {
		app.Disabled = false
		s.saveCurrentStartupApps()
	}
}

func (s *service) DisableApp(appId int) {
	app := s.mustFindAppById(appId)
	if !app.Disabled {
		app.Disabled = true
		s.saveCurrentStartupApps()
	}
}

func (s *service) ClearStatusProgress(appId int) {
	app := s.mustFindAppById(appId)
	app.StatusProgress = []string{}
}

func NewDefaultStartupAppsService(logger Logger, jsonFilePath string) StartupAppsService {
	return &service{
		logger,
		jsonFilePath,
		&sync.RWMutex{},
		nil,
		nil,
		false,
		&sync.RWMutex{},
	}
}
