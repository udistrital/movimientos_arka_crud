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
// @Description creates NovedadElemento and deletes previous NovedadElemento
// @Param	body		body 	models.NovedadElemento	true		"body for NovedadElemento content"
// @Success 201 {int} models.NovedadElemento
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *DepreciacionController) Post() {

	defer errorctrl.ErrorControlController(c.Controller, "DepreciacionController - Unhandled Error!")

	var v models.NovedadElemento
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		logs.Error(err)
		panic(errorctrl.Error(`Post - json.Unmarshal(c.Ctx.Input.RequestBody, &v)`, err, "400"))
	}

	if _, err := models.AddTrNovedadElemento(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
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

	defer errorctrl.ErrorControlController(c.Controller, "DepreciacionController")

	var fecha string
	if fecha_ := c.GetString("fechaCorte"); fecha_ == "" {
		err := "Debe especificar una fecha de corte"
		panic(errorctrl.Error(`GetCorte - c.GetString("fechaCorte")`, err, "400"))
	} else {
		fecha = fecha_
	}

	if v, err := models.GetCorteDepreciacion(fecha); err != nil {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("404")
	} else {
		c.Data["json"] = v
	}

	c.ServeJSON()
}
