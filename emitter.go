package main

import (
	"time"
)

type Emitter interface {
	Add(Event)
	StartEmitter(PriorityQueue)
	StopEmitter()
}

type ChannelEmitter struct {
	channel chan Event
	done    chan bool
}

func NewChannelEmitter(buffer int) Emitter {
	channel := make(chan Event, buffer)
	done := make(chan bool)
	ce := ChannelEmitter{channel: channel, done: done}
	return &ce
}

func (ce *ChannelEmitter) Add(event Event) {
	ce.channel <- event
}

func (ce *ChannelEmitter) StartEmitter(pq PriorityQueue) {
	var event Event
	for {
		select {
		case <-ce.done:
			return
		case event = <-ce.channel:
			if event.job.IsRepeating && event.EndTime.After(time.Now()) {
				pq.InsertTask(updateEvent(&event).job)
				// Updates and Reinsertion
				updatedEvent := updateTime(&event)
				ce.Add(updatedEvent)
			}
		}
	}
}

// Prepares event for insertion
func updateEvent(event *Event) Event {
	event.job.IsRepeating = false
	return *event
}

// Updates time for reinsertion
func updateTime(event *Event) Event {
	event.EndTime = event.EndTime.Add(event.job.RepeatInterval)
	return *event
}

// Stops emitter
func (ce *ChannelEmitter) StopEmitter() {
	ce.done <- true
}
