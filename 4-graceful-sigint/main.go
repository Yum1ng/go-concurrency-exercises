//////////////////////////////////////////////////////////////////////
//
// Given is a mock process which runs indefinitely and blocks the
// program. Right now the only way to stop the program is to send a
// SIGINT (Ctrl-C). Killing a process like that is not graceful, so we
// want to try to gracefully stop the process first.
//
// Change the program to do the following:
//   1. On SIGINT try to gracefully stop the process using
//          `proc.Stop()`
//   2. If SIGINT is called again, just kill the program (last resort)
//

package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// Create a process
	proc := MockProcess{}

	cancel := make(chan struct{}, 1)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		<-sigs
		cancel <- struct{}{}
		signal.Stop(sigs)
		proc.Stop()
		wg.Done()
	}()
	// Run the process (blocking)
	proc.Run(cancel)
	wg.Wait()

	fmt.Println("here")
}
