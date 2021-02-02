package models

// Book model
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Books book map
var Books = map[int]*Book{}
