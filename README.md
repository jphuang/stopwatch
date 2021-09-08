# stopwatch

stopwatch for golang by spring stopwatch

##  Example

```
package main

import (
  "fmt"

  "github.com/jphuang/stopwatch"
)

func main() {
  stopwatch := &stop_watch.StopWatch{
		Id:           "TryToAddTaskPoint:" + strconv.Itoa(int(req.HostId)),
		KeepTaskList: true,
	}
	defer func() {
		logger.Info(stopwatch.PrettyPrint())
	}()
   _= stopwatch.Start("GetUserCluster")
  // Do some work.

   _=  watch.Stop()
}
```
