package app

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type MySQLStruct struct {
	db   *gorm.DB
	once sync.Once
}

type StarRocksStruct struct {
	db   *gorm.DB
	once sync.Once
}

var MySQL MySQLStruct
var StarRocks StarRocksStruct

func (s *MySQLStruct) DB() *gorm.DB {
	s.once.Do(func() {
		var err error
		// fmt.Println(AppConfig.GetString("dsn.youzu-ad-mysql"))
		fmt.Println("root:123456@tcp(127.0.0.1:3306)/mercury?charset=utf8mb4&parseTime=True&loc=Local")

		// s.db, err = gorm.Open(mysql.Open(AppConfig.GetString("dsn.youzu-ad-mysql")), &gorm.Config{})
		s.db, err = gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/mercury?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})

		if err != nil {
			panic(err.Error())
		}
		//sqlDb, _ := s.db.DB()
		//sqlDb.SetConnMaxLifetime(time.Duration(5) * time.Second)
		//sqlDb.SetMaxIdleConns(2)
		//sqlDb.SetMaxIdleConns(5)
		//sqlDb.SetConnMaxIdleTime(time.Duration(5) * time.Second)
	})
	return s.db
}
func (s *StarRocksStruct) DB() *gorm.DB {
	s.once.Do(func() {
		var err error
		// s.db, err = gorm.Open(mysql.Open(AppConfig.GetString("dsn.youzu-star-rocks")), &gorm.Config{})
		s.db, err = gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/mercury?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})

		if err != nil {
			panic(err.Error())
		}
		//sqlDb, _ := s.db.DB()
		//sqlDb.SetConnMaxLifetime(time.Duration(5) * time.Second)
		//sqlDb.SetMaxIdleConns(2)
		//sqlDb.SetMaxIdleConns(5)
		//sqlDb.SetConnMaxIdleTime(time.Duration(5) * time.Second)
	})
	return s.db
}
