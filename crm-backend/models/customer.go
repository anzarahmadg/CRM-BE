package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
    ID          primitive.ObjectID `bson:"_id,omitempty"`
    Name        string             `bson:"name" validate:"required"`
    ContactInfo string             `bson:"contact_info" validate:"required"`
    Company     string             `bson:"company"`
    Status      string             `bson:"status"`
    Notes       string             `bson:"notes"`
}
