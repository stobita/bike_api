package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stobita/bike_api/model"
)

func GetPrefecture(c *gin.Context) {
	result := model.NewPrefecture().GetAll()
	c.JSON(200, result)
}
