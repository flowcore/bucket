package bucket

import "time"

type clock interface {
	Now() time.Time
}

type realClock struct {
}

func (r realClock) Now() time.Time {
	return time.Now()
}
