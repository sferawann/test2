package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sferawann/test2/controller"
	"github.com/sferawann/test2/middleware"
	"github.com/sferawann/test2/repository"
)

func NewRouter(borrowerRepository repository.BorrowerRepository, authController *controller.AuthController, borrowerController *controller.BorrowerController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	authRouter := router.Group("/auth")
	authRouter.POST("/register", authController.Register)
	authRouter.POST("/login", authController.Login)

	borRouter := router.Group("/borrowers")
	borRouter.POST("/", borrowerController.Insert)
	borRouter.GET("/", middleware.DeserializeBorrower(borrowerRepository), borrowerController.FindAll)

	return service
}
