package main

import (
	"time"
)

type PriorityQueue interface {
	InsertTask(job Job)
	PopHighest() Job
}

type ChannelQueue struct {
	channels []chan Job
}

func NewChannelQueue(n, buffer int) PriorityQueue {
	ch := InitChannels(n, buffer)
	return &ChannelQueue{channels: ch}
}

func (cq *ChannelQueue) InsertTask(job Job) {
	InsertTask(cq.channels, job)
}

func (cq *ChannelQueue) PopHighest() Job {
	return PopHighest(cq.channels)
}

// Channel Specific:

// Initializes n channels each with buffer
// Heap allocation is handles automatically
func InitChannels(n int, buffer int) []chan Job {
	channels := make([]chan Job, n)
	if n < 1 || buffer < 1 {
		return channels
	}

	for i := 0; i < n; i++ {
		newChan := make(chan Job, buffer)
		channels = append(channels, newChan)
	}
	return channels
}

// Gets channel containing highest priority task
func getHighestPriority(channels []chan Job) int {
	for i, channel := range channels {
		select {
		case <-channel:
			return i
		default:
		}
	}
	return -1
}

func InsertTask(channels []chan Job, job Job) {
	channels[job.Priority] <- job
}

func getTask(channels []chan Job, priority int) Job {
	return <-(channels[priority])
}

// Creates goRoutine to delay reInsertion
func reInsert(channels []chan Job, job Job) {
	go func() {
		time.Sleep(job.RepeatInterval)
		InsertTask(channels, job)
	}()
}

// Pops highest priority task
func PopHighest(channels []chan Job) Job {
	job := getTask(channels, getHighestPriority(channels))
	if job.IsRepeating {
		reInsert(channels, job)
	}
	return job
}

// Create set for removal
