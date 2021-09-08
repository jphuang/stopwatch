package stop_watch

import (
	"fmt"
	"github.com/jphuang/stopwatch"
	"testing"
	"time"
)

func TestStopWatch(t *testing.T) {
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
