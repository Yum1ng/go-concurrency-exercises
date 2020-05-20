//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"sync"
)

func producer(stream Stream) <-chan *Tweet {
	res := make(chan *Tweet, 3)
	go func() {
		defer close(res)
		for {
			tweet, err := stream.Next()
			if err == ErrEOF {
				return
			}
			res <- tweet
		}
	}()
	return res
}

func consumer(wg *sync.WaitGroup, tweets <-chan *Tweet, worker string) {
	go func() {
		defer func() {
			wg.Done()
		}()
		for t := range tweets {
			if t.IsTalkingAboutGo() {
				fmt.Println(fmt.Sprintf("%s \ttweets about golang : %s", t.Username, worker))
			} else {
				fmt.Println(fmt.Sprintf("%s \tdoes not tweet about golang : %s", t.Username, worker))
			}
		}
	}()
}
