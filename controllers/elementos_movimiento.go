package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/udistrital/movimientos_arka_crud/models"
	"github.com/udistrital/utils_oas/errorctrl"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// ElementosMovimientoController operations for ElementosMovimiento
type ElementosMovimientoController struct {
	beego.Controller
}

// URLMapping ...
func (c *ElementosMovimientoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetByFuncionario", c.GetByFuncionario)
}

// Post ...
// @Title Post
// @Description create ElementosMovimiento
// @Param	body		body 	models.ElementosMovimiento	true		"body for ElementosMovimiento content"
// @Success 201 {int} models.ElementosMovimiento
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *ElementosMovimientoController) Post() {
	var v models.ElementosMovimiento
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddElementosMovimiento(&v); err == nil {
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

// GetOne ...
// @Title Get One
// @Description get ElementosMovimiento by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ElementosMovimiento
// @Failure 404 not found resource
// @router /:id [get]
func (c *ElementosMovimientoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetElementosMovimientoById(id)
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

// GetAll ...
// @Title Get All
// @Description get ElementosMovimiento
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ElementosMovimiento
// @Failure 404 not found resource
// @router / [get]
func (c *ElementosMovimientoController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllElementosMovimiento(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		if l == nil {
			l = []interface{}{}
		}
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the ElementosMovimiento
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.ElementosMovimiento	true		"body for ElementosMovimiento content"
// @Success 200 {object} models.ElementosMovimiento
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *ElementosMovimientoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.ElementosMovimiento{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateElementosMovimientoById(&v); err == nil {
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

// Delete ...
// @Title Delete
// @Description delete the ElementosMovimiento
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *ElementosMovimientoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteElementosMovimiento(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}

// GetByFuncionario ...
// @Title Get By Funcionario
// @Description get Elementos by funcionario_id
// @Param	funcionarioId path string true "tercero_id del funcionario a consultar"
// @Success 200 []int
// @Failure 404 not found resource
// @router /funcionario/:funcionarioId [get]
func (c *ElementosMovimientoController) GetByFuncionario() {

	defer errorctrl.ErrorControlController(c.Controller, "ElementosMovimientoController - Unhandled Error!")

	var id int
	if v, err := c.GetInt(":funcionarioId"); err != nil || v <= 0 {
		if err == nil {
			err = errors.New("Se debe especificar un funcionario válido")
		}
		panic(errorctrl.Error("GetByFuncionario - c.GetInt(\":funcionarioId\")", err, "400"))
	} else {
		id = v
	}

	if el, err := models.GetElementosFuncionario(id); err == nil {
		c.Data["json"] = el
	} else {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}

// GetHistorial ...
// @Title Get Historial de un elemento
// @Description Consulta los movimientos que ha tenido un elemento
// @Param	id path string true "id del elemento"
// @Param	final	query 	bool	false	"Indica si se incluye unicamente el ultimo traslado"
// @Success 200 {object} models.Historial
// @Failure 404 not found resource
// @router /historial/:id [get]
func (c *ElementosMovimientoController) GetHistorial() {

	defer errorctrl.ErrorControlController(c.Controller, "ElementosMovimientoController - Unhandled Error!")

	var id int
	if v, err := c.GetInt(":id"); err != nil || v <= 0 {
		if err == nil {
			err = errors.New("Se debe especificar un elemento válido")
		}
		panic(errorctrl.Error("GetHistorial - c.GetInt(\":id\")", err, "400"))
	} else {
		id = v
	}

	var final bool
	if v, err := c.GetBool("final", false); err != nil {
		panic(errorctrl.Error("GetHistorial - c.GetBool(\"final\", false)", err, "400"))
	} else {
		final = v
	}

	if el, err := models.GetHistorialElemento(id, final); err == nil {
		c.Data["json"] = el
	} else {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}
