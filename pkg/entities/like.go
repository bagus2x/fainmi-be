package entities

// Like struct represent like entity
type Like struct {
	LinkID    int   `json:"linkID"`
	LikerID   int   `json:"likerID"`
	CreatedAt int64 `json:"createdAt"`
}
