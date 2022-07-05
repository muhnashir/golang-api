package main

import (
	"bwastartup/auth"
	"bwastartup/handler"
	"bwastartup/helper"
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
	userService := user.NewService(userRepository)
	authService :=auth.NewService()

	// token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo5fQ.0qlEgibgOJ9dFtv1vHbjZuJ813y-wdpoN-7z8UkpZME")
	// if err !=nil{
	// 	fmt.Println("error")
	// 	fmt.Println("error")
	// }

	// if token.Valid{
	// 	fmt.Println("valid")
	// }else{
	// 	fmt.Println("Invalidddd")
	// }

	
	// userService.SaveAvatar(9, "images/foto.png")

	// input := user.LoginInput{
	// 	Email:    "nashir@transisi.id",
	// 	Password: "12345678",
	// }
	// user, err := userService.Login(input)
	// if err != nil {
	// 	fmt.Println("terjadi kesalahan")
	// 	fmt.Println(err.Error())
	// }
	// fmt.Printf(user.Name)
	// fmt.Printf(user.Email)
	// userByEmail, err := userRepository.FindByEmail("lutfi@gmail.com")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// if userByEmail.ID == 0 {
	// 	fmt.Println("tidak ada")
	// } else {
	// 	fmt.Println(userByEmail.ID)
	// }

	userHandler := handler.NewUserHandler(userService,authService)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatar",authMiddleware(authService, userService), userHandler.UploadAvatar)

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
