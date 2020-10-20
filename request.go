package gopinba

import (
	"time"
)

type Request struct {
	timeStart   time.Time
	memoryUsage uint64
	schema      string
	scriptName  string
	timers      []*Timer
}

func (request *Request) SetSchema(schema string) {

	request.schema = schema
}

func (request *Request) SetScriptName(scriptName string) {

	request.scriptName = scriptName
}

func (request *Request) TimerStart(tags *Tags) *Timer {

	timer := &Timer{
		started:   true,
		timeStart: time.Now(),
		tags:      tags,
	}

	request.timers = append(request.timers, timer)

	return timer
}

func (request *Request) TimerStop(timer *Timer) {

	timer.started = false
	timer.timeEnd = time.Now()
}
