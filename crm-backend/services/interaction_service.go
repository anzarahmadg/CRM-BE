package services

import (
    "context"
    "crm-backend/models"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var interactionCollection *mongo.Collection

func init() {
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("your_mongo_uri"))
    if err != nil {
        panic(err)
    }
    interactionCollection = client.Database("crm").Collection("interactions")
}

func CreateInteraction(interaction *models.Interaction) error {
    _, err := interactionCollection.InsertOne(context.TODO(), interaction)
    return err
}

func GetInteractionsByCustomerId(customerId string) ([]models.Interaction, error) {
    var interactions []models.Interaction
    cursor, err := interactionCollection.Find(context.TODO(), bson.M{"customer_id": customerId})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())
    for cursor.Next(context.TODO()) {
        var interaction models.Interaction
        if err := cursor.Decode(&interaction); err != nil {
            return nil, err
        }
        interactions = append(interactions, interaction)
    }
    return interactions, nil
}

func UpdateInteraction(interactionId string, interaction *models.Interaction) error {
    filter := bson.M{"_id": interactionId}
    _, err := interactionCollection.UpdateOne(context.TODO(), filter, bson.M{"$set": interaction})
    return err
}

func DeleteInteraction(interactionId string) error {
    filter := bson.M{"_id": interactionId}
    _, err := interactionCollection.DeleteOne(context.TODO(), filter)
    return err
}
