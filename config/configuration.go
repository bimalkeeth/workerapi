package config

import (
	"encoding/json"
	"io/ioutil"
	rec "workerapi/errorhandle"
)

type IReadConfig interface {
	ReadConfig()
}

type Config struct {
	User       string `json:"user"`
	DbName  string `json:"dbname"`
	Password string `json:"password"`
	Client       string `json:"client"`
	DbAddress string `json:"dbaddress"`
	Generate bool `json:"generate"` 
}

func NewConfig() *Config{
	return &Config{}
}

func(c *Config)ReadConfig(){

	data, err := ioutil.ReadFile("config.json")
	rec.Error("error in reading config file",err)
	config := &Config{}
	err = json.Unmarshal(data, config)
	c.Client=config.Client
	c.DbName=config.DbName
	c.User=config.User;
	c.Password=config.Password
	c.Generate=config.Generate
}