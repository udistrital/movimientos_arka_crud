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

// MovimientoController operations for Movimiento
type MovimientoController struct {
	beego.Controller
}

// URLMapping ..
func (c *MovimientoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("GetAllTrasladoByTerceroId", c.GetAllTrasladoByTerceroId)
	c.Mapping("GetAllBajasByTerceroId", c.GetAllBajasByTerceroId)
	c.Mapping("GetAllBodegaByTerceroId", c.GetAllBodegaByTerceroId)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ..
// @Title Post
// @Description create Movimiento
// @Param	body		body 	models.Movimiento	true		"body for Movimiento content"
// @Success 201 {int} models.Movimiento
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *MovimientoController) Post() {
	logs.Info("[controller:Post] ===== POST /v1/movimiento INICIO =====")
	logs.Info("[controller:Post] Method=%s URL=%s RemoteAddr=%s", c.Ctx.Request.Method, c.Ctx.Request.URL.String(), c.Ctx.Request.RemoteAddr)
	logs.Info("[controller:Post] Content-Type=%s Body.len=%d", c.Ctx.Request.Header.Get("Content-Type"), len(c.Ctx.Input.RequestBody))
	logs.Info("[controller:Post] Body crudo:\n%s", string(c.Ctx.Input.RequestBody))

	var v models.Movimiento
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		logs.Info("[controller:Post] Unmarshal OK")
		logs.Info("[controller:Post] v.Id=%d (debe ser 0 para insert nuevo)", v.Id)
		logs.Info("[controller:Post] v.Observacion=%q", v.Observacion)
		logs.Info("[controller:Post] v.Detalle=%q", v.Detalle)
		logs.Info("[controller:Post] v.Activo=%v", v.Activo)
		logs.Info("[controller:Post] v.FechaCorte=%v", v.FechaCorte)

		if v.ConsecutivoId != nil {
			logs.Info("[controller:Post] v.ConsecutivoId=%d", *v.ConsecutivoId)
		} else {
			logs.Info("[controller:Post] v.ConsecutivoId=nil")
		}
		if v.Consecutivo != nil {
			logs.Info("[controller:Post] v.Consecutivo=%q", *v.Consecutivo)
		} else {
			logs.Info("[controller:Post] v.Consecutivo=nil")
		}

		if v.FormatoTipoMovimientoId != nil {
			logs.Info("[controller:Post] v.FormatoTipoMovimientoId.Id=%d", v.FormatoTipoMovimientoId.Id)
		} else {
			logs.Error("[controller:Post] v.FormatoTipoMovimientoId=NIL — FK vacío, el insert probablemente fallará o usará 0")
		}

		if v.EstadoMovimientoId != nil {
			logs.Info("[controller:Post] v.EstadoMovimientoId.Id=%d", v.EstadoMovimientoId.Id)
		} else {
			logs.Error("[controller:Post] v.EstadoMovimientoId=NIL — FK vacío, el insert probablemente fallará o usará 0")
		}

		if v.MovimientoPadreId != nil {
			logs.Info("[controller:Post] v.MovimientoPadreId.Id=%d", v.MovimientoPadreId.Id)
		} else {
			logs.Info("[controller:Post] v.MovimientoPadreId=nil")
		}

		reser, _ := json.MarshalIndent(v, "", "  ")
		logs.Info("[controller:Post] Struct completo re-serializado:\n%s", string(reser))

		logs.Info("[controller:Post] Invocando models.AddMovimiento...")
		if id, err := models.AddMovimiento(&v); err == nil {
			logs.Info("[controller:Post] INSERT OK — id retornado por ORM: %d, v.Id después del insert: %d", id, v.Id)

			respJSON, _ := json.MarshalIndent(v, "", "  ")
			logs.Info("[controller:Post] JSON que se retornará al cliente:\n%s", string(respJSON))

			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			logs.Error("[controller:Post] INSERT FALLÓ — error: %v", err)
			logs.Error("[controller:Post] Struct que falló: %+v", v)
			c.Data["system"] = err
			c.Abort("400")
		}
	} else {
		logs.Error("[controller:Post] Unmarshal FALLÓ — error: %v", err)
		logs.Error("[controller:Post] Body que causó el error: %s", string(c.Ctx.Input.RequestBody))

		// Intentar unmarshal genérico para ver qué campos llegan
		var raw map[string]interface{}
		if json.Unmarshal(c.Ctx.Input.RequestBody, &raw) == nil {
			logs.Error("[controller:Post] Campos presentes en el body: %v", keysOf(raw))
		}

		c.Data["system"] = err
		c.Abort("400")
	}

	logs.Info("[controller:Post] ===== POST /v1/movimiento FIN =====")
	c.ServeJSON()
}

