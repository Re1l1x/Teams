package tests

import (
	"net/http"

	"github.com/Re1l1x/Teams/internal/config"
	"github.com/Re1l1x/Teams/internal/database"
	"github.com/Re1l1x/Teams/internal/handlers"
	"github.com/Re1l1x/Teams/internal/repository"
	"github.com/Re1l1x/Teams/internal/services"
	"github.com/gin-gonic/gin"
)

func setup() http.Handler {
	gin.SetMode(gin.TestMode)

	cfg := &config.Config{
		DBHost:     "localhost",
		DBPort:     "5433",
		DBUser:     "postgres",
		DBPassword: "postgres",
		DBName:     "people_groups",
		ServerPort: "8080",
	}

	db := database.New(cfg)

	groupRepo := repository.NewGroupRepository(db)
	personRepo := repository.NewPersonRepository(db)

	groupService := services.NewGroupService(groupRepo, personRepo)
	personService := services.NewPersonService(personRepo, groupRepo)

	router := handlers.SetupRouter(groupService, personService)

	return router
}
