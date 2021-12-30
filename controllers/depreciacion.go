package controllers

import (
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
}

// GetCorte ...
// @Title GetCorte
// @Description Retorna lista de elementos disponibles para liquidar depreciacion a una fecha de corte determinada
// @Param	fechaCorte query string true "Fecha de corte en formato YYYY-MM-DD"
// @Success 200 {object} []models.detalle
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
