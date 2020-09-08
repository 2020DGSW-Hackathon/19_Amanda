package main

import (
	"Amanda_Server/config"
	"Amanda_Server/controller"
	"Amanda_Server/database"
	"Amanda_Server/library/jwt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type serverMethod interface {
	main()
}

func main() {
	config.InitConfig()
	database.Connect()
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Static("/image", "profileimage")
	e.GET("/showMyInfo", controller.ShowMyInfo, middleware.JWT([]byte("secret")), jwt.VerifyAccessToken)
	e.GET("/showUserInfo", controller.ShowUserInfo)
	e.GET("/showUserStar", controller.ShowUserStar)
	e.POST("/signin", controller.SignIn)
	e.POST("/signup", controller.SignUp)
	e.POST("/updateStar", controller.UpdateStar, middleware.JWT([]byte("secret")), jwt.VerifyAccessToken)
	e.POST("/updateMyInfo", controller.UpdateMyInfo, middleware.JWT([]byte("secret")), jwt.VerifyAccessToken)
	e.POST("/updateProfileImage", controller.UpdateProfileImage, middleware.JWT([]byte("secret")), jwt.VerifyAccessToken)
	e.POST("/createComment", controller.CreateComment, middleware.JWT([]byte("secret")), jwt.VerifyAccessToken)
	e.Logger.Fatal(e.Start(":3000"))
}
