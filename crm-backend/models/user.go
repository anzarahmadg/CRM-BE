package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    ID       primitive.ObjectID `bson:"_id,omitempty"`
    Name     string             `bson:"name" validate:"required"`
    Email    string             `bson:"email" validate:"required,email"`
    Password string             `bson:"password" validate:"required"`
    Role     string             `bson:"role" validate:"required"`
}
