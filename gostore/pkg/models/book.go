package models

import (
	"github.com/jinzhu/gorm"
	"github.com/amisharyan/gostore/pkg/config"
	
)
var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `gorm:"column:name" json:"name"`
	Author string `gorm:"column:author" json:"author"`
	Publication string `gorm:"column:publication" json:"publication"`
}
func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})


}
func(b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllBooks() []Book{
	var books []Book
	db.Find(&books)
	return books
}
func GetBookById(Id int64) (*Book, *gorm.DB){
	var book Book
	db.First(&book, Id)
	return &book,db
}
func DeleteBook(Id int64) *Book{
	var book Book
	db.Where("id=?",Id).Delete(book)
	return &book
}
func UpdateBook(Id int64) *Book{
	var book Book
	db.Where("id=?",Id).First(&book)
	return &book
}