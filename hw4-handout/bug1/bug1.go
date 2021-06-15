package bug1
import (
	"sync"
//	"fmt"
)
// Counter stores a count.
//var mutuxIns sync.Mutex
type Counter struct {
	n int64
}

// Inc increments the count in the Counter.
var mutexIns sync.Mutex
func (c *Counter) Inc() {
	//critical section
	mutexIns.Lock()
	c.n++
	mutexIns.Unlock()
	//end of critical section
}
