package model

type Movie struct {
	ID          string `json:"id" grom:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
