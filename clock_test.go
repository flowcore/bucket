package bucket

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_realClock_uses_real_system_time(t *testing.T) {
	realNow := time.Now()
	clockNow := realClock{}.Now()
	assert.Less(t, realNow, clockNow)
	assert.Less(t, clockNow.Sub(realNow), time.Second)
}

type fakeClock struct {
	t time.Time
}

func newFakeClock(t time.Time) *fakeClock {
	return &fakeClock{t: t}
}

func (fc *fakeClock) Now() time.Time {
	return fc.t
}

func (fc *fakeClock) Set(t time.Time) {
	fc.t = t
}

func (fc *fakeClock) Add(d time.Duration) {
	fc.t = fc.t.Add(d)
}
