package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Interaction struct {
    ID         primitive.ObjectID `bson:"_id,omitempty"`
    CustomerID primitive.ObjectID `bson:"customer_id" validate:"required"`
    Type       string             `bson:"type" validate:"required"`
    Notes      string             `bson:"notes"`
    Scheduled  string             `bson:"scheduled"`
    Resolved   bool               `bson:"resolved"`
}
