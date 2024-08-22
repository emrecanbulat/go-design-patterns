package movie

type JohnWick struct {
	Movie
}

func NewJohnWick() IMovie {
	movie := &Movie{}
	movie.SetTitle("John Wick")
	movie.SetRuntime(105)

	return &JohnWick{*movie}
}
