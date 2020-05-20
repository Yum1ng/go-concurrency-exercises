package pipeline

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestPipeline(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	var numbers []int
	for i := 0; i < 6; i++ {
		numbers = append(numbers, i)
	}
	start := time.Now()
	in := gen(done, numbers...)
	c1 := sq(done, in, "a")
	c2 := sq(done, in, "b")
	res := merge(done, c1, c2)
	for curr := range res {
		fmt.Println(fmt.Sprintf("res : %d", curr))
	}

	elapsed := time.Since(start)
	log.Printf("time took %s", elapsed)
}
