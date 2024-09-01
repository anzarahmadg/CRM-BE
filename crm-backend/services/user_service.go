package services

import (
    "context"
    "crm-backend/models"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "golang.org/x/crypto/bcrypt"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection

func init() {
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("your_mongo_uri"))
    if err != nil {
        panic(err)
    }
    userCollection = client.Database("crm").Collection("users")
}

func CreateUser(user *models.User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)
    _, err = userCollection.InsertOne(context.TODO(), user)
    return err
}

func GetUsers() ([]models.User, error) {
    var users []models.User
    cursor, err := userCollection.Find(context.TODO(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())
    for cursor.Next(context.TODO()) {
        var user models.User
        if err := cursor.Decode(&user); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}

func UpdateUser(userId string, user *models.User) error {
    filter := bson.M{"_id": userId}
    _, err := userCollection.UpdateOne(context.TODO(), filter, bson.M{"$set": user})
    return err
}

func DeleteUser(userId string) error {
    filter := bson.M{"_id": userId}
    _, err := userCollection.DeleteOne(context.TODO(), filter)
    return err
}
