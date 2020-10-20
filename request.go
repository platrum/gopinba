package gopinba

import (
	"time"
)

type request struct {
	timeStart   time.Time
	memoryUsage uint64
	schema      string
	scriptName  string
	timers      []*Timer
}

func (request *request) SetSchema(schema string) {

	request.schema = schema
}

func (request *request) SetScriptName(scriptName string) {

	request.scriptName = scriptName
}

func (request *request) TimerStart(tags *Tags) *Timer {

	timer := &Timer{
		started:   true,
		timeStart: time.Now(),
		tags:      tags,
	}

	request.timers = append(request.timers, timer)

	return timer
}

func (request *request) TimerStop(timer *Timer) {

	timer.started = false
	timer.timeEnd = time.Now()
}
