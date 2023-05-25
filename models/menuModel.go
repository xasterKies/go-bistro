package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Menu struct {
	ID primitive.ObjectID `bson:"_id"`
	Name string `json:"name" validate:"required"`
	Category string `json:"category" validate:"required"`
	Start_Date *time.Time `json:"start_date"`
	End_Date *time.Time `json:"end_date`
	Created_at time.Time ``
	Updated_at time.Time
	Menu_id string
}