package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type Subject struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type Recommendation struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	UserID    uint `json:"user_id"`
	SubjectID uint `json:"subject_id"`
}

var db *gorm.DB

func initDB() {
	dsn := "user:password@tcp(db:3306)/education?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&User{}, &Subject{}, &Recommendation{})
}

func main() {
	initDB()
	r := gin.Default()

	r.POST("/register", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&user)
		c.JSON(http.StatusOK, user)
	})

	r.GET("/subjects", func(c *gin.Context) {
		var subjects []Subject
		db.Find(&subjects)
		c.JSON(http.StatusOK, subjects)
	})

	r.Run(":8080")
}
