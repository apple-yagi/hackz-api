package routes

import (
	"hackz-api/pkg/api/infrastructure/persistence/gorm"
	"hackz-api/pkg/api/interface/controller"
	"hackz-api/pkg/api/usecase"
	"hackz-api/pkg/db/pg"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	conn := pg.Connect

	// CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	// Init Handler
	taskRepository := gorm.NewTaskRepository(conn())
	taskInteractor := usecase.NewTaskUsecase(taskRepository)
	taskController := controller.NewTaskController(taskInteractor)

	// Task API (Public)
	v1 := router.Group("/v1")
	v1.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "hello world") })
	v1.GET("/tasks", func(c *gin.Context) { taskController.Index(c) })
	v1.POST("/tasks", func(c *gin.Context) { taskController.Create(c) })
	v1.GET("/tasks/:id", func(c *gin.Context) { taskController.Show(c) })
	v1.PUT("/tasks/:id", func(c *gin.Context) { taskController.Put(c) })
	v1.DELETE("/tasks/:id", func(c *gin.Context) { taskController.Delete(c) })

	Router = router
}
