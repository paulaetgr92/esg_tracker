package models

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)



func ConectarComBanco() {
	dsn := "root:senha123@tcp(127.0.0.1:3306)/esg_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar com o banco de dados:", err)
	}

	// Auto migrate da tabela Emissao
	if err := DB.AutoMigrate(&Emissao{}); err != nil {
		log.Fatal("Erro no AutoMigrate:", err)
	}
	fmt.Println("Conectado ao banco com sucesso!")
}
