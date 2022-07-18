package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	BookName string             `json:"bookname,omitempty" validate:"required"`
	Author   string             `json:"author,omitempty" validate:"required"`
}
