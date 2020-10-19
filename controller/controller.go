package controller

import (
	"fmt"
	"net/http"

	"github.com/vrindavanstudio/haldiayatra/dblayer"

	"github.com/gin-gonic/gin"
	"github.com/vrindavanstudio/haldiayatra/models"
)

type ControllerInterface interface {
	AddMember(c *gin.Context)
}

type Controller struct {
	client dblayer.Dao
}

func NewController() (ControllerInterface, error) {
	client, err := dblayer.GetConnection()
	if err != nil {
		fmt.Println("Database Connection error ", err)
	}
	return &Controller{
		client: client,
	}, nil
}

//AddMember Extract member info from request and insert in db
func (controller *Controller) AddMember(c *gin.Context) {

	var member models.Member
	err := c.BindJSON(&member)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Recieved Member  ", member)

	res, err := controller.client.CreateMember(member)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, res)
	return

}
