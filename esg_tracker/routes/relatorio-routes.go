package routes

import (
	"github.com/paulaetgr92/esg_tracker/controller"


	"github.com/gin-gonic/gin"
)

func RegistrarRotasRelatorio(r *gin.Engine) {
	relatorio := r.Group("/relatorio")
	{
		relatorio.GET("/:id", controllers.GerarRelatorio)
	}
}
