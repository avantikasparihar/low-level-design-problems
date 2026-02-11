package internal

type Observer interface {
	UpdateVotes(count int)
}

type userObserver struct {
	id int
}

func NewUserObserver(id int) Observer {
	return &userObserver{
		id,
	}
}

func (uo *userObserver) UpdateVotes(count int) {
	// todo: update reputation for user id
}
