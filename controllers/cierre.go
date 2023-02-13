package controllers

import (
	"encoding/json"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/movimientos_arka_crud/models"
	"github.com/udistrital/utils_oas/errorctrl"
)

// CierreController
type CierreController struct {
	beego.Controller
}

// URLMapping ...
func (c *CierreController) URLMapping() {
	c.Mapping("GetCorte", c.GetCorte)
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Post
// @Description Crea las novedades correspondientes a un cierre determinado y actualiza el cierre
// @Param	body	body	models.Movimiento	true	"body for NovedadElemento content"
// @Success	201	{object}	models.Movimiento
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *CierreController) Post() {

	var v models.Movimiento
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		logs.Error(err)
		panic(errorctrl.Error(`Post - json.Unmarshal(c.Ctx.Input.RequestBody, &v)`, err, "400"))
	}

	if v.Id == 0 {
		err := "Debe especificar un cierre para ser aprobado"
		logs.Error(err)
		panic(errorctrl.Error(`Post - v.MovimientoId == 0`, err, "400"))
	}

	if err := models.SubmitCierre(&v); err == nil {
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
func (c *CierreController) GetCorte() {

	var fecha string
	if fecha_ := c.GetString("fechaCorte"); fecha_ == "" {
		err := "Debe especificar una fecha de corte"
		panic(errorctrl.Error(`GetCorte - c.GetString("fechaCorte")`, err, "400"))
	} else {
		fecha = fecha_
	}

	_, err := time.Parse("2006-01-02", fecha)
	if err != nil {
		panic(errorctrl.Error(`GetCorte - time.Parse("2006-01-02", fecha)`, err, "400"))
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
