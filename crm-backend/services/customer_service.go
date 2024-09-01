package services

import (
    "context"
    "crm-backend/models"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var customerCollection *mongo.Collection

func init() {
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("your_mongo_uri"))
    if err != nil {
        panic(err)
    }
    customerCollection = client.Database("crm").Collection("customers")
}

func CreateCustomer(customer *models.Customer) error {
    _, err := customerCollection.InsertOne(context.TODO(), customer)
    return err
}

func GetCustomers() ([]models.Customer, error) {
    var customers []models.Customer
    cursor, err := customerCollection.Find(context.TODO(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())
    for cursor.Next(context.TODO()) {
        var customer models.Customer
        if err := cursor.Decode(&customer); err != nil {
            return nil, err
        }
        customers = append(customers, customer)
    }
    return customers, nil
}

func UpdateCustomer(customerId string, customer *models.Customer) error {
    filter := bson.M{"_id": customerId}
    _, err := customerCollection.UpdateOne(context.TODO(), filter, bson.M{"$set": customer})
    return err
}

func DeleteCustomer(customerId string) error {
    filter := bson.M{"_id": customerId}
    _, err := customerCollection.DeleteOne(context.TODO(), filter)
    return err
}
