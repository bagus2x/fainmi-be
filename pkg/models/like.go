package models

// AddLikeReq -
type AddLikeReq struct {
	LinkID  int `json:"linkID"`
	OwnerID int `json:"ownerID"`
}

// LikeDetailRes -
type LikeDetailRes struct {
	LinkID    int   `json:"linkID"`
	OwnerID   int   `json:"ownerID"`
	LikerID   int   `json:"likerID"`
	CreatedAt int64 `json:"createdAt"`
}
