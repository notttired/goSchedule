package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/notttired/goSchedule/models"
	"github.com/notttired/goSchedule/services"
)

type PriorityQueueHandler struct {
	PQ services.PriorityQueue
}

func (handler PriorityQueueHandler) PriorityQueueAddHandler(c *gin.Context) {
	var newJob = models.Job{}

	if err := c.BindJSON(&newJob); err != nil {
		return
	}

	handler.PQ.InsertTask(newJob)
	c.IndentedJSON(http.StatusCreated, newJob)
}

func (handler PriorityQueueHandler) PriorityQueueStartHandler(c *gin.Context) {
	// Start the priority queue
	handler.PQ.Start()
	c.IndentedJSON(http.StatusOK, "Priority queue started")
}

func (handler PriorityQueueHandler) PriorityQueueStopHandler(c *gin.Context) {
	// Stop the priority queue
	handler.PQ.Stop()
	c.IndentedJSON(http.StatusOK, "Priority queue stopped")
}
