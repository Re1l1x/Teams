package main

import (
	"log"

	"github.com/Re1l1x/Teams/internal/config"
	"github.com/Re1l1x/Teams/internal/database"
	"github.com/Re1l1x/Teams/internal/handlers"
	"github.com/Re1l1x/Teams/internal/repository"
	"github.com/Re1l1x/Teams/internal/services"
)

func main() {

	cfg := config.Load()

	db := database.New(cfg)
	defer db.Close()

	groupRepo := repository.NewGroupRepository(db)
	personRepo := repository.NewPersonRepository(db)

	groupService := services.NewGroupService(groupRepo, personRepo)
	personService := services.NewPersonService(personRepo, groupRepo)

	router := handlers.SetupRouter(groupService, personService)

	log.Printf("Server started on port %s", cfg.ServerPort)

	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal(err)
	}
}
