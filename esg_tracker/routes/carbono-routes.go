package routes


import (
	"github.com/paulaetgr92/esg_tracker/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/emissoes", controllers.CriarEmissao)
	r.GET("/emissoes", controllers.ListarEmissoes)
	r.GET("/emissoes/:id", controllers.ObterEmissaoPorID)
	r.PUT("/emissoes/:id", controllers.AtualizarEmissao)
	r.DELETE("/emissoes/:id", controllers.DeletarEmissao)

	
	r.GET("/relatorio/:id", controllers.GerarRelatorio)


	return r
}
