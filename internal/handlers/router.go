package handlers

import (
	"github.com/Re1l1x/Teams/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	groupService *services.GroupService,
	personService *services.PersonService,
) *gin.Engine {

	router := gin.Default()

	groupHandler := NewGroupHandler(groupService)
	personHandler := NewPersonHandler(personService)

	groups := router.Group("/groups")
	{
		groups.POST("", groupHandler.Create)
		groups.GET("", groupHandler.GetAll)
		groups.GET("/:id", groupHandler.GetByID)
		groups.PUT("/:id", groupHandler.Update)
		groups.DELETE("/:id", groupHandler.Delete)

		groups.GET("/:id/people", groupHandler.GetPeople)
		groups.GET("/:id/people/all", groupHandler.GetPeopleRecursive)
		groups.GET("/:id/count", groupHandler.CountPeople)
		groups.GET("/:id/count/all", groupHandler.CountPeopleRecursive)
	}

	people := router.Group("/people")
	{
		people.POST("", personHandler.Create)
		people.GET("", personHandler.GetAll)
		people.GET("/:id", personHandler.GetByID)
		people.PUT("/:id", personHandler.Update)
		people.DELETE("/:id", personHandler.Delete)
	}

	return router
}
