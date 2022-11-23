package main

import (
	"fmt"

	"github.com/renatozhang/goorm/app"
	"github.com/renatozhang/goorm/model"
	"gorm.io/gorm"
)

var TestDb *gorm.DB

func main() {
	TestDb = app.MySQL.DB()
	// user := model.User{}
	// 获取第一条记录（主键升序）
	// TestDb.Debug().First(&user) // SELECT * FROM `user` ORDER BY `user`.`user_id` LIMIT 1
	// fmt.Printf("user:%#v", user)
	// 获取一条记录，没有指定排序字段
	// TestDb.Debug().Take(&user) // SELECT * FROM `user` LIMIT 1
	// fmt.Printf("user:%#v", user)
	// 获取最后一条记录（主键降序）
	// TestDb.Debug().Last(&user) //SELECT * FROM `user` ORDER BY `user`.`user_id` DESC LIMIT 1
	// fmt.Printf("%#v", user)
	// result := TestDb.Debug().First(&user)
	// result.RowsAffected // 返回找到的记录数
	// result.Error        // returns error or nil

	// 检查 ErrRecordNotFound 错误
	// errors.Is(result.Error, gorm.ErrRecordNotFound)

	// # 检索全部对象
	// var users []*model.User
	// result := TestDb.Find(&users)
	// fmt.Println(result.RowsAffected)
	// fmt.Println(result.Error)

	// # 条件
	// user := model.User{}
	// Get first matched record
	// TestDb.Where("username=?", "admin").First(&user)
	// fmt.Printf("%#v", user)
	// Get all matched records
	// var users []*model.User
	// TestDb.Where("username <> ?", "admin").Find(&users)
	// fmt.Printf("%#v", users)
	// IN
	// TestDb.Where("username IN ?", []string{"admin", "zhangzeng"}).Find(&users)
	// fmt.Printf("%#v", users)

	//LIKE
	// TestDb.Where("username LIKE ?", "%zhang%").Find(&users)
	// fmt.Printf("%#v", users)

	//AND
	// TestDb.Where("username = ? AND age >=?", "admin", "22").Find(&users)

	// TIME
	// TestDb.Where("update_time < ?", time.Now()).Find(&users)
	// fmt.Printf("%#v", users)

	// BETWEEN
	// db.Where("created_time BETWEEN ? AND ?", lastWeek, today).Find(&users)

	// 	var user = User{ID: 10}
	// db.Where("id = ?", 20}.First(&user)
	// SELECT * FROM users WHERE id = 10 and id = 20 ORDER BY id ASC LIMIT 1

	var user model.User
	TestDb.Debug().Where(&model.User{Username: "admin", NickName: "admin"}).First(&user)
	fmt.Printf("%#v", user)

}
