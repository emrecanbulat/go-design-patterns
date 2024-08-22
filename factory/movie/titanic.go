package movie

type Titanic struct {
	Movie
}

func NewTitanic() IMovie {
	movie := &Movie{}
	movie.SetTitle("Titanic")
	movie.SetRuntime(120)

	return &Titanic{*movie}
}
