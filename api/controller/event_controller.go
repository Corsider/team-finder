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
	err0 := c.BindJSON(&request)
	if err0 != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err0.Error()})
		return
	}
	eventId, err := ec.EventUsecase.RegEvent(request, utils.First(strconv.Atoi(c.Query("user_id"))))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.EventRegResponse{EventId: strconv.Itoa(eventId)})
}

func (ec *EventController) AddTeamToEvent(c *gin.Context) {
	eventId := utils.First(strconv.Atoi(c.Param("event_id")))
	teamId := utils.First(strconv.Atoi(c.Param("team_id")))
	err := ec.EventUsecase.AddTeamToEvent(eventId, teamId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.NormalResponse{Server: "1"})
}

func (ec *EventController) DeleteEvent(c *gin.Context) {
	eventId := utils.First(strconv.Atoi(c.Param("id")))
	err := ec.EventUsecase.DeleteEvent(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.NormalResponse{Server: "1"})
}
