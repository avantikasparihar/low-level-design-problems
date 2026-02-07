package internal

import "sync"

var showMgr *showManager

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
	mut      sync.Mutex
}

func (ms *MovieShow) BlockCapacity(seats [][]int) {
	ms.mut.Lock()
	defer ms.mut.Unlock()

	for _, s := range seats {
		// block seat
		ms.capacity[s[0]][s[1]] = 1
	}
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

type showManager struct {
	showList map[int]*MovieShow
}

func (s showManager) AddMovieShow(req AddMovieShowReq) MovieShow {
	id := len(s.showList) + 1
	show := &MovieShow{
		Id:      id,
		movie:   req.MovieId,
		theatre: req.TheatreId,
		time:    req.Time,
		hall:    req.HallId,
	}
	s.showList[id] = show

	return *show
}

func (s showManager) RemoveMovieShow(id int) {
	//TODO implement me
	panic("implement me")
}

func GetMovieShowManager() MovieShowManager {
	if showMgr != nil {
		return showMgr
	}
	showMgr = &showManager{
		showList: make(map[int]*MovieShow),
	}
	return showMgr
}
