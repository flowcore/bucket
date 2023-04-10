package bucket

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewLeakyBucket(t *testing.T) {
	bucket := NewLeakyBucket(uint64(4), time.Minute).(*leakyBucket)

	assert.Equal(t, uint64(4), bucket.capacity)
	assert.Equal(t, time.Minute, bucket.rate)
	assert.Equal(t, time.Second*15, bucket.leakInterval) // 1/4 of a minute is 15 seconds
	assert.NotNil(t, bucket.clock)
}

func Test_leakyBucket_Capacity(t *testing.T) {
	bucket := NewLeakyBucket(10, time.Minute)
	assertCapacity(t, bucket, 10)
}

func Test_leakyBucket_Count(t *testing.T) {
	bucket := NewLeakyBucket(10, time.Minute)
	assertCount(t, bucket, 0)
}

func Test_leakyBucket_Fill(t *testing.T) {
	clock := newFakeClock(time.Now())
	bucket := testLeakyBucket(3, time.Minute, clock)

	fillNotFull(t, bucket)
	assertBucketState(t, bucket, 3, 1)
	fillNotFull(t, bucket)
	assertBucketState(t, bucket, 3, 2)
	fillNotFull(t, bucket)
	assertBucketState(t, bucket, 3, 3)
	fillFull(t, bucket)
	assertBucketState(t, bucket, 3, 3)

	clock.Add(time.Second * 20)
	assertBucketState(t, bucket, 3, 2)
	clock.Add(time.Second * 20)
	assertBucketState(t, bucket, 3, 1)
	clock.Add(time.Second * 20)
	assertBucketState(t, bucket, 3, 0)

	clock.Add(time.Hour)
	assertBucketState(t, bucket, 3, 0)
	fillNotFull(t, bucket)
	assertBucketState(t, bucket, 3, 1)
	fillNotFull(t, bucket)
	assertBucketState(t, bucket, 3, 2)
	fillNotFull(t, bucket)
	assertBucketState(t, bucket, 3, 3)
	fillFull(t, bucket)
	assertBucketState(t, bucket, 3, 3)
}

func Test_leakyBucket_Remaining(t *testing.T) {
	bucket := NewLeakyBucket(10, time.Minute)
	assertRemaining(t, bucket, 10)
}

func Test_leakyBucket_Empty(t *testing.T) {
	clock := newFakeClock(time.Now())
	bucket := testLeakyBucket(2, time.Minute, clock)

	fillNotFull(t, bucket)
	assertBucketState(t, bucket, 2, 1)
	fillNotFull(t, bucket)
	assertBucketState(t, bucket, 2, 2)

	fillFull(t, bucket)
	assertBucketState(t, bucket, 2, 2)

	bucket.Empty()
	assertBucketState(t, bucket, 2, 0)

	fillNotFull(t, bucket)
	assertBucketState(t, bucket, 2, 1)
}

func testLeakyBucket(capacity uint64, rate time.Duration, clock clock) Bucket {
	return &leakyBucket{capacity: capacity, rate: rate, leakInterval: rate / time.Duration(capacity), clock: clock}
}
