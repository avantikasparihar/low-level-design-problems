package main

import (
	. "github.com/avantikasparihar/low-level-design-problems/movie-ticket-booking-system/internal"
	"log"
)

/*
Entities:
	Movie:
		- id int
		- name string
		- theatres []int
	Theatre:
		- id int
		- name string
		- location string
		- movies []int
		- halls []int
		- movieShows []int
	Hall:
		- id int
		- theatre int
		- location string
		- seatMatrix [][]int
	MovieShow:
		- id int
		- movie int
		- theatre int
		- time string
		- hall int
		- capacity [][]int
	Reservation:
		- id int
		- movieShow: int
		- quantity: int
		- seats [2][]int

Operations:
	User:
		- showMovie(name)
		- showTheatre(name)
		- bookMovieShow(id, seats) Reservation <handles concurrency>
		- cancelReservation(id)
	Theatre Admin:
		- addMovieShow(movieShowRequest)
		- removeMovieShow(id)

Interfaces:
	MovieBrowser:
		ListMovies()
		GetMovie()
		GetTheatre()
	MovieReservationManager:
		CreateReservation()
		CancelReservation()
	MovieShowManager:
		AddMovieShow()
		RemoveMovieShow()
*/

func main() {
	movieBrowser := NewMovieBrowser()
	showMgr := NewMovieShowManager()
	resMgr := NewReservationManager()

	movList := movieBrowser.ListMovies()
	log.Println(movList)

	mov := movieBrowser.GetMovie("movie-1")
	log.Println(mov)

	th := movieBrowser.GetTheatre("theatre-1")
	log.Println(th)

	movShow := showMgr.AddMovieShow(AddMovieShowReq{
		MovieId:   1,
		TheatreId: 1,
		Time:      "",
		HallId:    1,
	})
	log.Println(movShow)

	res := resMgr.CreateReservation(CreateReservationReq{
		MovieShowId: movShow.Id,
		Seats: [][]int{
			{0, 0},
			{0, 1},
		},
	})
	log.Println(res)
}
