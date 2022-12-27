package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/movimientos_arka_crud/models"
	"github.com/udistrital/utils_oas/errorctrl"
)

// BajasController
type DepreciacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *DepreciacionController) URLMapping() {
	c.Mapping("GetCorte", c.GetCorte)
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Post
// @Description Crea las novedades correspondientes a un cierre determinado y actualiza el cierre
// @Param	body	body	models.TransaccionCierre	true	"body for NovedadElemento content"
// @Success 201 {int} models.TransaccionCierre
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *DepreciacionController) Post() {

	var v models.TransaccionCierre
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		logs.Error(err)
		panic(errorctrl.Error(`Post - json.Unmarshal(c.Ctx.Input.RequestBody, &v)`, err, "400"))
	}

	var m models.Movimiento
	if err := models.SubmitCierre(&v, &m); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = m
	} else {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// GetCorte ...
// @Title GetCorte
// @Description Retorna lista de elementos disponibles para liquidar depreciacion a una fecha de corte determinada
// @Param	fechaCorte query string true "Fecha de corte en formato YYYY-MM-DD"
// @Success 200 {object} []models.DepreciacionElemento
// @Failure 404 not found resource
// @router / [get]
func (c *DepreciacionController) GetCorte() {

	var fecha string
	if fecha_ := c.GetString("fechaCorte"); fecha_ == "" {
		err := "Debe especificar una fecha de corte"
		panic(errorctrl.Error(`GetCorte - c.GetString("fechaCorte")`, err, "400"))
	} else {
		fecha = fecha_
	}

	elementos := make([]*models.DepreciacionElemento, 0)
	if err := models.GetCorteDepreciacion(fecha, &elementos); err != nil {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("404")
	} else {
		c.Data["json"] = elementos
	}

	c.ServeJSON()
}
