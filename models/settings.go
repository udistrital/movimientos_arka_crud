package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var esquema string

func init() {
	esquema = beego.AppConfig.String("PGschema")
	if esquema == "" {
		logs.Critical("ERROR: Esquema no definido")
	}
}
