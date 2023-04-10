package bucket

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func assertCapacity(t *testing.T, bucket Bucket, expected uint64) {
	assert.Equal(t, expected, bucket.Capacity())
}

func assertCount(t *testing.T, bucket Bucket, expected uint64) {
	assert.Equal(t, expected, bucket.Count())
}

func assertRemaining(t *testing.T, bucket Bucket, expected uint64) {
	assert.Equal(t, expected, bucket.Remaining())
}

func fillNotFull(t *testing.T, bucket Bucket) {
	assert.NoError(t, bucket.Fill())
}

func fillFull(t *testing.T, bucket Bucket) {
	assert.ErrorIs(t, bucket.Fill(), ErrorFull)
}

func assertBucketState(t *testing.T, bucket Bucket, capacity uint64, count uint64) {
	assertCapacity(t, bucket, capacity)
	assertCount(t, bucket, count)
	assertRemaining(t, bucket, capacity-count)
}
