package movie

type Movie struct {
	Title   string
	Runtime int
}

type IMovie interface {
	SetTitle(title string)
	GetTitle() string
	SetRuntime(runtime int)
	GetRuntime() int
}

func (m *Movie) SetTitle(title string) {
	m.Title = title
}

func (m *Movie) GetTitle() string {
	return m.Title
}

func (m *Movie) SetRuntime(runtime int) {
	m.Runtime = runtime
}

func (m *Movie) GetRuntime() int {
	return m.Runtime
}
