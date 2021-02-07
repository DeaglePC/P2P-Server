package server

import "sync/atomic"

var globalID uint64

func getID() uint64 {
	atomic.AddUint64(&globalID, 1)
	return globalID
}
