package dblayer

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"context"
	"fmt"
	"time"
	//"github.com/gin-gonic/gin"
	"github.com/vrindavanstudio/haldiayatra/models"
)




//AddMember adds member object into db
func (client *DBClient) CreateMember(member models.Member) (models.MemberWithID, error) {

	

	collection := client.Database(MongoConfig.MongoDBName).Collection(MongoConfig.MembersCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	member.CreatedAt = time.Now()
	member.UpdatedAt = time.Now()
	response, err := collection.InsertOne(ctx, member)
	if err != nil {
		fmt.Println("ERROR: Insertion Unsuccessfull", err)
		return models.MemberWithID{}, err
	}
	id := response.InsertedID.(primitive.ObjectID)
	var memberWithID models.MemberWithID
	memberWithID.Member = member
	memberWithID.ID = id.String()
	fmt.Println("Successfully Inserted...", id)

	
	return memberWithID, nil
}
