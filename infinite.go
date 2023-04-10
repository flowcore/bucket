package bucket

import "math"

type infiniteBucket struct {
}

func NewInfiniteBucket() Bucket {
	return &infiniteBucket{}
}

func (_ infiniteBucket) Capacity() uint64 {
	return math.MaxUint64
}

func (_ infiniteBucket) Remaining() uint64 {
	return math.MaxUint64
}

func (_ infiniteBucket) Count() uint64 {
	return 0
}

func (_ infiniteBucket) Fill() error {
	return nil
}

func (_ infiniteBucket) Empty() {
	return
}
