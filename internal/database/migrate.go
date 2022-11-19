package database

import (
	"github.com/jinzhu/gorm"
	"github.com/melinaco4/companies-manager/internal/company"
)

// MigrateDB - migrates our database and creates our comment table
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&company.Company{}); result.Error != nil {
		return result.Error
	}
	return nil
}
