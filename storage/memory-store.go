package storage

import (
	"sync"
)

var (
	store = make(map[string]int)
	mu    sync.Mutex
)

// Stores the receipt ID and points in memory.
func SaveReceipt(id string, points int) {
	mu.Lock()
	defer mu.Unlock()
	store[id] = points
}

// Retrieves the points for a given receipt ID.
func GetPoints(id string) (int, bool) {
	mu.Lock()
	defer mu.Unlock()
	points, exists := store[id]
	return points, exists
}
