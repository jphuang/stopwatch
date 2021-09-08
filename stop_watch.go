package stop_watch

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

type StopWatch struct {
	Id              string
	TaskList        []TaskInfo
	StartTimeMillis int64
	Running         bool
	KeepTaskList    bool
	CurrentTaskName string
	LastTaskInfo    TaskInfo
	TaskCount       int
	TotalTimeMillis int64
}

func (s *StopWatch) Start(taskName string) error {
	if s.Running {
		return errors.New("Can't start StopWatch: it's already running")
	} else {
		s.StartTimeMillis = time.Now().UnixNano() / 1e6
		s.Running = true
		s.CurrentTaskName = taskName
	}
	return nil
}
func (s *StopWatch) Stop() error {
	if !s.Running {
		return errors.New("Can't stop StopWatch: it's not running")
	} else {
		lastTime := time.Now().UnixNano()/1e6 - s.StartTimeMillis
		s.TotalTimeMillis += lastTime
		s.LastTaskInfo = TaskInfo{
			TaskName:   s.CurrentTaskName,
			TimeMillis: lastTime,
		}
		if s.KeepTaskList {
			s.TaskList = append(s.TaskList, s.LastTaskInfo)
		}

		s.TaskCount += 1
		s.Running = false
		s.CurrentTaskName = ""
	}
	return nil
}

func (s *StopWatch) ShortSummary() string {
	return "StopWatch '" + s.Id + "': running time (millis) = " + strconv.FormatInt(s.TotalTimeMillis, 10)
}

func (s *StopWatch) PrettyPrint() string {
	var sb strings.Builder
	sb.WriteString(s.ShortSummary())
	sb.WriteString("\n")
	if !s.KeepTaskList {
		sb.WriteString("No task info kept")
	} else {
		sb.WriteString("-----------------------------------------\n")
		sb.WriteString("ms     %     Task name\n")
		sb.WriteString("-----------------------------------------\n")
		for _, info := range s.TaskList {
			sb.WriteString(strconv.FormatInt(info.TimeMillis, 10) + "  ");
			sb.WriteString(strconv.FormatInt(info.TimeMillis * 100  / s.TotalTimeMillis,10) +  "%  ");
			sb.WriteString(info.TaskName + "\n");
		}
	}

	return sb.String()
}
