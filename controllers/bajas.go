package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/udistrital/movimientos_arka_crud/models"
)

// BajasController
type BajasController struct {
	beego.Controller
}

// URLMapping ...
func (c *BajasController) URLMapping() {
	c.Mapping("Put", c.Put)
}

// Put ...
// @Title Put
// @Description Actualiza el estado de las solicitudes una vez se registra la revision del comite de almacen
// @Param	body	body 	models.TrRevisionBaja	true	"Informacion de la revision"
// @Success 200 {object} []int
// @Failure 404 not found resource
// @router / [put]
func (c *BajasController) Put() {

	var trBaja *models.TrRevisionBaja
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &trBaja); err != nil {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("400")
	}

	if v, err := models.PostRevisionComite(trBaja); err != nil {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("404")
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}
