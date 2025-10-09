package repository

import (
	"context"
	"errors"
	"fmt"
	"time"
	"user-service/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BoardRepository interface {
	CreateBoard(board *models.Board) error
	GetBoardByID(id string) (*models.Board, error)
	GetUserBoards(userID string) ([]*models.Board, error)
	UpdateBoard(board *models.Board) error
	DeleteBoard(id string) error
}

type boardRepository struct {
	collection *mongo.Collection
}

func NewBoardRepository(db *mongo.Database) BoardRepository {
	return &boardRepository{
		collection: db.Collection("boards"),
	}
}

func (r *boardRepository) CreateBoard(board *models.Board) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	board.ID = primitive.NewObjectID()

	board.Members = []models.BoardMember{
		{
			UserID:   board.OwnerID,
			Role:     "owner",
			JoinedAt: time.Now(),
		},
	}

	if board.Settings.Color == "" {
		board.Settings.Color = "#3498db"
	}

	_, err := r.collection.InsertOne(ctx, board)
	return err
}

func (r *boardRepository) GetBoardByID(id string) (*models.Board, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid board ID: %w", err)
	}

	var board models.Board
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&board)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get board: %w", err)
	}

	return &board, nil
}

func (r *boardRepository) GetUserBoards(userID string) ([]*models.Board, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"members.user_id": userID,
	}

	cursor, err := r.collection.Find(ctx, filter, options.Find().SetSort(bson.M{"created_at": -1}))
	if err != nil {
		return nil, fmt.Errorf("failed to find boards: %w", err)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var boards []*models.Board
	for cursor.Next(ctx) {
		var board models.Board
		if err := cursor.Decode(&board); err != nil {
			return nil, fmt.Errorf("failed to decode board: %w", err)
		}
		boards = append(boards, &board)
	}

	return boards, nil
}

func (r *boardRepository) UpdateBoard(board *models.Board) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"title":       board.Title,
			"description": board.Description,
			"settings":    board.Settings,
		},
	}

	_, err := r.collection.UpdateByID(ctx, board.ID, update)
	return err
}

func (r *boardRepository) DeleteBoard(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid board ID: %w", err)
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}
