package bucket

import "errors"

var ErrorFull = errors.New("bucket is full")

type Bucket interface {
	Capacity() uint64
	Remaining() uint64
	Count() uint64
	Fill() error
	Empty()
}
