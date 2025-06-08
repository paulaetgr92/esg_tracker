package controllers

import (
	"github.com/paulaetgr92/esg_tracker/models"
	
	"net/http"

	"github.com/gin-gonic/gin"
)

func CriarEmissao(c *gin.Context) {
	var dados models.Emissao
	if err := c.ShouldBindJSON(&dados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.Criar(dados); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar emissão"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Emissão registrada com sucesso"})
}

func ListarEmissoes(c *gin.Context) {
	emissoes, err := models.ListarTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar emissões"})
		return
	}
	c.JSON(http.StatusOK, emissoes)
}

func ObterEmissaoPorID(c *gin.Context) {
	id := c.Param("id")
	dados, err := models.BuscarPorID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Empresa não encontrada"})
		return
	}
	c.JSON(http.StatusOK, dados)
}

func AtualizarEmissao(c *gin.Context) {
	id := c.Param("id")
	var novosDados models.Emissao
	if err := c.ShouldBindJSON(&novosDados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if id != novosDados.EmpresaID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID da URL e do corpo devem ser iguais"})
		return
	}
	if err := models.Atualizar(id, novosDados); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Empresa não encontrada"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Emissão atualizada com sucesso"})
}

func DeletarEmissao(c *gin.Context) {
	id := c.Param("id")
	if err := models.Deletar(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Empresa não encontrada"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Emissão deletada com sucesso"})
}

func CalcularEmissao(c *gin.Context) {
	id := c.Param("id")
	dados, err := models.BuscarPorID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Empresa não encontrada"})
		return
	}
	total := models.CalcularEmissao(dados)
	c.JSON(http.StatusOK, gin.H{
		"empresa_id":   id,
		"emissao_CO2":  total,
	})
}
