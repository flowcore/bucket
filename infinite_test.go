package bucket

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_infiniteBucket_Capacity(t *testing.T) {
	bucket := NewInfiniteBucket()
	assert.True(t, math.MaxUint64 == bucket.Capacity())
}

func Test_infiniteBucket_Count(t *testing.T) {
	bucket := NewInfiniteBucket()
	assert.True(t, 0 == bucket.Count())
}

func Test_infiniteBucket_Fill(t *testing.T) {
	bucket := NewInfiniteBucket()

	fillNotFull(t, bucket)
	assert.True(t, 0 == bucket.Count())
	assert.True(t, math.MaxUint64 == bucket.Remaining())

	fillNotFull(t, bucket)
	assert.True(t, 0 == bucket.Count())
	assert.True(t, math.MaxUint64 == bucket.Remaining())
}

func Test_infiniteBucket_Remaining(t *testing.T) {
	bucket := NewInfiniteBucket()
	assert.True(t, math.MaxUint64 == bucket.Remaining())
}

func Test_infiniteBucket_Empty(t *testing.T) {
	bucket := NewInfiniteBucket()
	bucket.Empty()
}
