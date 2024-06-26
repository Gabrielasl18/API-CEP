package main

import (
	"fmt"
	"log"
	"modules-app/config"
	"modules-app/controllers"
	"modules-app/routes"
	"modules-app/storage"
	"net/http"
)

func main() {
	// Configuração inicial do banco de dados
	c := config.Config{
		MySQLUser:     "root",
		MySQLPassword: "root",
		MySQLHost:     "mysql", // Nome do serviço definido no docker-compose.yml
		MySQLPort:     "3306",
		MySQLDB:       "ceps_db",
	}

	db, err := storage.NewDB(c)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&controllers.CEP{})
	if err != nil {
		log.Fatalf("Erro ao aplicar migrations: %v", err)
	}

	// Inicializar o repository do CEP
	cepRepository := storage.NewCepRepository(db)

	// Criar instância do controlador com o repository do CEP
	controller := controllers.NewController(cepRepository)

	// Configurar rotas
	router := routes.SetupRoutes(controller)

	// Iniciar servidor HTTP
	fmt.Println("Servidor iniciado na porta :8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
