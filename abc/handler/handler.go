package handler

import (
	"abc/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerAdapter struct {
	s service.ServicePort
}

func NewhandlerAdapter(s service.ServicePort) handlerAdapter {
	return handlerAdapter{s: s}
}

func (h handlerAdapter) GetHand(c *gin.Context) {
	data, err := h.s.GetSer()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": data})
}
