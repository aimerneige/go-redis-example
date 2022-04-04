package db

import "time"

// SleepTime how long to sleep
const SleepTime = time.Second * 2

// data private data
var data string = "default"

// SlowQueryGet fake database operation to emulate slow query
func SlowQueryGet() (string, error) {
	time.Sleep(SleepTime)
	return "This is something you wants to get from database but slow. " + data, nil
}

// DataUpdate fake database operation to update data
func DataUpdate(d string) error {
	data = d
	return nil
}
