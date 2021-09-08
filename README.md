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

## log

```
2021-09-08 19:15:51.113 [INF]  StopWatch 'TryToAddTaskPoint:10176': running time (millis) = 703
-----------------------------------------
ms     %     Task name
-----------------------------------------
343  48%  GetUserCluster
292  41%  findRunningAngelTask
68  9%  findRunningAngelMultiTask

```