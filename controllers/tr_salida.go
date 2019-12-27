package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

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
	c.Mapping("Get", c.GetOne)
}

// GetOne ...
// @Title Get One
// @Description get SoporteMovimiento by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Movimiento
// @Failure 404 not found resource
// @router /:id [get]
func (c *TrSalidaController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetTransaccionSalida(id)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// Post ...
// @Title Create
// @Description create SalidaGeneral
// @Param	body		body 	models.SalidaGeneral	true		"body for SalidaGeneral content"
// @Success 201 {object} models.SalidaGeneral
// @Failure 403 body is empty
// @router / [post]
func (c *TrSalidaController) Post() {
	var v models.SalidaGeneral
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		fmt.Println(v);
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
