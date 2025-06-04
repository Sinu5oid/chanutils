package chanutils

import "sync"

// Merge merges provided channels into single channel
func Merge[T any](cs ...<-chan T) chan T {
	out := make(chan T, len(cs))
	if len(cs) == 0 {
		close(out)
		return out
	}

	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan T) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
