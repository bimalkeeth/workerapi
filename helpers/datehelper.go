package helpers

import (
	"github.com/araddon/dateparse"
	"time"
rec "workerapi/errorhandle"
)


//--------------------------------------------------
//Helper function convert string to date time
//--------------------------------------------------
func StrToDateTime(str string) time.Time{
	parsedTime,err:=dateparse.ParseStrict(str)
	rec.Error("error parsing time ",err)
	return parsedTime
}