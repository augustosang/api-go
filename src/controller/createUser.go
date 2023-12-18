package controller

import (
	"fmt"
	"log"

	"github.com/augustosang/api-go/src/configuration/validation"
	"github.com/augustosang/api-go/src/model/dtos"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest dtos.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s\n", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
	response := dtos.UserResponse{
		ID:    "teste",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	c.JSON(201, response)
}
