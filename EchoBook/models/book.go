package models

// Book model
type Book struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

// Books book map
var Books = map[int]*Book{}
