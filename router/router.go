package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sferawann/test2/controller"
	"github.com/sferawann/test2/repository"
)

func NewRouter(borrowerRepository repository.BorrowerRepository, authController *controller.AuthController, borrowerController *controller.BorrowerController, lenderController *controller.LenderController, loanProductController *controller.LoanProductController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	// authRouter := router.Group("/auth")
	// authRouter.POST("/register", authController.Register)
	// authRouter.POST("/login", authController.Login)

	borRouter := router.Group("/borrowers")
	borRouter.POST("/", borrowerController.Insert)
	borRouter.GET("/", borrowerController.FindAll)
	borRouter.GET("/:id", borrowerController.FindByID)
	borRouter.DELETE("/:id", borrowerController.Delete)
	borRouter.GET("/username/:username", borrowerController.FindByUsername)
	borRouter.PUT("/:id", borrowerController.Update)
	// borRouter.GET("/", middleware.DeserializeBorrower(borrowerRepository), borrowerController.FindAll)

	// lenRouter := router.Group("/lenders")
	// lenRouter.POST("/", lenderController.Insert)
	// lenRouter.GET("/", lenderController.FindAll)
	// lenRouter.GET("/:id", lenderController.FindByID)
	// lenRouter.DELETE("/:id", lenderController.Delete)
	// lenRouter.GET("/name/:name", lenderController.FindByName)
	// lenRouter.PUT("/:id", lenderController.Update)

	// lpRouter := router.Group("/loan_products")
	// lpRouter.POST("/", loanProductController.Insert)
	// lpRouter.GET("/", loanProductController.FindAll)
	// lpRouter.GET("/:id", loanProductController.FindByID)
	// lpRouter.DELETE("/:id", loanProductController.Delete)
	// lpRouter.GET("/name/:name", loanProductController.FindByName)
	// lpRouter.PUT("/:id", loanProductController.Update)

	return service
}
