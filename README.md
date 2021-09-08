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
		Id:           "test:" ,
		KeepTaskList: true,
	}
	defer func() {
		logger.Info(stopwatch.PrettyPrint())
	}()
   _= stopwatch.Start("GetUserCluster")
  // Do some work.

   _=  stopwatch.Stop()
}
```
