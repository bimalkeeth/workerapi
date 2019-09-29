package entities

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)
//-------------------------------------------
// table structure for shift table
//--------------------------------------------

type TableShifts struct{
	ID  uint64           `gorm:"column:id;type:bigint(19) unsigned auto_increment;not_null;primary_key"`
	Worker string        `gorm:"column:worker"`
	StartedAt time.Time  `gorm:"column:startat;not_null"`
	EndAt time.Time      `gorm:"column:endat;not_null"`
}

func (t TableShifts) TableName() string {
	return "shifts"
}

func (t TableShifts) Validate(db *gorm.DB) {
	if t.StartedAt.IsZero(){

		_ = db.AddError(errors.New("started time should be defined"))
	}
	if t.EndAt.IsZero() {
		_ = db.AddError(errors.New("end time should not be empty"))
	}
}