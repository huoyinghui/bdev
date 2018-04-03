package main

import (
	"bdev/models"
	_ "bdev/routers"
	_ "github.com/lib/pq"

	"bdev/config"
	"github.com/astaxie/beego"
)

func init() {
	beego.BConfig.AppName = config.AppConf.AppName
	beego.BConfig.RunMode = config.AppConf.RunMode
	beego.BConfig.CopyRequestBody = config.AppConf.CopyRequestBody
	beego.BConfig.Listen.HTTPPort = config.AppConf.HttpPort
	beego.BConfig.WebConfig.EnableDocs = config.AppConf.EnableDocs
	beego.BConfig.WebConfig.AutoRender = config.AppConf.AutoRender
}

func main() {
	models.InitDB()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
