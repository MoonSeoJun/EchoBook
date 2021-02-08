package controllers

import (
	"net/http"

	"github.com/MoonSeoJun/EchoBook/models"
	"github.com/labstack/echo/v4"
)

// CreateBook create one book
func CreateBook(c echo.Context) error {
	u := &models.Book{}

	if err := c.Bind(u); err != nil {
		return err
	}

	if u.Title == "" || u.Author == "" || u.Content == "" {
		return c.Render(http.StatusBadRequest, "badrequest.html", nil)
	}

	createbook := models.Book{Title: u.Title, Author: u.Author, Content: u.Content}

	models.DB.Create(&createbook)

	return c.Render(http.StatusCreated, "onepage.html", createbook)
}

// GetBook get one book
func GetBook(c echo.Context) error {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.Render(http.StatusOK, "onepage.html", book)
}

// GetBookToUpdate get to update one book
func GetBookToUpdate(c echo.Context) error {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.Render(http.StatusOK, "updatepage.html", book)
}

// UpdateBook update one book
func UpdateBook(c echo.Context) error {
	var book models.Book
	u := &models.Book{}

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	models.DB.Model(&book).Updates(u)

	return c.Render(http.StatusOK, "onepage.html", book)
}

// DeleteBook delete one book
func DeleteBook(c echo.Context) error {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	models.DB.Delete(&book)

	return c.NoContent(http.StatusNoContent)
}

// GetAllBooks get all books
func GetAllBooks(c echo.Context) error {
	var books []models.Book

	models.DB.Find(&books)

	return c.Render(http.StatusOK, "index.html", books)
}
