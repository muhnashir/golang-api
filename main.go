package main

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/payment"
	"bwastartup/transaction"
	"bwastartup/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "admin:admin123@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	userService := user.NewService(userRepository)
	authService :=auth.NewService()
	campaignService :=campaign.NewService(campaignRepository)
	paymentService := payment.NewService()
	transactionService :=transaction.NewService(transactionRepository, campaignRepository, paymentService)

	

	userHandler := handler.NewUserHandler(userService,authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	router := gin.Default()
	router.Static("/images","./images")
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatar",authMiddleware(authService, userService), userHandler.UploadAvatar)

	api.GET("campaigns",campaignHandler.GetCampaigns )
	api.GET("campaigns/:id",campaignHandler.GetCampaign )
	api.POST("campaigns",authMiddleware(authService, userService),campaignHandler.CreateCampiagn )
	api.PUT("campaigns/:id",authMiddleware(authService, userService),campaignHandler.UpdateCampign )
	api.POST("campaign-images",authMiddleware(authService, userService),campaignHandler.UploadImage )

	api.GET("campaigns/:id/transactions",authMiddleware(authService, userService),transactionHandler.GetCampaignTransactions )

	api.GET("transactions",authMiddleware(authService, userService),transactionHandler.GetUserTransactions)	
	api.POST("transactions",authMiddleware(authService, userService),transactionHandler.CreateTransaction)
	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service)(gin.HandlerFunc){
	return func(c *gin.Context){	
		authHeader :=c.GetHeader("Authorization")
		if !strings.Contains(authHeader,"Bearer"){
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
	
		tokenString := ""
		arrayToken :=strings.Split(authHeader, " ")
		if len(arrayToken) == 2{
			tokenString = arrayToken[1]
		}
	
		token , err := authService.ValidateToken(tokenString)
		if err !=nil{
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid{
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId:= int(claim["user_id"].(float64))

		user, err:= userService.GetUserById(userId)
		if err !=nil{
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)

	}
}
