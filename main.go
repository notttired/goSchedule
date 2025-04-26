package main

import (
	"github.com/gin-gonic/gin"
	"github.com/notttired/goSchedule/controllers"
	"github.com/notttired/goSchedule/services"
)

func main() {
	// Initialize the emitter and priority queue
	const emitterBuffer = 10
	const queueBuffer = 10
	const priorities = 10

	// Initialize the emitter and priority queue
	emitter := services.NewChannelEmitter(emitterBuffer)
	pq := services.NewChannelQueue(priorities, queueBuffer)

	// Initialize the handlers
	emitterHandler := controllers.EmitterHandler{Emitter: emitter, PQ: pq}
	pqHandler := controllers.PriorityQueueHandler{PQ: pq}

	router := gin.Default()

	router.POST("/emitter/add", func(c *gin.Context) {
		emitterHandler.EmitterAddHandler(c)
	})

	router.POST("/emitter/start", func(c *gin.Context) {
		emitterHandler.EmitterStartHandler(c)
	})

	router.POST("/emitter/stop", func(c *gin.Context) {
		emitterHandler.EmitterStopHandler(c)
	})

	router.POST("/priorityqueue/add", func(c *gin.Context) {
		pqHandler.PriorityQueueAddHandler(c)
	})

	router.POST("/priorityqueue/start", func(c *gin.Context) {
		pqHandler.PriorityQueueStartHandler(c)
	})

	router.POST("/priorityqueue/stop", func(c *gin.Context) {
		pqHandler.PriorityQueueStopHandler(c)
	})

	// Start the server
	router.Run(":8080")
}
