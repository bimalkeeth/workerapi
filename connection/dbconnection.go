package connection

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	conf "workerapi/config"
	rec "workerapi/errorhandle"
)
import  _ "github.com/go-sql-driver/mysql"

type IDatabase interface {
	Open()(*gorm.DB,error)
}
type DB struct{}

func NewDB() *DB{return &DB{}}

func(db DB)Open()(*gorm.DB, error){

	config:=conf.NewConfig()
	config.ReadConfig()
	Db, err := gorm.Open("mysql",  fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",config.User,config.Password,config.DbName))
	if err!=nil{
		if err.(*mysql.MySQLError).Number==1049{
			if ok:=createDatabaseIfNotExists(*config);ok{
				Db, err = gorm.Open("mysql",  fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",config.User,config.Password,config.DbName))
				return Db,nil
			}else{
				return Db,err
			}
		}
	  fmt.Println(err)
	}
	return Db,nil
}

func createDatabaseIfNotExists(con conf.Config) bool{
	 db, err := sql.Open("mysql",    fmt.Sprintf("%s:%s@/%s",con.User,con.Password,con.DbAddress))
	 defer db.Close()
	 rec.Error("error in opening database",err)
	_,err = db.Exec("CREATE DATABASE "+con.DbName)
	if err==nil{
		return true
	}
	return false
}