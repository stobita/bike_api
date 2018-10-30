package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stobita/bike_api/handler"
	"github.com/stobita/bike_api/middleware"
)

func main() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware)
	r.POST("/signup", handler.SignUp)
	r.POST("/signin", handler.SignIn)
	r.GET("/recruitments", handler.GetRecruitment)
	r.GET("/prefecture", handler.GetPrefecture)
	authorized := r.Group("/", middleware.TokenAuthMiddleware)
	{
		authorized.POST("/recruitment", handler.PostRecruitment)
		authorized.POST("/recruitment/comment", handler.PostRecruitmentComment)
		authorized.GET("/recruitment/:id/comments", handler.GetRecruitmentComments)
		authorized.GET("/user/recruitment/post", handler.GetUserRecruitmentPost)
		// authorized.GET("/user/recruitment/commented", handler.GetUserRecruitmentCommented)
	}
	port := os.Getenv("PORT")
	r.Run(":" + port)
}
