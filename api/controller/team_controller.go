package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"team-finder/boot"
	"team-finder/domain"
	"team-finder/internal/utils"
)

type TeamController struct {
	TeamUsecase domain.TeamUsecase
	Env         *boot.Env
}

func (t *TeamController) GetTeamById(c *gin.Context) {
	team, err := t.TeamUsecase.GetByTeamId(utils.First(strconv.Atoi(c.Param("id"))))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, team)
}

func (t *TeamController) GetTeamsByUser(c *gin.Context) {
	teams, err := t.TeamUsecase.GetByUserId(utils.First(strconv.Atoi(c.Param("id"))))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, teams)
}

func (t *TeamController) GetTeamsByEvent(c *gin.Context) {
	teams, err := t.TeamUsecase.GetByEventId(utils.First(strconv.Atoi(c.Param("id"))))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, teams)
}

func (t *TeamController) GetAllTeams(c *gin.Context) {
	teams, err := t.TeamUsecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, teams)
}

func (t *TeamController) RegTeam(c *gin.Context) {
	var request domain.TeamsRegRequest
	c.BindJSON(&request)
	creatorId := c.Query("user_id")
	creator, _ := strconv.Atoi(creatorId)

	teamId, err := t.TeamUsecase.RegTeam(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
	} else {
		// adding creator to user_team
		err = t.TeamUsecase.AddUserToTeam(creator, teamId)
		if err != nil {
			t.TeamUsecase.DeleteTeamById(teamId) // bad!
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
			return
		}
		c.JSON(http.StatusOK, domain.TeamsRegResponse{TeamId: strconv.Itoa(teamId)})
	}
}
