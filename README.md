# stopwatch

stopwatch for golang by spring stopwatch

##  Example

```
package main


import (
	"fmt"
	"github.com/jphuang/stopwatch"
	"testing"
	"time"
)

func main() {
stopwatch := &stop_watch.StopWatch{
		Id:           "test:" ,
		KeepTaskList: true,
	}
	defer func() {
		fmt.Println(stopwatch.PrettyPrint())
	}()
	_= stopwatch.Start("GetUserCluster")
	// Do some work.
	time.Sleep(time.Duration(2) * time.Second)
	_= stopwatch.Stop()
}
```

## log

```
StopWatch 'test:': running time (millis) = 2008
-----------------------------------------
ms     %     Task name
-----------------------------------------
2008  100%  GetUserCluster

--- PASS: TestStopWatch (2.01s)
PASS

```