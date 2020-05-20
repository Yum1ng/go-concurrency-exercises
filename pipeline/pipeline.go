// Practice go concurrency patterns: https://blog.golang.org/pipelines
package pipeline

import (
	"fmt"
	"sync"
)

func gen(done chan struct{}, nums ...int) <-chan int {
	res := make(chan int)
	go func() {
		defer func() {
			fmt.Println(fmt.Sprintf("closing gen"))
			close(res)
		}()
		for _, num := range nums {
			select {
			case res <- num:
				fmt.Println(fmt.Sprintf("gen %d", num))
			case <-done:
				return
			}
		}

	}()
	return res
}

func sq(done chan struct{}, nums <-chan int, name string) <-chan int {
	res := make(chan int)
	go func() {
		defer func() {
			fmt.Println(fmt.Sprintf("closing worker %s", name))
			close(res)
		}()
		for num := range nums {
			select {
			case res <- num * num:
				fmt.Println(fmt.Sprintf("sq %d, worker %s", num*num, name))
			case <-done:
				return
			}
		}

	}()
	return res
}

func merge(done chan struct{}, cs ...<-chan int) <-chan int {
	res := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(currC <-chan int) {
			defer func() {
				fmt.Println("merge done")
				wg.Done()
			}()
			for num := range currC {
				select {
				case res <- num:
					fmt.Println(fmt.Sprintf("merge %d", num))
				case <-done:
					return
				}
			}
		}(c)
	}
	// without starting a separate goroutine to close the channel, deadlock will occur.
	// Because here is using unbuffered channel which will block _until_ both sender and receiver is ready.
	// The receiver is ready until wait() is ready. And wait() will never be ready.
	go func() {
		wg.Wait()
		close(res)
	}()
	return res
}
