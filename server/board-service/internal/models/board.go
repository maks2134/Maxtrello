package models

type Board struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	OwnerID     string `json:"owner_id"`
}

type Column struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Position int    `json:"position"`
	BoardID  string `json:"board_id"`
}

type BoardMember struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	BoardID string `json:"board_id"`
	Role    string `json:"role"` // гланый, участник
}

type CreateBoardRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type CreateColumnRequest struct {
	Title string `json:"title" binding:"required"`
}

type UpdateBoardRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type AddMemberRequest struct {
	UserID string `json:"user_id" binding:"required"`
	Role   string `json:"role"` // по дефолту участник
}

type BoardWithDetails struct {
	Board   *Board    `json:"board"`
	Columns []*Column `json:"columns"`
	Members []string  `json:"members"`
}
