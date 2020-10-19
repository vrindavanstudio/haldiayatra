package dblayer

import (
	"github.com/vrindavanstudio/haldiayatra/models"
)

//Dao Interface for db operations
type Dao interface {
	
	CreateMember(models.Member) (models.MemberWithID, error) 
}