package repository

import (
	"board-service/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type BoardQueryRepository interface {
	GetBoardByID(id string) (*models.Board, error)
	GetUserBoards(userID string) ([]*models.Board, error)
	GetBoardWithDetails(boardID string) (*models.BoardWithDetails, error)
}

type boardQueryRepository struct {
	boardsCollection  *mongo.Collection
	columnsCollection *mongo.Collection
}

func (b boardQueryRepository) GetBoardByID(id string) (*models.Board, error) {

	panic("implement me")
}

func (b boardQueryRepository) GetUserBoards(userID string) ([]*models.Board, error) {
	//TODO implement me
	panic("implement me")
}

func (b boardQueryRepository) GetBoardWithDetails(boardID string) (*models.BoardWithDetails, error) {
	//TODO implement me
	panic("implement me")
}

func NewBoardQueryRepository(db *mongo.Database) BoardQueryRepository {
	return &boardQueryRepository{
		boardsCollection:  db.Collection("boards"),
		columnsCollection: db.Collection("columns"),
	}
}
