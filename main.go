// From https://gorm.io/docs/

package main

import (
	"fmt"

	"github.com/kortschak/utter"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func msg(msg string, product *Product) {
	fmt.Printf("%s: %s\n", msg, utter.Sdump(product))
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

	var product Product

	// Read
	if err := db.First(&product, 1).Error; err != nil { // find product with integer primary key
		fmt.Printf("Product with id=1 is not found\n")
	} else {
		msg("Found with id of 1", &product)
	}
	db.First(&product, "code = ?", "D42") // find product with code D42
	msg("Found with code=D42", &product)

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	msg("Updated price to 200", &product)

	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	msg("Updated Price to 200 and Code to F42 method 1", &product)
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	msg("Updated Price to 200 and Code to F42 method 2", &product)

	// Delete - delete product
	db.Delete(&product, 1)
	msg("Deleted id 1", &product)
}
