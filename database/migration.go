package database

import (
	"fmt"
	"synapsis-test/models"
	"synapsis-test/pkg/postgree"
)

// automatic migration
func RunMigration() {
	err := postgree.DB.AutoMigrate(&models.Product{}, &models.Category{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
