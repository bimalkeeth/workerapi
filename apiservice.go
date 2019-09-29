package main

import (
    ec "github.com/labstack/echo/v4"
       "github.com/labstack/echo/v4/middleware"
       "workerapi/api"
       "workerapi/config"
    ma "workerapi/entities"
    rec"workerapi/errorhandle"
)

func main() {

  rec.Panicking()
  createSchema()

  echoServer:=ec.New()
  echoServer.Use(middleware.Logger())
  echoServer.Use(middleware.Recover())

  apiBundle := api.NewApi()
  apiBundle.BundleApi(echoServer)
  echoServer.Logger.Fatal(echoServer.Start(":9000"))
}


func createSchema(){
    conf:=config.NewConfig()
    conf.ReadConfig()
    if conf.Generate {
        sch:= ma.NewSchema()
        err:= sch.CreateSchema()
        rec.Error("error in schema generation ",err)
    }
}