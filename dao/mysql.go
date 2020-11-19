package dao

import "gorm.io/gorm"
import "gorm.io/driver/mysql"
var (
	DB *gorm.DB
)
func InitMysql() (err error){
	dsn:="root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	DB,err =gorm.Open(mysql.Open(dsn),&gorm.Config{}) //这不要 :=
	if err!=nil{
		return err
	}
	return nil
}