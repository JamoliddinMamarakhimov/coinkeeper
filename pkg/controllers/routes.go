package controllers

import (
	"coinkeeper/configs"
	_ "coinkeeper/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", PingPong)

	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", SignUp)
		auth.POST("/sign-in", SignIn)
	}

	apiG := r.Group("/api", checkUserAuthentication)

	UserG := apiG.Group("/users")
	{
		UserG.GET("", GetAllUsers)
		UserG.POST("", CreateUser)
		UserG.GET("/:id", GetUserByID)
		UserG.PUT("/:id", UpdateUser)
		UserG.DELETE("/:id", DeleteUser)
	}

	incomeG := apiG.Group("/incomes")
	{
		incomeG.GET("", GetAllIncomes)
		incomeG.POST("", CreateIncome)
		incomeG.GET("/:id", GetIncomeByID)
		incomeG.PUT("/:id", UpdateIncome)
		incomeG.DELETE("/:id", DeleteIncome)
	}

	outcomeG := apiG.Group("/outcomes")
	{
		outcomeG.GET("", GetAllOutcomes)
		outcomeG.POST("", CreateOutcome)
		outcomeG.GET("/:id", GetOutcomeByID)
		outcomeG.PUT("/:id", UpdateOutcome)
		outcomeG.DELETE("/:id", DeleteOutcome)
	}

	accountG := apiG.Group("/accounts")
	{
		accountG.GET("", GetAllAccounts)
		accountG.POST("", CreateAccount)
		accountG.GET("/:id", GetAccountByID)
		accountG.PUT("/:id", UpdateAccount)
		accountG.DELETE("/:id", DeleteAccount)
	}

	expenseG := apiG.Group("/expenses")
	{
		expenseG.GET("", GetAllExpenses)
		expenseG.POST("", CreateExpense)
		expenseG.GET("/:id", GetExpenseByID)
		expenseG.PUT("/:id", UpdateExpense)
		expenseG.DELETE("/:id", DeleteExpense)
	}
	return r
}

func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
