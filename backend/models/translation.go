package models

type Translation struct {
	ID           int    `json:"ID"`
	Source       string `json:"Source" binding:"required"`
	Destionation string `jscon:"Destionation" binding:"required"`
}
