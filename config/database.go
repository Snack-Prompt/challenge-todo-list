package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	for {
		DB, err = gorm.Open(postgres.Open("host=postgres-todo-go user=postgres password=postgres dbname=todo-go port=5432 sslmode=disable"), &gorm.Config{})
		if err == nil {
			break
		}
		log.Println("Aguardando conexão com o banco de dados...")
		time.Sleep(2 * time.Second)
	}
	fmt.Println("Conectado ao banco de dados!")
}
