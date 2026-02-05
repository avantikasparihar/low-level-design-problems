package internal

type Reservation struct {
	Id        int
	MovieShow int
	Seats     [][]int
}

type CreateReservationReq struct {
	MovieShowId int
	Seats       [][]int
}

type ReservationManager interface {
	CreateReservation(req CreateReservationReq) Reservation
	CancelReservation(id int)
}

type resMgr struct {
	resList map[int]Reservation
}

func (r resMgr) CreateReservation(req CreateReservationReq) Reservation {
	id := len(r.resList) + 1
	res := Reservation{
		Id:        id,
		MovieShow: req.MovieShowId,
		Seats:     req.Seats,
	}
	r.resList[id] = res
	// todo block seats in theatre hall

	return res
}

func (r resMgr) CancelReservation(id int) {
	//TODO implement me
	panic("implement me")
}

func NewReservationManager() ReservationManager {
	return &resMgr{
		resList: make(map[int]Reservation),
	}
}
