package entities

import (
	"github.com/jinzhu/gorm"

)


//---------------------------------------------
//Mapping Table shipt structure to MySql Schema
//---------------------------------------------
func MapShiftsTable(db *gorm.DB) bool {
	if !db.HasTable(&TableShifts{}){
		db.CreateTable(&TableShifts{})
		return true
	}
	return false
}

