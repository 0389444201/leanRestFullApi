package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"leanRestFullApi/models"
	"net/http"
)

type Repository struct {
	DB *gorm.DB
}
type Books struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

func NewConnection() *gorm.DB {
	dsn := "host=localhost user=postgres password=thai dbname=books port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("connect successfully")
	return db
}
func CreateBook(r *Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		books := Books{}
		if err := c.BindJSON(&books); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Request failed"})
			return
		}
		if err := r.DB.Create(&books).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Could not create book"})
			return
		}
		c.JSON(200, gin.H{"message": "Book has been added"})
	}
}
func DeleteBook(r *Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		books := Books{}
		if err := c.BindJSON(&books); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Request failed"})
			return
		}
		if err := r.DB.Delete(&books).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Deleted"})
	}
}
func GetBookByID(r *Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookId := c.Param("book_id")
		booksModels := &[]models.Books{}
		err := r.DB.First(&booksModels, bookId).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to find book"})
			return
		}
		c.JSON(200, booksModels)
	}
}
func GetBooks(r *Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		booksModels := &[]models.Books{}
		if err := r.DB.Find(&booksModels); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get book "})
			return
		}
		c.JSON(200, booksModels)
	}
}
func (r *Repository) SetupRoutes(app *gin.Engine) {
	api := app.Group("/api")
	{
		api.POST("/create_books", CreateBook(r))
		api.DELETE("/delete_books/:id", DeleteBook(r))
		api.GET("/get_books/:id", GetBookByID(r))
		api.GET("/books", GetBooks(r))
	}
}

func main() {
	db := NewConnection()
	r := Repository{
		DB: db,
	}
	app := gin.Default()
	r.SetupRoutes(app)
	app.Run(":8000")
}
