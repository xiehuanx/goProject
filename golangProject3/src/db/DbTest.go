package main

import (
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type str_order struct {
	ID        uint `json:"id"`
	Member_id string
	Order_sn  string
	Ceshi     string `gorm:"column:member_username"`
}

func main() {
	var result str_order
	var results []str_order

	dsn := "root:Xiehuan@3@tcp(sh-cdb-4tu4mj0t.sql.tencentcdb.com:63257)/cp?charset=utf8mb4&parseTime=True&loc=Local"
	db, errDb := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if errDb != nil {
		panic("连接数据库失败")
	}
	//db.Raw("select id,member_id, order_sn from str_order limit 10").Scan(&result)
	//查询第一条数据
	db.First(&result)
	jsonByte, _ := json.Marshal(result)
	//查询最后一条数据
	db.Last(&result)
	fmt.Println(string(jsonByte))
	//查询主键为400的数据
	db.First(&result, 400)
	//select * from str_order where receiver_city = 广州市 的数据
	db.Where("receiver_city = ?", "广州市").Find(&results)
	jsonBytes, _ := json.Marshal(results[0])
	fmt.Println("results的长度是：", len(results))
	fmt.Println(string(jsonBytes))

	db.Where("receiver_city = ? and receiver_region = ? ", "广州市", "海珠区").Find(&results)
	fmt.Println("results的长度是：", len(results))
	//按照
	db.Find("receiver_city = ? and receiver_region = ? ", "广州市", "海珠区").Find(&results)
	fmt.Println("results的长度是：", len(results))
	// 按照道理来说，这种方式应该是最符合Java程序员的用法的
	order := str_order{
		Order_sn: "STR416910415678669797",
		Ceshi:    "18834002180",
	}
	db.Debug().Where(order).Not("member_id = ?", 0).Or("order_sn = ?", "xiehuan").Order("order_sn desc").Find(&results)
	fmt.Println("results的长度是：", len(results))
	//db.Create(&order)
	//fmt.Println("主键id是：", order.ID)
	order2 := str_order{
		Order_sn: "STR416910415678669797",
		Ceshi:    "ceshi ceshiceshi",
	}
	result_1 := db.Where(order).Updates(order2)
	// 更新的记录数
	num := result_1.RowsAffected
	// 更新的错误
	error := result_1.Error

	fmt.Println(num, error)
}
