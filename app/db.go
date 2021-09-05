package main
import (
	"formllm/dao"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDB()(d *dao.Dao,close func() error,err error){
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3309)/test?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		return
	}
	d = &dao.Dao{DB:db}
	close = db.Close
	return
}