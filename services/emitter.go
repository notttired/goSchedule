package services

import (
	"time"

	"github.com/notttired/goSchedule/models"
)

// Emitter emits items at a set interval
type Emitter interface {
	Add(models.Event)
	StartEmitter(PriorityQueue)
	StopEmitter()
}

// Restriction: All items must have same interval
type ChannelEmitter struct {
	channel chan models.Event
	done    chan bool
}

// Makes new channel emitter with buffer
func NewChannelEmitter(buffer int) Emitter {
	channel := make(chan models.Event, buffer)
	done := make(chan bool)
	ce := ChannelEmitter{channel: channel, done: done}
	return &ce
}

func (ce *ChannelEmitter) Add(event models.Event) {
	ce.channel <- event
}

func (ce *ChannelEmitter) StartEmitter(pq PriorityQueue) {
	go func() {
		var event models.Event
		for {
			select {
			case <-ce.done:
				return
			case event = <-ce.channel:
				if event.Job.IsRepeating && event.EndTime.After(time.Now()) {
					pq.InsertTask(updateEvent(&event).Job)
					// Updates and Reinsertion
					updatedEvent := updateTime(&event)
					ce.Add(updatedEvent)
				}
			}
		}
	}()
}

// Prepares event for insertion
func updateEvent(event *models.Event) models.Event {
	event.Job.IsRepeating = false
	return *event
}

// Updates time for reinsertion
func updateTime(event *models.Event) models.Event {
	event.EndTime = event.EndTime.Add(event.Job.RepeatInterval)
	return *event
}

// Stops emitter
func (ce *ChannelEmitter) StopEmitter() {
	ce.done <- true
}
