package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Board struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	OwnerID     string             `json:"owner_id" bson:"owner_id"`
	Members     []BoardMember      `json:"members" bson:"members"`
	Settings    BoardSettings      `json:"settings" bson:"settings"`
}

type BoardMember struct {
	UserID   string    `json:"user_id" bson:"user_id"`
	Role     string    `json:"role" bson:"role"` // основатель, админ, участник
	JoinedAt time.Time `json:"joined_at" bson:"joined_at"`
}

type BoardSettings struct {
	Color    string `json:"color" bson:"color"`
	IsPublic bool   `json:"is_public" bson:"is_public"`
}

type Column struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Position  int                `json:"position" bson:"position"`
	BoardID   string             `json:"board_id" bson:"board_id"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type CreateBoardRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

type CreateColumnRequest struct {
	Title string `json:"title" binding:"required"`
}

type UpdateBoardRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Color       string `json:"color"`
	IsPublic    *bool  `json:"is_public"`
}

type AddMemberRequest struct {
	UserID string `json:"user_id" binding:"required"`
	Role   string `json:"role"` // базово участник
}

type BoardWithDetails struct {
	Board   *Board    `json:"board"`
	Columns []*Column `json:"columns"`
}