func keysOf(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// GetOne ..
// @Title Get One
// @Description get Movimiento by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Movimiento
// @Failure 404 not found resource
// @router /:id [get]
func (c *MovimientoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetMovimientoById(id)
	if err != nil {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("404")
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ..
// @Title Get All
// @Description get Movimiento
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Movimiento
// @Failure 404 not found resource
// @router / [get]
func (c *MovimientoController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
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

	l, count, err := models.GetAllMovimiento(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("404")
	} else {
		if l == nil {
			l = []interface{}{}
		}
		c.Ctx.Output.Header("total-count", strconv.Itoa(int(count)))
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ..
// @Title Put
// @Description update the Movimiento
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Movimiento	true		"body for Movimiento content"
// @Success 200 {object} models.Movimiento
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *MovimientoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Movimiento{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateMovimientoById(&v); err == nil {
			c.Data["json"] = v
		} else {
			logs.Error(err)
			c.Data["system"] = err
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// Delete ..
// @Title Delete
// @Description delete the Movimiento
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *MovimientoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteMovimiento(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}

// GetEntradaByActa ..
// @Title Get By Acta
// @Description get Movimiento by acta_recibido_id
// @Param	acta_recibido_id path string true "id del acta asociada a la entrada"
// @Success 200 {object} models.Movimiento
// @Failure 404 not found resource
// @router /entrada/:acta_recibido_id [get]
func (c *MovimientoController) GetMovimientoByActa() {
	ActaRecibidoIdStr := c.Ctx.Input.Param(":acta_recibido_id")
	ActaRecibidoId, _ := strconv.Atoi(ActaRecibidoIdStr)
	v, err := models.GetEntradaByActa(ActaRecibidoId)
	if err != nil {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("404")
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAllTrasladoByTerceroId ..
// @Title Get Traslados By Tercero
// @Description Consulta traslados asociados a un tercero determinado.
// @Param	tercero_id	path	string	true	"TerceroId de quien solicita los traslados"
// @Param	confirmar	query	bool	false	"Consulta los traslados que están pendientes de ser confirmados por el tercero que consulta."
// @Success 200 {object} []models.Movimiento
// @Failure 404 not found resource
// @router /traslado/:tercero_id [get]
func (c *MovimientoController) GetAllTrasladoByTerceroId() {

	var (
		terceroId int
		recibir   bool
	)

	if v, err := c.GetInt(":tercero_id", 0); err != nil {
		panic(errorctrl.Error(`GetAll - c.GetInt(":tercero_id", 0)`, err, "400"))
	} else if v > 0 {
		terceroId = v
	} else {
		panic(errorctrl.Error(`GetAll - Se debe especificar un tercero válido`, err, "400"))
	}

	if v, err := c.GetBool("confirmar", false); err != nil {
		panic(errorctrl.Error(`GetAll - c.GetBool("confirmar", false)`, err, "400"))
	} else {
		recibir = v
	}

	var traslados = make([]models.Movimiento, 0)
	err := models.GetTrasladosByTerceroId(terceroId, recibir, &traslados)
	if err != nil {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("404")
	} else {
		c.Data["json"] = traslados
	}
	c.ServeJSON()
}

// GetAllBajasByTerceroId ..
// @Title Get Bajas Solicitadas By Tercero
// @Description Consulta las bajas solicitadas por un tercero determinado.
// @Param	tercero_id	path	string	true	"TerceroId de quien consulta la lista de bajas"
// @Success 200 {object} []models.Movimiento
// @Failure 404 not found resource
// @router /baja/:tercero_id [get]
func (c *MovimientoController) GetAllBajasByTerceroId() {

	var terceroId int

	if v, err := c.GetInt(":tercero_id", 0); err != nil {
		panic(errorctrl.Error(`GetAll - c.GetInt(":tercero_id", 0)`, err, "400"))
	} else if v > 0 {
		terceroId = v
	} else {
		panic(errorctrl.Error(`GetAll - Se debe especificar un tercero válido`, err, "400"))
	}

	var bajas = make([]interface{}, 0)
	if err := models.GetBajasByTerceroId(terceroId, &bajas); err != nil {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("404")
	} else {
		c.Data["json"] = bajas
	}

	c.ServeJSON()
}

// GetAllBodegaByTerceroId ..
// @Title Get Solicitudes de bodega By Tercero
// @Description Consulta las solicitudes de bodega de consumo solicitadas por un tercero determinado.
// @Param	tercero_id	path	string	true	"TerceroId de quien consulta la lista de solicitudes"
// @Success 200 {object} []models.Movimiento
// @Failure 404 not found resource
// @router /bodega/:tercero_id [get]
func (c *MovimientoController) GetAllBodegaByTerceroId() {

	var terceroId int

	if v, err := c.GetInt(":tercero_id", 0); err != nil {
		panic(errorctrl.Error(`GetAllBodegaByTerceroId - c.GetInt(":tercero_id", 0)`, err, "400"))
	} else if v > 0 {
		terceroId = v
	} else {
		panic(errorctrl.Error(`GetAllBodegaByTerceroId - Se debe especificar un tercero válido`, err, "400"))
	}

	var solicitudes = make([]interface{}, 0)
	if err := models.GetBodegaByTerceroId(terceroId, &solicitudes); err != nil {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("404")
	} else {
		c.Data["json"] = solicitudes
	}

	c.ServeJSON()
}
