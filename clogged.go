package bucket

import "sync/atomic"

type cloggedBucket struct {
	capacity uint64
	count    atomic.Uint64
}

func NewCloggedBucket(capacity uint64) Bucket {
	return &cloggedBucket{capacity: capacity}
}

func (cb *cloggedBucket) Capacity() uint64 {
	return cb.capacity
}

func (cb *cloggedBucket) Remaining() uint64 {
	return cb.capacity - cb.count.Load()
}

func (cb *cloggedBucket) Count() uint64 {
	return cb.count.Load()
}

func (cb *cloggedBucket) Fill() error {
	for true {
		count := cb.count.Load()
		newCount := count + 1
		if newCount > cb.capacity {
			return ErrorFull
		}
		if !cb.count.CompareAndSwap(count, newCount) {
			continue
		}
		break
	}
	return nil
}

func (cb *cloggedBucket) Empty() {
	for true {
		count := cb.count.Load()
		if !cb.count.CompareAndSwap(count, 0) {
			continue
		}
		break
	}
}
