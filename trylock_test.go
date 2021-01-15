package trylock

import (
	"fmt"
	"testing"
	"time"
)

func TestTrylock_Lock(t *testing.T) {
	tl := NewTrylock()
	for i := 0; i < 10; i++ {
		go func() {
			if tl.Lock() {
				fmt.Println("yes")
			} else {
				fmt.Println("No")
			}
		}()
	}
	time.Sleep(time.Second)
	tl.Unlock()
	for i := 0; i < 10; i++ {
		go func() {
			if tl.Lock() {
				fmt.Println("yes")
			} else {
				fmt.Println("No")
			}
		}()
	}
	time.Sleep(time.Second)
}
