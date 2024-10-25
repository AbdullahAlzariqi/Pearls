// services/content_service.go
package services

import (
	"context"
	"time"

	"github.com/AbdullahAlzariqi/Pearls/db"
	"github.com/AbdullahAlzariqi/Pearls/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var contentCollection *mongo.Collection

func init() {
	contentCollection = db.MongoDB.Collection("content")
}

// CreateContent inserts a new content document
func CreateContent(ctx context.Context, content *models.Content) (*mongo.InsertOneResult, error) {
	content.CreatedAt = time.Now()
	content.UpdatedAt = time.Now()
	return contentCollection.InsertOne(ctx, content)
}

// GetContentByID retrieves a content document by ID
func GetContentByID(ctx context.Context, id string) (*models.Content, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var content models.Content
	filter := bson.M{"_id": objID}
	err = contentCollection.FindOne(ctx, filter).Decode(&content)
	if err != nil {
		return nil, err
	}
	return &content, nil
}

// UpdateContent updates a content document
func UpdateContent(ctx context.Context, id string, update bson.M) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update = bson.M{
		"$set": update,
	}

	_, err = contentCollection.UpdateOne(ctx, filter, update)
	return err
}

// DeleteContent deletes a content document by ID
func DeleteContent(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	_, err = contentCollection.DeleteOne(ctx, filter)
	return err
}
