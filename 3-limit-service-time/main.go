//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import (
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
}

const limitInSeconds = 10

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(chan<- struct{}), u *User) bool {
	if !u.IsPremium && u.TimeUsed >= limitInSeconds {
		return false
	}
	timer := time.NewTimer(time.Duration(limitInSeconds-u.TimeUsed) * time.Second)
	done := make(chan struct{})
	go process(done)
	start := time.Now()
	select {
	case <-done:
		t := time.Now()
		u.TimeUsed += int64(t.Sub(start).Seconds())
		return true
	case <-timer.C:
		t := time.Now()
		u.TimeUsed += int64(t.Sub(start).Seconds())
		return u.IsPremium
	}
}

func main() {
	RunMockServer()
}
