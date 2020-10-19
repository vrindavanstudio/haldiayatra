package main

import (
	"github.com/vrindavanstudio/haldiayatra/controller"
	//"github.com/vrindavanstudio/haldiayatra/rest"
	// "context"
	// "fmt"
	// "time"

	// "go.mongodb.org/mongo-driver/bson"

	// "github.com/vrindavanstudio/haldiayatra/dblayer"
)

func main() {

	controller.RunAPI(":8080")

	//It is rarely necessary to close a DB could be removed
	//defer dblayer.CloseConnection(dblayer.Client)

	

}
