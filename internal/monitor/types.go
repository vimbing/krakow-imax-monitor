package monitor

import (
	http "github.com/vimbing/fhttp"
)

const OPPENHIMER_ID string = "5297s2r"

type Monitor struct {
	Client *http.Client
}

type MovieEntry struct {
	Id          string
	Day         string
	Time        string
	BookingLink string
}
