package probe

import "time"

type Result struct {
	RelayName string
	RTT       time.Duration
	Err       error
}
