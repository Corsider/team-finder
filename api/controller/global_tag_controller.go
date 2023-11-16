package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"team-finder/boot"
	"team-finder/domain"
)

type GlobalTagController struct {
	GlobaTagUsecase domain.GlobalTagUsecase
	Env             *boot.Env
}

func (g *GlobalTagController) GetAll(c *gin.Context) {
	gtags, err := g.GlobaTagUsecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gtags)
}
