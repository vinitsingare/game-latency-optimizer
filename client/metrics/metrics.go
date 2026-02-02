package metrics

import (
	"time"
)

type RTTStats struct {
	Samples []time.Duration
	Size    int
}

func NewRTTStats(size int) *RTTStats {
	return &RTTStats{Size: size}
}

func (r *RTTStats) Add(d time.Duration) {
	if len(r.Samples) == r.Size {
		r.Samples = r.Samples[1:]
	}
	r.Samples = append(r.Samples, d)
}

func (r *RTTStats) Avg() time.Duration {
	var sum time.Duration
	for _, d := range r.Samples {
		sum += d
	}
	if len(r.Samples) == 0 {
		return 0
	}
	return sum / time.Duration(len(r.Samples))
}
