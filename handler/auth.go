package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stobita/bike_api/lib"
	"github.com/stobita/bike_api/model"
)

// SignUpJSON signup params
type SignUpJSON struct {
	Email                string `json:"email" binding:"required"`
	Name                 string `json:"name" binding:"required"`
	Password             string `json:"password" binding:"required"`
	ConfirmationPassword string `json:"confirmationPassword" binding:"required"`
}

// SignInJSON signin params
type SignInJSON struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// SignUp User signup
func SignUp(c *gin.Context) {
	var json SignUpJSON

	if c.BindJSON(&json) != nil {
		c.JSON(400, lib.ErrorResponse("invalid Params"))
		return
	}

	email := json.Email
	name := json.Name
	password := json.Password
	confirmationPassword := json.ConfirmationPassword

	if password != confirmationPassword {
		c.JSON(400, lib.ErrorResponse("Password Mismatch"))
		return
	}

	encryptedPassword, err := lib.GetEncryptedPassword(password)
	if err != nil || encryptedPassword == "" {
		c.JSON(400, lib.ErrorResponse("encrypt error"))
		return
	}

	err = model.NewUser().Create(name, email, encryptedPassword)

	if err != nil {
		c.JSON(400, lib.ErrorResponse(err.Error()))
		return
	}

	c.AbortWithStatus(200)
}

// SignIn user sign in
func SignIn(c *gin.Context) {
	var json SignInJSON
	if c.BindJSON(&json) != nil {
		c.JSON(400, lib.ErrorResponse("Invalid Params"))
		return
	}

	email := json.Email
	password := json.Password
	user := model.NewUser()

	if !user.FindOneByEmail(email) {
		c.JSON(400, lib.ErrorResponse("User not found"))
		return
	}

	if lib.ComparePassword(password, user.Password) {
		if tokenString, err := lib.GenerateTokenString(user.ID); err == nil {
			c.JSON(200, gin.H{"token": tokenString})
		} else {
			c.JSON(400, lib.ErrorResponse("token generate error"))
		}
	} else {
		c.JSON(400, lib.ErrorResponse("Invalid email or password"))
	}
	return
}
