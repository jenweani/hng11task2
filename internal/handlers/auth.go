package handlers

import (
	"encoding/json"
	"hng11task2/services"
	"hng11task2/typ"
	"hng11task2/typ/jwt"
	"os"

	"github.com/gin-gonic/gin"
)

type RegisterRequestBody struct{
	FirstName string
	LastName string
	Email string
	Password string
	Phone string
}

type LoginRequestBody struct{
	Email string
	Password string
}


func RegisterUserHandler(c *gin.Context) {
	requestBody := c.Request.Body

	var body RegisterRequestBody
	err := json.NewDecoder(requestBody).Decode(&body)
	if err != nil{
		c.JSON(400, typ.UnSuccessResponse{
			Message: err.Error(),
			Status: "Bad request",
			StatusCode: 400,
		})
	}

	user, err := services.CreateNewUser(body.FirstName, body.LastName, body.Email, body.Password , body.Phone)
	if err != nil{
		c.JSON(400, typ.UnSuccessResponse{
			Message: err.Error(),
			Status: "Bad request",
			StatusCode: 400,
		})
	}

	token, err := jwt.CreateToken(user.ResponseMap(), os.Getenv("JWTSECRET"), 72)
	if err != nil{
		c.JSON(400, typ.UnSuccessResponse{
			Message: err.Error(),
			Status: "Bad request",
			StatusCode: 400,
		})
	}

	c.JSON(201, typ.SuccessResponse{
		Status: "success",
		Message: "Registeration Successful",
		Data: map[string]interface{}{
			"accessToken": token,
			"user": user.ResponseMap(),
		},
	})
}

func LoginHandler(c *gin.Context){
	requestBody := c.Request.Body

	var body LoginRequestBody
	err := json.NewDecoder(requestBody).Decode(&body)
	if err != nil{
		c.JSON(400, typ.UnSuccessResponse{
			Message: err.Error(),
			Status: "Bad request",
			StatusCode: 400,
		})
	}

	user, err := services.GetUsersByEmailAndPassword(body.Email, body.Password)
	if err != nil{
		c.JSON(400, typ.UnSuccessResponse{
			Message: err.Error(),
			Status: "Bad request",
			StatusCode: 400,
		})
	}

	token, err := jwt.CreateToken(user.ResponseMap(), os.Getenv("JWTSECRET"), 72)
	if err != nil{
		c.JSON(400, typ.UnSuccessResponse{
			Message: err.Error(),
			Status: "Bad request",
			StatusCode: 400,
		})
	}

	c.JSON(201, typ.SuccessResponse{
		Status: "success",
		Message: "Registeration Successful",
		Data: map[string]interface{}{
			"accessToken": token,
			"user": user.ResponseMap(),
		},
	})
}

