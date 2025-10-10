package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"board-service/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BoardMemberRepository interface {
	AddBoardMember(boardID, userID, role string) error
	GetBoardMembers(boardID string) ([]string, error)
	RemoveBoardMember(boardID, userID string) error
	IsBoardMember(boardID, userID string) (bool, error)
	GetUserRole(boardID, userID string) (string, error)
}

type boardMemberRepository struct {
	collection *mongo.Collection
}

func NewBoardMemberRepository(db *mongo.Database) BoardMemberRepository {
	return &boardMemberRepository{
		collection: db.Collection("boards"),
	}
}

func (r *boardMemberRepository) AddBoardMember(boardID, userID, role string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(boardID)
	if err != nil {
		return fmt.Errorf("invalid board ID: %w", err)
	}

	member := models.BoardMember{
		UserID:   userID,
		Role:     role,
		JoinedAt: time.Now(),
	}

	update := bson.M{
		"$addToSet": bson.M{
			"members": member,
		},
	}

	result, err := r.collection.UpdateByID(ctx, objectID, update)
	if err != nil {
		return fmt.Errorf("failed to add board member: %w", err)
	}

	if result.MatchedCount == 0 {
		return errors.New("board not found")
	}

	return nil
}

func (r *boardMemberRepository) GetBoardMembers(boardID string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(boardID)
	if err != nil {
		return nil, fmt.Errorf("invalid board ID: %w", err)
	}

	var board models.Board
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&board)
	if err != nil {
		return nil, fmt.Errorf("failed to get board: %w", err)
	}

	members := make([]string, len(board.Members))
	for i, member := range board.Members {
		members[i] = member.UserID
	}

	return members, nil
}

func (r *boardMemberRepository) RemoveBoardMember(boardID, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(boardID)
	if err != nil {
		return fmt.Errorf("invalid board ID: %w", err)
	}

	update := bson.M{
		"$pull": bson.M{
			"members": bson.M{"user_id": userID},
		},
	}

	result, err := r.collection.UpdateByID(ctx, objectID, update)
	if err != nil {
		return fmt.Errorf("failed to remove board member: %w", err)
	}

	if result.MatchedCount == 0 {
		return errors.New("board not found")
	}

	return nil
}

func (r *boardMemberRepository) IsBoardMember(boardID, userID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(boardID)
	if err != nil {
		return false, fmt.Errorf("invalid board ID: %w", err)
	}

	filter := bson.M{
		"_id":             objectID,
		"members.user_id": userID,
	}

	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, fmt.Errorf("failed to check board membership: %w", err)
	}

	return count > 0, nil
}

func (r *boardMemberRepository) GetUserRole(boardID, userID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(boardID)
	if err != nil {
		return "", fmt.Errorf("invalid board ID: %w", err)
	}

	var board models.Board
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&board)
	if err != nil {
		return "", fmt.Errorf("failed to get board: %w", err)
	}

	for _, member := range board.Members {
		if member.UserID == userID {
			return member.Role, nil
		}
	}

	return "", errors.New("user is not a member of this board")
}
