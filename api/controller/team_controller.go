package controller

import (
	"github.com/gin-gonic/gin"
	"team-finder/boot"
	"team-finder/domain"
)

type TeamController struct {
	TeamUsecase domain.TeamUsecase
	Env         *boot.Env
}

func (t *TeamController) GetTeamById(c *gin.Context) {

}

func (t *TeamController) GetTeamsByUser(c *gin.Context) {

}

func (t *TeamController) GetTeamsByEvent(c *gin.Context) {

}

func (t *TeamController) GetAllTeams(c *gin.Context) {

}

func (t *TeamController) RegTeam(c *gin.Context) {

}
