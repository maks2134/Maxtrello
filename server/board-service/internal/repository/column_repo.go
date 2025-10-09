package repository

import (
	"board-service/internal/models"
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ColumnRepository interface {
	CreateColumn(column *models.Column) error
	GetBoardColumns(boardID string) ([]*models.Column, error)
	UpdateColumn(column *models.Column) error
	DeleteColumn(id string) error
	GetColumnByID(id string) (*models.Column, error)
}

type columnRepository struct {
	collection *mongo.Collection
}

func NewColumnRepository(db *mongo.Database) ColumnRepository {
	return &columnRepository{
		collection: db.Collection("columns"),
	}
}

func (r *columnRepository) CreateColumn(column *models.Column) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	column.ID = primitive.NewObjectID()

	_, err := r.collection.InsertOne(ctx, column)
	return err
}

func (r *columnRepository) GetBoardColumns(boardID string) ([]*models.Column, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"board_id": boardID}
	cursor, err := r.collection.Find(ctx, filter, options.Find().SetSort(bson.M{"position": 1}))
	if err != nil {
		return nil, fmt.Errorf("failed to find columns: %w", err)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Printf("failed to close cursor: %v", err)
		}
	}(cursor, ctx)

	var columns []*models.Column
	for cursor.Next(ctx) {
		var column models.Column
		if err := cursor.Decode(&column); err != nil {
			return nil, fmt.Errorf("failed to decode column: %w", err)
		}
		columns = append(columns, &column)
	}

	return columns, nil
}

func (r *columnRepository) UpdateColumn(column *models.Column) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"title":    column.Title,
			"position": column.Position,
		},
	}

	_, err := r.collection.UpdateByID(ctx, column.ID, update)
	return err
}

func (r *columnRepository) DeleteColumn(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid column ID: %w", err)
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}

func (r *columnRepository) GetColumnByID(id string) (*models.Column, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid column ID: %w", err)
	}

	var column models.Column
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&column)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get column: %w", err)
	}

	return &column, nil
}
