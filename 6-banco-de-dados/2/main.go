package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dns := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 1000,
	// })

	// products := []Product{
	// 	{
	// 		Name:  "smartphone",
	// 		Price: 1000,
	// 	},
	// 	{
	// 		Name:  "Mouse",
	// 		Price: 50,
	// 	},
	// 	{
	// 		Name:  "Tablet",
	// 		Price: 3500,
	// 	},
	// }

	// db.Create(&products)

	//select one
	// var product Product

	// db.First(&product, 1)
	// fmt.Println(product)

	// db.First(&product, "name = ?", "smartphone")
	// fmt.Println(product)

	//select all
	// var products []Product
	// db.Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p)
	// }

	//update and delete
	// var p Product
	// db.First(&p, 1)
	// p.Name = "new Mouse"
	// db.Save(&p)

	var p2 Product
	db.First(&p2, 5)
	fmt.Println(p2)
	db.Delete(&p2)

}
