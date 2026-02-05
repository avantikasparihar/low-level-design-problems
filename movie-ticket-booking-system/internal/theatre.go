package internal

type Theatre struct {
	id         int
	name       string
	location   string
	movies     []int
	halls      []int
	movieShows []int
}

type Hall struct {
	id         int
	theatre    int
	location   string
	seatMatrix [][]int
}

type MovieShow struct {
	Id       int
	movie    int
	theatre  int
	time     string
	hall     int
	capacity [][]int
}

type AddMovieShowReq struct {
	MovieId   int
	TheatreId int
	Time      string
	HallId    int
}

type MovieShowManager interface {
	AddMovieShow(req AddMovieShowReq) MovieShow
	RemoveMovieShow(id int)
}

type showMgr struct {
	showList map[int]MovieShow
}

func (s showMgr) AddMovieShow(req AddMovieShowReq) MovieShow {
	id := len(s.showList) + 1
	show := MovieShow{
		Id:      id,
		movie:   req.MovieId,
		theatre: req.TheatreId,
		time:    req.Time,
		hall:    req.HallId,
	}
	s.showList[id] = show

	return show
}

func (s showMgr) RemoveMovieShow(id int) {
	//TODO implement me
	panic("implement me")
}

func NewMovieShowManager() MovieShowManager {
	return &showMgr{
		showList: make(map[int]MovieShow),
	}
}
