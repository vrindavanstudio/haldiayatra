package controller

import (
	"github.com/gin-gonic/gin"

)

func RunAPIWithController(address string, controller ControllerInterface) error{

	router := gin.Default()
	
	membersGroup := router.Group("/members")
	membersGroup.POST("/add", controller.AddMember)

	return router.Run(address)
}

func RunAPI(address string) error {
	controller, err := NewController()
	if err != nil {
		return err
	}
	return RunAPIWithController(address, controller)
}