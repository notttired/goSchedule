package main

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
	job     Job
	EndTime time.Time
}
