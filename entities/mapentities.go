package entities
import (
	"github.com/jinzhu/gorm"
con "workerapi/connection"
hel "workerapi/helpers"
)
import rec "workerapi/errorhandle"


type IEntityGenerator interface {
	CreateSchema() error
}

type EntityGenerator struct{}

func NewSchema() IEntityGenerator{
	return EntityGenerator{}
}

//--------------------------------------------------
// Create schema if the database table not exists
//--------------------------------------------------
func(t EntityGenerator)CreateSchema() error{
	db:=con.NewDB()
	dbase,err:= db.Open()
	rec.Error("database open not successful", err)
	if ok:=MapShiftsTable(dbase);ok{
		generateShiftData(dbase)
	}
	err=dbase.Close()
	rec.Error("error in closing database",err)
	return nil
}

//--------------------------------------------------
// Insert migration data for shift table
//--------------------------------------------------

func generateShiftData(db *gorm.DB ){
   db.Create(&TableShifts{ Worker:"George",StartedAt:hel.StrToDateTime("2019-08-23 00:00:00"),EndAt:hel.StrToDateTime("2019-08-23 00:30:00")})
   db.Create(&TableShifts{ Worker:"Chris",StartedAt:hel.StrToDateTime("2019-08-23 00:00:00"),EndAt:hel.StrToDateTime("2019-08-23 00:30:00")})
   db.Create(&TableShifts{ Worker:"Scott",StartedAt:hel.StrToDateTime("2018-08-29 13:45:00"),EndAt:hel.StrToDateTime("2018-08-30 14:00:00")})
   db.Create(&TableShifts{ Worker:"Alex",StartedAt:hel.StrToDateTime("2018-06-20 09:00:00"),EndAt:hel.StrToDateTime("2018-06-20 11:00:00")})

}