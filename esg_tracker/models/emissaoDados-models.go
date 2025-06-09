package models

import (
	"gorm.io/gorm"
)

type Emissao struct {
	EmpresaID         string  `gorm:"primaryKey" json:"empresa_id"`
	ConsumoEnergia    float64 `json:"consumo_energia"`
	ConsumoTransporte float64 `json:"consumo_transporte"`
}

var DB *gorm.DB


func Criar(dados Emissao) error {
	return DB.Create(&dados).Error
}


func ListarTodos() ([]Emissao, error) {
	var lista []Emissao
	err := DB.Find(&lista).Error
	return lista, err
}


func BuscarPorID(empresaID string) (Emissao, error) {
	var dados Emissao
	err := DB.First(&dados, "empresa_id = ?", empresaID).Error
	return dados, err
}


func Atualizar(empresaID string, novosDados Emissao) error {
	var dados Emissao
	if err := DB.First(&dados, "empresa_id = ?", empresaID).Error; err != nil {
		return err
	}
	return DB.Model(&dados).Updates(novosDados).Error
}

func Deletar(empresaID string) error {
	return DB.Delete(&Emissao{}, "empresa_id = ?", empresaID).Error
}


func CalcularEmissao(dados Emissao) float64 {
	return dados.ConsumoEnergia*0.0005 + dados.ConsumoTransporte*0.0003
}
