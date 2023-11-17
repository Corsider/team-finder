package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"team-finder/boot"
	"team-finder/domain"
	"team-finder/internal/utils"
)

type EventController struct {
	EventUsecase domain.EventUsecase
	Env          *boot.Env
}

func (ec *EventController) GetEventById(c *gin.Context) {
	event, err := ec.EventUsecase.GetEventById(utils.First(strconv.Atoi(c.Param("id"))))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, event)
}

func (ec *EventController) GetAll(c *gin.Context) {
	events, err := ec.EventUsecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

func (ec *EventController) RegEvent(c *gin.Context) {
	var request domain.EventRegRequest
	eventId, err := ec.EventUsecase.RegEvent(request, utils.First(strconv.Atoi(c.Param("user_id"))))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.EventRegResponse{EventId: strconv.Itoa(eventId)})
}
