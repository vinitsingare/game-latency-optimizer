package routing

import (
	"time"
)

type RouteScore struct {
	Name string
	RTT  time.Duration
}

func ChooseBest(routes []RouteScore) RouteScore {
	best := routes[0]
	for _, r := range routes {
		if r.RTT < best.RTT {
			best = r
		}
	}
	return best
}
