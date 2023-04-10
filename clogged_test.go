package bucket

import (
	"testing"
)

func Test_cloggedBucket_Capacity(t *testing.T) {
	bucket := NewCloggedBucket(10)
	assertCapacity(t, bucket, 10)
}

func Test_cloggedBucket_Count(t *testing.T) {
	bucket := NewCloggedBucket(10)
	assertCount(t, bucket, 0)
}

func Test_cloggedBucket_Fill(t *testing.T) {
	capacity := uint64(3)
	bucket := NewCloggedBucket(capacity)

	fillNotFull(t, bucket)
	assertBucketState(t, bucket, capacity, 1)

	fillNotFull(t, bucket)
	assertBucketState(t, bucket, capacity, 2)
}

func Test_cloggedBucket_Full(t *testing.T) {
	capacity := uint64(3)
	bucket := NewCloggedBucket(capacity)

	fillNotFull(t, bucket)
	fillNotFull(t, bucket)
	fillNotFull(t, bucket)
	fillFull(t, bucket)
}

func Test_cloggedBucket_Remaining(t *testing.T) {
	bucket := NewCloggedBucket(10)
	assertRemaining(t, bucket, 10)
}

func Test_cloggedBucket_Empty(t *testing.T) {
	bucket := NewCloggedBucket(10)
	fillNotFull(t, bucket)
	fillNotFull(t, bucket)
	assertCount(t, bucket, 2)
	bucket.Empty()
	assertCount(t, bucket, 0)
}
