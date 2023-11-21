package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"team-finder/boot"
	"team-finder/domain"
	"team-finder/internal/utils"
)

type TagController struct {
	TagUsecase domain.TagUsecase
	Env        *boot.Env
}

func (t *TagController) GetAll(c *gin.Context) {
	tags, err := t.TagUsecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, tags)
}

func (t *TagController) GetByUserId(c *gin.Context) {
	tags, err := t.TagUsecase.GetByUserId(utils.First(strconv.Atoi(c.Param("id"))))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, tags)
}

func (t *TagController) GetByTeamId(c *gin.Context) {
	tags, err := t.TagUsecase.GetByTeamId(utils.First(strconv.Atoi(c.Param("id"))))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, tags)
}

func (t *TagController) GetByEventId(c *gin.Context) {
	tags, err := t.TagUsecase.GetByEventId(utils.First(strconv.Atoi(c.Param("id"))))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, tags)
}

func (t *TagController) GetByGlobalTagId(c *gin.Context) {
	tags, err := t.TagUsecase.GetByGlobalTagId(utils.First(strconv.Atoi(c.Param("id"))))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, tags)
}

//////////////////// POST ////////////////////

func (t *TagController) PostTagsToUser(c *gin.Context) {
	var request domain.PostTagsRequest
	userId, _ := strconv.Atoi(c.Query("user_id"))
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	err = t.TagUsecase.PostTagsToUser(userId, request.Tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.NormalResponse{Server: "1"})
}

func (t *TagController) PostTagsToTeam(c *gin.Context) {
	var request domain.PostTagsRequest
	teamId, _ := strconv.Atoi(c.Query("team_id"))
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	err = t.TagUsecase.PostTagsToTeam(teamId, request.Tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.NormalResponse{Server: "1"})
}

func (t *TagController) PostTagsToEvent(c *gin.Context) {
	var request domain.PostTagsRequest
	eventId, _ := strconv.Atoi(c.Query("event_id"))
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	err = t.TagUsecase.PostTagsToEvent(eventId, request.Tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.NormalResponse{Server: "1"})
}

func (t *TagController) DeleteTagsFromUser(c *gin.Context) {
	var request domain.PostTagsRequest
	userId, _ := strconv.Atoi(c.Param("id"))
	c.BindJSON(&request)
	err := t.TagUsecase.DeleteTagsFromUser(request, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.NormalResponse{Server: "1"})
}

func (t *TagController) DeleteTagsFromTeam(c *gin.Context) {
	var request domain.PostTagsRequest
	teamId, _ := strconv.Atoi(c.Param("id"))
	c.BindJSON(&request)
	err := t.TagUsecase.DeleteTagsFromTeam(request, teamId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.NormalResponse{Server: "1"})
}

func (t *TagController) DeleteTagsFromEvent(c *gin.Context) {
	var request domain.PostTagsRequest
	eventId, _ := strconv.Atoi(c.Param("id"))
	c.BindJSON(&request)
	err := t.TagUsecase.DeleteTagsFromEvent(request, eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.NormalResponse{Server: "1"})
}
