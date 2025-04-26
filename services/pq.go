package services

import (
	"time"

	"github.com/notttired/goSchedule/models"
)

// Priority queue allows for interactions
type PriorityQueue interface {
	InsertTask(job models.Job)
	PopHighest() models.Job
	Start()
	Stop()
}

// Restriction: Requires priority <= n
type ChannelQueue struct {
	channels []chan models.Job
	done     chan bool
}

// Creates n new channel queues with buffer
func NewChannelQueue(n, buffer int) PriorityQueue {
	ch := InitChannels(n, buffer)
	return &ChannelQueue{channels: ch, done: make(chan bool)}
}

func (cq *ChannelQueue) InsertTask(job models.Job) {
	InsertTask(cq.channels, job)
}

func (cq *ChannelQueue) PopHighest() models.Job {
	return PopHighest(cq.channels)
}

func (cq *ChannelQueue) Start() {
	go func() {
		for {
			select {
			case <-cq.done:
				return
			default:
				job := PopHighest(cq.channels)
				if job.Run != nil {
					job.Run()
				}
			}
		}
	}()
}

func (cq *ChannelQueue) Stop() {
	cq.done <- true
}

// Channel Specific:

// Initializes n channels each with buffer
// Heap allocation is handles automatically
func InitChannels(n int, buffer int) []chan models.Job {
	channels := make([]chan models.Job, n)
	if n < 1 || buffer < 1 {
		return channels
	}

	for i := 0; i < n; i++ {
		newChan := make(chan models.Job, buffer)
		channels = append(channels, newChan)
	}
	return channels
}

// Gets channel containing highest priority task
func getHighestPriority(channels []chan models.Job) int {
	for i, channel := range channels {
		select {
		case <-channel:
			return i
		default:
		}
	}
	return -1
}

func InsertTask(channels []chan models.Job, job models.Job) {
	channels[job.Priority] <- job
}

func getTask(channels []chan models.Job, priority int) models.Job {
	return <-(channels[priority])
}

// Creates goRoutine to delay reInsertion
func reInsert(channels []chan models.Job, job models.Job) {
	go func() {
		time.Sleep(job.RepeatInterval)
		InsertTask(channels, job)
	}()
}

// Pops highest priority task
func PopHighest(channels []chan models.Job) models.Job {
	job := getTask(channels, getHighestPriority(channels))
	if job.IsRepeating {
		reInsert(channels, job)
	}
	return job
}

// Create set for removal
