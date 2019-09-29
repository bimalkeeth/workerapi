package dbrepository

import (
	"github.com/jinzhu/gorm"
	ent "workerapi/entities"
)

type IRepoShift interface {
	GetWorkerById(id uint64)([]ent.TableShifts,error)
	GetWorkersExcludeById(id uint64)([]ent.TableShifts,error)
	GetAllWorkerShifts()([]ent.TableShifts,error)
}

type RepoShift struct {
	DB *gorm.DB
}

func NewShift() *RepoShift {
	return &RepoShift{}
}

//------------------------------------------------------------------
//get worker by shift id
//------------------------------------------------------------------
func(s *RepoShift) GetWorkerById(id uint64)(ent.TableShifts,error){
	found:=ent.TableShifts{}
	s.DB.First(&found)
	return found,nil
}

//-------------------------------------------------------------------
// Get workers excluding given shift id
//-------------------------------------------------------------------
func(s *RepoShift)GetWorkersExcludeById(id uint64)([]ent.TableShifts,error){

	var otherShifts []ent.TableShifts
    s.DB.Where("id <> ?",id).Find(&otherShifts)
	return otherShifts,nil
}
//------------------------------------------------------------------
//Get all shift data
//------------------------------------------------------------------
func(s *RepoShift)GetAllWorkerShifts()([]ent.TableShifts,error){

	var allShifts []ent.TableShifts
	s.DB.Find(&allShifts)
	return allShifts,nil
}