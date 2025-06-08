package models

import (
	"gorm.io/gorm"
)

// Struct que representa a tabela 'emissaos'
type Emissao struct {
	EmpresaID         string  `gorm:"primaryKey" json:"empresa_id"`
	ConsumoEnergia    float64 `json:"consumo_energia"`
	ConsumoTransporte float64 `json:"consumo_transporte"`
}


var DB *gorm.DB

// Criar uma nova emissão
func Criar(dados Emissao) error {
	return DB.Create(&dados).Error
}

// Listar todas as emissões
func ListarTodos() ([]Emissao, error) {
	var lista []Emissao
	err := DB.Find(&lista).Error
	return lista, err
}

// Buscar uma emissão por ID
func BuscarPorID(empresaID string) (Emissao, error) {
	var dados Emissao
	err := DB.First(&dados, "empresa_id = ?", empresaID).Error
	return dados, err
}

// Atualizar dados de emissão
func Atualizar(empresaID string, novosDados Emissao) error {
	var dados Emissao
	if err := DB.First(&dados, "empresa_id = ?", empresaID).Error; err != nil {
		return err
	}
	return DB.Model(&dados).Updates(novosDados).Error
}

// Deletar uma emissão
func Deletar(empresaID string) error {
	return DB.Delete(&Emissao{}, "empresa_id = ?", empresaID).Error
}

// Calcular a emissão de CO2
func CalcularEmissao(dados Emissao) float64 {
	return dados.ConsumoEnergia*0.0005 + dados.ConsumoTransporte*0.0003
}
