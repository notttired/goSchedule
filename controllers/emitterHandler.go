package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/notttired/goSchedule/models"
	"github.com/notttired/goSchedule/services"
)

type EmitterHandler struct {
	Emitter services.Emitter
	PQ      services.PriorityQueue
}

func (handler EmitterHandler) EmitterAddHandler(c *gin.Context) {
	var newEvent = models.Event{}

	// Binds received JSON to newAlbum
	if err := c.BindJSON(&newEvent); err != nil {
		return
	}

	handler.Emitter.Add(newEvent)
	c.IndentedJSON(http.StatusCreated, newEvent)
}

func (handler EmitterHandler) EmitterStartHandler(c *gin.Context) {
	// Start the emitter
	handler.Emitter.StartEmitter(handler.PQ)
	c.IndentedJSON(http.StatusOK, "Emitter started")
}

func (handler EmitterHandler) EmitterStopHandler(c *gin.Context) {
	// Stop the emitter
	handler.Emitter.StopEmitter()
	c.IndentedJSON(http.StatusOK, "Emitter stopped")
}
