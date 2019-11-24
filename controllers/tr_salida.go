package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/movimientos_arka_crud/models"
)

// TrSalidaController operations for Tr_salida
type TrSalidaController struct {
	beego.Controller
}

// URLMapping ...
func (c *TrSalidaController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Create
// @Description create Tr_salida
// @Param	body		body 	models.Tr_salida	true		"body for Tr_salida content"
// @Success 201 {object} models.Tr_salida
// @Failure 403 body is empty
// @router / [post]
func (c *TrSalidaController) Post() {
	var v models.TrSalida
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.AddTransaccionSalida(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			logs.Error(err)
			//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = err
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}
