package business

import (
	"github.com/jinzhu/gorm"
	db "workerapi/connection"
	con "workerapi/contracts"
	"workerapi/entities"
	rc "workerapi/errorhandle"
)
import rep "workerapi/dbrepository"


type IBuShifts interface {
  SwapShift(id uint64)(con.ShiftResultBO,error)
}

type BuShifts struct{}

var DB *gorm.DB

//-----------------------------------------------------
//Get database connection
//-----------------------------------------------------
func getDb(){

  dbn:=db.NewDB()
  dbc,err:= dbn.Open()
  rc.Error("error in opening db",err)
  DB=dbc
}

//------------------------------------------------------
//Create new instance
//------------------------------------------------------
func NewBuShift() IBuShifts{
	return BuShifts{}
}
//-------------------------------------------------------
//Swap worker shifts
//-------------------------------------------------------
func(b BuShifts)SwapShift(id uint64)(con.ShiftResultBO,error){
  getDb()
  shift:=rep.NewShift()
  shift.DB=DB
  workers,err:= shift.GetAllWorkerShifts()
  rc.Error("record not found for worker",err)

  defer shift.DB.Close()

  resultData:=con.ShiftResultBO{}

  var otherWorkers []entities.TableShifts
  var worker = entities.TableShifts{}

  for _,item :=range workers {
  	  if item.ID!=id{
		  otherWorkers=append(otherWorkers,item)
	  }else{
	  	worker.ID=item.ID
	  	worker.Worker=item.Worker
	  	worker.StartedAt=item.StartedAt
	  	worker.EndAt=item.EndAt
	  }
  }
  if worker.ID==0{
	  resultData.ShiftExists=false

  }else{
	  resultData.ShiftExists=true
  }

  var result []con.ShiftBO
  if resultData.ShiftExists{
	  for _,item :=range otherWorkers{
		  if (item.StartedAt.Unix()<=worker.StartedAt.Unix() && item.EndAt.Unix() > worker.StartedAt.Unix()) ||
			  (item.StartedAt.Unix()>=worker.StartedAt.Unix() && item.StartedAt.Unix()< worker.EndAt.Unix()){

		  }else{
			  result=append(result,con.ShiftBO{Id:item.ID,Worker:item.Worker,StartedAt:item.StartedAt,EndAt:item.EndAt})
		  }
	  }
  }else{
	  result=[]con.ShiftBO{}
  }
  resultData.Shifts=result

  return resultData,nil
}


