package health

import "time"

type RelayHealth struct {
	consecutiveFailures int
	maxFailures         int
	lastRTT             time.Duration
}

func New(maxFailures int) *RelayHealth {
	return &RelayHealth{
		maxFailures: maxFailures,
	}
}

func (h *RelayHealth) RecordSuccess(rtt time.Duration) {
	h.consecutiveFailures = 0
	h.lastRTT = rtt
}

func (h *RelayHealth) RecordFailure() {
	h.consecutiveFailures++
}

func (h *RelayHealth) IsHealthy() bool {
	return h.consecutiveFailures < h.maxFailures
}
