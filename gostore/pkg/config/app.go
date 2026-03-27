package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)
func Connect(){
	d,err := gorm.Open("mysql","amisharyan:aryan248@/simplerest?charset=utf8parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	db = d
}
func GetDB() *gorm.DB{
	return db
}