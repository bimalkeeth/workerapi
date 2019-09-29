package api

import (
	ec "github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	bu "workerapi/business"
	rec "workerapi/errorhandle"
)

type IApi interface {

	BundleApi(echo  *ec.Echo)
}
type Api struct {}

func NewApi() IApi{
	return Api{}
}

//------------------------------------------------
// Bundling Api
//------------------------------------------------
func (a Api)BundleApi(echo  *ec.Echo){

	echo.GET("/",home)
	echo.GET("/shifts/:id/swaps/",getShiftForSwapping)
}

func getShiftForSwapping(context ec.Context) error {
   id:=context.Param("id")
   if len(id)==0 || id==" "{
   	 context.HTML(http.StatusBadRequest,"<strong>provided id is not correct or empty</string>")
   }

   if _, err := strconv.ParseInt(id,10,64); err != nil {
	   context.HTML(http.StatusBadRequest,"<strong>provided id is not a  number</string>")
   }

   val,err:=strconv.ParseUint(id,10,64)
   rec.Error("error in id conversion",err)

   shift:=bu.NewBuShift()
   result,err:= shift.SwapShift(val)
   rec.Error("error in getting swap shifts",err)

   if !result.ShiftExists {
	   context.HTML(http.StatusBadRequest,"<strong>provided shfit id not found</string>")
   }

   return context.JSONPretty(http.StatusOK,result.Shifts," ")
}

func home(context ec.Context) error {
	return context.String(http.StatusOK,"Welcome to Opensimsim test")
}