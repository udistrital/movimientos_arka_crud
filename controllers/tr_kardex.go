package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/udistrital/movimientos_arka_crud/models"
)

// TrkardexController operations for Tr_Kardex
type TrkardexController struct {
	beego.Controller
}

// URLMapping ...
func (c *TrkardexController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Post", c.PostRespuestaSolicitud)
	c.Mapping("Get", c.GetOne)
}

// GetOne ...
// @Title Get One
// @Description get SoporteMovimiento by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Movimiento
// @Failure 404 not found resource
// @router /:id [get]
func (c *TrkardexController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetTransaccionKardex(id)
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

// GetExistencias
// @Title Consulta el detalle de las fichas kardex y el saldo actual.
// @Description get SoporteMovimiento by id
// @Param	ConSaldo	query	bool	true	"Filtra las fichas kardex que tienen existencias"
// @Success 200 {object} models.ElementosMovimiento
// @Failure 404 not found resource
// @router /aperturas [get]
func (c *TrkardexController) GetExistencias() {

	saldo, _ := c.GetBool("ConSaldo")
	v, err := models.GetAllAperturas(saldo)
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
// @Param	body		body 	models.KardexGeneral	true		"body for SalidaGeneral content"
// @Success 201 {object} models.KardexGeneral
// @Failure 403 body is empty
// @router / [post]
func (c *TrkardexController) Post() {
	var v models.KardexGeneral
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		fmt.Println(v);
		if err := models.AddTransaccionKardex(&v); err == nil {
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

// Post ...
// @Title Create
// @Description create SalidaGeneral
// @Param	body		body 	models.KardexGeneral	true		"body for SalidaGeneral content"
// @Success 201 {object} models.KardexGeneral
// @Failure 403 body is empty
// @router /responder_solicitud/ [post]
func (c *TrkardexController) PostRespuestaSolicitud() {
	var v models.KardexGeneral
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		fmt.Println(v);
		if err := models.ResponderSolicitud(&v); err == nil {
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

// Post ...
// @Title Create
// @Description create SalidaGeneral
// @Param	body		body 	models.Movimiento	true		"body for SalidaGeneral content"
// @Success 201 {object} models.Movimiento
// @Failure 403 body is empty
// @router /rechazar_solicitud/ [post]
func (c *TrkardexController) PostRechazarSolicitud() {
	var v models.Movimiento
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		fmt.Println(v);
		if err := models.RechazarSolicitud(&v); err == nil {
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
