package app

type Env interface {
	Prepare()
}

func NewEnv() Env {
	return &env{
		children: []Env{},
	}
}

type env struct {
	children []Env
}

func (e *env) Prepare() {
	for _, child := range e.children {
		child.Prepare()
	}
}

// ------------------------------------------------------------------

func NewRDBSetting()

type RDBSetting struct {
	User        string
	Password    string
	InstanceStr string
	DBName      string
	// TODO: その他、MaxConnection等
}

func (s *RDBSetting) Prepare() {

}
