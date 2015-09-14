package DefaultStartupAppsService

import (
	"encoding/json"
	. "github.com/francoishill/golang-web-dry/errors/checkerror"
	"io/ioutil"

	. "github.com/francoishill/windows-startup-manager/Domain/Entity/StartupApps"
)

type tmpApp struct {
	Name     string
	Exe      string
	Args     []string `json:",omitempty"`
	Disabled bool     `json:",omitempty"`
}

type tmpApps struct {
	Apps []*tmpApp
}

func (s *service) loadCurrentStartupApps() []*App {
	s.lock.Lock()
	defer s.lock.Unlock()

	fileContentBytes, err := ioutil.ReadFile(s.jsonFilePath)
	CheckError(err)

	currentApps := &tmpApps{}
	err = json.Unmarshal(fileContentBytes, currentApps)
	CheckError(err)

	apps := []*App{}
	for _, a := range currentApps.Apps {
		apps = append(apps, &App{
			TmpId:    generateNewAppId(),
			Name:     a.Name,
			Exe:      a.Exe,
			Args:     a.Args,
			Disabled: a.Disabled,
		})
	}
	return apps
}

func (s *service) saveCurrentStartupApps() {
	s.lock.Lock()
	defer s.lock.Unlock()

	allApps := []*tmpApp{}
	for _, a := range s.currentApps {
		allApps = append(allApps, &tmpApp{
			Name:     a.Name,
			Exe:      a.Exe,
			Args:     a.Args,
			Disabled: a.Disabled,
		})
	}

	jsonBytes, err := json.MarshalIndent(&tmpApps{allApps}, "", "    ")
	CheckError(err)

	err = ioutil.WriteFile(s.jsonFilePath, jsonBytes, 0600)
	CheckError(err)
}
