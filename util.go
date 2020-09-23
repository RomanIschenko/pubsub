package pubsub

import (
	"hash/fnv"
)

// a single id length (like client id or user id)
//const IDLength = 36
//const IDSeparator = "-"

func hash(data []byte) int {
	h := fnv.New32a()
	if _, err := h.Write(data); err != nil {
		return 1
	}
	return int(h.Sum32())
}