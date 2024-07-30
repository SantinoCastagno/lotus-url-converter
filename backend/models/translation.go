package models

type Translation struct {
	ID          int    `json:"ID"`
	Source      string `json:"Source" binding:"required"`
	Destination string `json:"Destination" binding:"required"`
}
