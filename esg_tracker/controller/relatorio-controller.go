package controllers

import (
	"github.com/paulaetgr92/esg_tracker/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GerarRelatorio(c *gin.Context) {
	id := c.Param("id")
	relatorio, err := models.GerarRelatorio(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, relatorio)
}
