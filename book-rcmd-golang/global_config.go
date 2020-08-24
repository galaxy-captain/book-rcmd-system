package rcmd

import (
	"runtime"
	"time"
)

var TaskTimeout = 1000 * time.Millisecond
var CurrentTaskNum = runtime.NumCPU()
