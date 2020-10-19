package models

import (
	"time"
)

//Member model for members to be added in db
type Member struct {
	//ID          primitive.ObjectID `bson:"_id" json:"id"`

	FirstName   string    `bson:"first_name" json:"firstname"`
	LastName    string    `bson:"last_name" json:"lastname"`
	Email       string    `bson:"email" json:"email"`
	Age         int       `bson:"age" json:"age"`
	Phone       int       `bson:"phone" json:"phone"`
	ConnectedTo int       `bson:"connected_to" json:"connected_to"`
	CreatedAt   time.Time `bson:"created_at" json:"createdat"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updatedat"`
}

//MemberWithID Member with ID
type MemberWithID struct {
	ID     string `json:"id"`
	Member Member
}

//For reference
// type Task struct {
//     ID        primitive.ObjectID `bson:"_id"`
//     CreatedAt time.Time          `bson:"created_at"`
//     UpdatedAt time.Time          `bson:"updated_at"`
//     Text      string             `bson:"text"`
//     Completed bool               `bson:"completed"`
// }
// . .

// type Tea struct {
// 	Id       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
// 	Name     string        `json:"name"`
// 	Category string        `json:"category"`
// }
