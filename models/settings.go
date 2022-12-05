package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var Esquema string

func init() {
	Esquema = beego.AppConfig.String("PGschemas")
	if Esquema == "" {
		logs.Critical("ERROR: Esquema no definido")
	}
}
