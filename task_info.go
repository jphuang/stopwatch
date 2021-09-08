package stop_watch

type TaskInfo struct {
	TaskName   string
	TimeMillis int64
}

func (t TaskInfo) GetTimeSeconds() int64 {
	return t.TimeMillis / 1e6
}
