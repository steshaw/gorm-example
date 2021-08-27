// From https://gorm.io/docs/

package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1) // find product with integer primary key
	fmt.Printf("Found with id=1: %+v\n", product)
	db.First(&product, "code = ?", "D42") // find product with code D42
	fmt.Printf("Found with code=D42: %+v\n", product)

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	fmt.Printf("Updated price to 200: %+v\n", product)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	fmt.Printf("Updated Price to 200 and Code to F42 method 1: %+v\n", product)
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	fmt.Printf("Updated Price to 200 and Code to F42 method 2: %+v\n", product)

	// Delete - delete product
	db.Delete(&product, 1)
}
