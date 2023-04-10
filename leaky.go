package bucket

import (
	"math"
	"sync"
	"time"
)

type leakyBucket struct {
	capacity uint64
	rate     time.Duration
	clock    clock

	mutex        sync.Mutex
	empty        time.Time
	leakInterval time.Duration
}

func NewLeakyBucket(capacity uint64, rate time.Duration) Bucket {
	return &leakyBucket{capacity: capacity, rate: rate, leakInterval: rate / time.Duration(capacity), clock: realClock{}}
}

func (b *leakyBucket) Capacity() uint64 {
	return b.capacity
}

func (b *leakyBucket) Remaining() uint64 {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	return b.capacity - b.count()
}

func (b *leakyBucket) Count() uint64 {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	return b.count()
}

func (b *leakyBucket) count() uint64 {
	if b.clock.Now().After(b.empty) {
		return 0
	}
	nsRemaining := b.empty.Sub(b.clock.Now())
	return uint64(math.Ceil(float64(nsRemaining) / float64(b.leakInterval)))
}

func (b *leakyBucket) Fill() error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	count := b.count()

	if count >= b.capacity {
		return ErrorFull
	}

	if b.clock.Now().After(b.empty) {
		b.empty = b.clock.Now()
	}

	b.empty = b.empty.Add(b.leakInterval)
	return nil
}

func (b *leakyBucket) Empty() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.empty = b.clock.Now()
}
