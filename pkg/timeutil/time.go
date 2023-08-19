package timeutil

import "time"

// UnixNano unix nano
func UnixNano(t int64) time.Time {
	sec := t / 1_000_000_000
	nsec := t % 1_000_000_000
	return time.Unix(sec, nsec)
}
