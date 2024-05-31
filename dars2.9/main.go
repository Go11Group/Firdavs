package main

import (
	"fmt"
	"n11/Firdavs/dars2.8/storage/postgres"
	"n11/Firdavs/dars2.9/model"

	_ "gorm.io/driver/postgres"
)

func main() {
	db, err := postgres.ConnectDB
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&model.Users{})

	//db.Create(&Product{Code: "D43", Price: 1654})
	//db.Create(&Product{Code: "D45433", Price: 56})

	var product model.Users
	db.First(&product, 1)
	//db.First(&product, "code = ? and price = ?", "D42", 100) // find product with code D42
	//db.Model(&product).Where("code = ? and price = ?", "D42", 100).Update("Price", 200)

	fmt.Println(product)

	db.Delete(&product, 1)

}
