package internal

type Movie struct {
	Id       int
	Name     string
	Theatres []int
}

type MovieBrowser interface {
	ListMovies() []Movie
	GetMovie(name string) Movie
	GetTheatre(name string) Theatre
}

type movieBrowser struct {
	movies   map[string]Movie
	theatres map[string]Theatre
}

func (m movieBrowser) ListMovies() []Movie {
	movList := []Movie{}
	for _, m := range m.movies {
		movList = append(movList, m)
	}

	return movList
}

func (m movieBrowser) GetMovie(name string) Movie {
	mov, found := m.movies[name]
	if !found {
		return Movie{}
	}
	return mov
}

func (m movieBrowser) GetTheatre(name string) Theatre {
	th, found := m.theatres[name]
	if !found {
		return Theatre{}
	}
	return th
}

func NewMovieBrowser() MovieBrowser {
	return &movieBrowser{
		movies: map[string]Movie{
			"movie-1": {
				Id:       1,
				Name:     "movie-1",
				Theatres: []int{1},
			},
			"movie-2": {
				Id:       1,
				Name:     "movie-2",
				Theatres: []int{1},
			},
		},
		theatres: map[string]Theatre{
			"theatre-1": {
				id:     1,
				name:   "theatre-1",
				movies: []int{1, 2},
				halls:  []int{1, 2},
			},
		},
	}
}
