package models

import (
	"errors"
)

type Relatorio struct {
	EmpresaID          string  `json:"empresa_id"`
	ConsumoEnergia     float64 `json:"consumo_energia"`
	ConsumoTransporte  float64 `json:"consumo_transporte"`
	TotalCO2Emissao    float64 `json:"total_co2_emissao"`
}

func GerarRelatorio(empresaID string) (Relatorio, error) {
	dados, err := BuscarPorID(empresaID)
	if err != nil {
		return Relatorio{}, errors.New("empresa não encontrada")
	}

	total := CalcularEmissao(dados)

	return Relatorio{
		EmpresaID:          dados.EmpresaID,
		ConsumoEnergia:     dados.ConsumoEnergia,
		ConsumoTransporte:  dados.ConsumoTransporte,
		TotalCO2Emissao:    total,
	}, nil
}
