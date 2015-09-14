package StartupApps

type App struct {
	TmpId    int
	Name     string
	Exe      string
	Args     []string
	Disabled bool

	CurrentProcessId int
	IsRunning        bool
	HasError         bool
	StatusProgress   []string
	CurrentStatus    string
}
