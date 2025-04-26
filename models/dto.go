package models

import (
	"time"
)

type Job struct {
	Priority       int
	IsRepeating    bool
	RepeatInterval time.Duration
	Run            func()
}

type Event struct {
	Job     Job
	EndTime time.Time
}
