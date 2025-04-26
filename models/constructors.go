package models

import (
	"time"
)

func NewJob(priority int, isRepeating bool, repeatInterval time.Duration, run func()) Job {
	return Job{
		Priority:       priority,
		IsRepeating:    isRepeating,
		RepeatInterval: repeatInterval,
		Run:            run,
	}
}

func NewEvent(job Job, endTime time.Time) Event {
	return Event{
		Job:     job,
		EndTime: endTime,
	}
}
