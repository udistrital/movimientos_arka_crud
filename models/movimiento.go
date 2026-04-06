package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/formatdata"
)

type Movimiento struct {
	Id                      int                    `orm:"column(id);pk;auto"`
	Observacion             string                 `orm:"column(observacion);null"`
	ConsecutivoId           *int                   `orm:"column(consecutivo_id);null"`
	Consecutivo             *string                `orm:"column(consecutivo);null"`
	FechaCorte              *time.Time             `orm:"column(fecha_corte);null"`
	Detalle                 string                 `orm:"column(detalle);type(jsonb)"`
	FechaCreacion           time.Time              `orm:"auto_now_add;column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion       time.Time              `orm:"auto_now;column(fecha_modificacion);type(timestamp without time zone)"`
	Activo                  bool                   `orm:"column(activo)"`
	MovimientoPadreId       *Movimiento            `orm:"column(movimiento_padre_id);rel(fk);null"`
	FormatoTipoMovimientoId *FormatoTipoMovimiento `orm:"column(formato_tipo_movimiento_id);rel(fk)"`
	EstadoMovimientoId      *EstadoMovimiento      `orm:"column(estado_movimiento_id);rel(fk)"`
}

func (t *Movimiento) TableName() string {
	return "movimiento"
}

func init() {
	orm.RegisterModel(new(Movimiento))
}

// AddMovimiento insert a new Movimiento into database and returns
// last inserted Id on success.
func AddMovimiento(m *Movimiento) (id int64, err error) {
	logs.Info("[model:AddMovimiento] ===== INICIO =====")
	logs.Info("[model:AddMovimiento] Struct recibido ANTES del insert:")
	logs.Info("[model:AddMovimiento]   Id=%d (debe ser 0 para auto-increment)", m.Id)
	logs.Info("[model:AddMovimiento]   Observacion=%q", m.Observacion)
	logs.Info("[model:AddMovimiento]   Activo=%v", m.Activo)
	logs.Info("[model:AddMovimiento]   Detalle=%q", m.Detalle)

	if m.FechaCorte != nil {
		logs.Info("[model:AddMovimiento]   FechaCorte=%v", *m.FechaCorte)
	} else {
		logs.Info("[model:AddMovimiento]   FechaCorte=nil")
	}
	if m.ConsecutivoId != nil {
		logs.Info("[model:AddMovimiento]   ConsecutivoId=%d", *m.ConsecutivoId)
	} else {
		logs.Info("[model:AddMovimiento]   ConsecutivoId=nil")
	}
	if m.Consecutivo != nil {
		logs.Info("[model:AddMovimiento]   Consecutivo=%q", *m.Consecutivo)
	} else {
		logs.Info("[model:AddMovimiento]   Consecutivo=nil")
	}

	if m.FormatoTipoMovimientoId != nil {
		logs.Info("[model:AddMovimiento]   FormatoTipoMovimientoId.Id=%d (este es el FK que se insertará)", m.FormatoTipoMovimientoId.Id)
	} else {
		logs.Error("[model:AddMovimiento]   FormatoTipoMovimientoId=NIL — esto causará error o insertará 0")
	}

	if m.EstadoMovimientoId != nil {
		logs.Info("[model:AddMovimiento]   EstadoMovimientoId.Id=%d (este es el FK que se insertará)", m.EstadoMovimientoId.Id)
	} else {
		logs.Error("[model:AddMovimiento]   EstadoMovimientoId=NIL — esto causará error o insertará 0")
	}

	if m.MovimientoPadreId != nil {
		logs.Info("[model:AddMovimiento]   MovimientoPadreId.Id=%d", m.MovimientoPadreId.Id)
	} else {
		logs.Info("[model:AddMovimiento]   MovimientoPadreId=nil")
	}

	mJSON, _ := json.MarshalIndent(m, "", "  ")
	logs.Info("[model:AddMovimiento] Struct completo serializado:\n%s", string(mJSON))

	o := orm.NewOrm()
	logs.Info("[model:AddMovimiento] Ejecutando o.Insert(m)...")
	id, err = o.Insert(m)

	if err != nil {
		logs.Error("[model:AddMovimiento] INSERT FALLÓ — error: %v", err)
		logs.Error("[model:AddMovimiento] Tipo de error: %T", err)
		logs.Error("[model:AddMovimiento] m.Id después del insert fallido: %d", m.Id)
	} else {
		logs.Info("[model:AddMovimiento] INSERT OK")
		logs.Info("[model:AddMovimiento]   id retornado por o.Insert: %d", id)
		logs.Info("[model:AddMovimiento]   m.Id después del insert: %d", m.Id)
		logs.Info("[model:AddMovimiento]   (ambos deben coincidir y ser > 0)")

		if id <= 0 {
			logs.Error("[model:AddMovimiento] ALERTA: id retornado es <= 0, posible insert fantasma")
		}
		if m.Id <= 0 {
			logs.Error("[model:AddMovimiento] ALERTA: m.Id es <= 0 después del insert, el ORM no asignó el auto-increment")
		}
	}

	logs.Info("[model:AddMovimiento] ===== FIN =====")
	return
}

// GetMovimientoById retrieves Movimiento by Id. Returns error if
// Id doesn't exist
func GetMovimientoById(id int) (v Movimiento, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(new(Movimiento)).RelatedSel().Filter("Id", id).One(&v)

	return
}

// GetAllMovimiento retrieves all Movimiento matches certain condition. Returns empty list if
// no records exist
func GetAllMovimiento(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Movimiento)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else if strings.Contains(k, "__in") {
			arr := strings.Split(v, "|")
			qs = qs.Filter(k, arr)
		} else {
			qs = qs.Filter(k, v)
		}
	}
	count, err = qs.Count()
	if err != nil {
		return
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					err = errors.New("error: Invalid order. Must be either [asc|desc]")
					return
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					err = errors.New("error: Invalid order. Must be either [asc|desc]")
					return
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			err = errors.New("error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
			return
		}
	} else {
		if len(order) != 0 {
			err = errors.New("error: unused 'order' fields")
			return
		}
	}

	var l []Movimiento
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
	}
	return
}

// UpdateMovimiento updates Movimiento by Id and returns error if
// the record to be updated doesn't exist
func UpdateMovimientoById(m *Movimiento) (err error) {
	o := orm.NewOrm()
	v := Movimiento{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMovimiento deletes Movimiento by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMovimiento(id int) (err error) {
	o := orm.NewOrm()
	v := Movimiento{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Movimiento{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// GetEntradaByActa Retorna la entrada asociada a un acta determinada
func GetEntradaByActa(acta_recibido_id int) (entrada *Movimiento, err error) {

	o := orm.NewOrm()
	var ids []int
	query :=
		`SELECT	m.id
	FROM movimientos_arka.movimiento m
	WHERE m.detalle ->> 'acta_recibido_id' = ?;`

	if _, err = o.Raw(query, acta_recibido_id).QueryRows(&ids); err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return nil, nil
	}

	if l, _, err := GetAllMovimiento(
		map[string]string{"Id": strconv.Itoa(ids[0])}, []string{}, nil, nil, 0, -1); err != nil {
		return nil, err
	} else {
		var movs []*Movimiento
		if err := formatdata.FillStruct(l, &movs); err != nil {
			return nil, err
		}
		entrada = movs[0]
	}

	return entrada, nil
}

// GetTrasladosByTerceroId Retorna la entrada asociada a un acta determinada
func GetTrasladosByTerceroId(terceroId int, porRecibir bool, traslados *[]Movimiento) (err error) {

	o := orm.NewOrm()

	var ids []int

	query :=
		`SELECT	m.id
	FROM ` + esquema + `.movimiento m,` +
			esquema + `.formato_tipo_movimiento fm`

	if porRecibir {
		query += `,
		` + esquema + `.estado_movimiento em`
	}

	query +=
		`
		WHERE fm.codigo_abreviacion = 'SOL_TRD'
		AND (m.detalle ->> 'FuncionarioDestino' = ?`

	if !porRecibir {
		query += `
		OR m.detalle ->> 'FuncionarioOrigen' = ?);`
	} else {
		query += `)
		AND em.nombre = 'Traslado Por Confirmar'
		AND m.estado_movimiento_id = em.id;`
	}

	if porRecibir {
		if _, err = o.Raw(query, terceroId).QueryRows(&ids); err != nil {
			return err
		}
	} else {
		if _, err = o.Raw(query, terceroId, terceroId).QueryRows(&ids); err != nil {
			return err
		}
	}

	if len(ids) == 0 {
		return nil
	}

	if l, _, err := GetAllMovimiento(
		map[string]string{"Id__in": strings.Trim(strings.Replace(fmt.Sprint(ids), " ", "|", -1), "[]")}, []string{}, nil, nil, 0, -1); err != nil {
		return err
	} else {
		var movs []Movimiento
		if err := formatdata.FillStruct(l, &movs); err != nil {
			return err
		}
		*traslados = movs
	}

	return
}

// GetBajasByTerceroId Retorna la lista de bajas solicitadas por un funcionario determinado
func GetBajasByTerceroId(terceroId int, bajas *[]interface{}) (err error) {

	o := orm.NewOrm()

	var ids []int

	query :=
		`
		SELECT	m.id
		FROM ` + esquema + `.movimiento m,` +
			esquema + `.estado_movimiento sm
		WHERE
			sm.nombre LIKE 'Baja%'
			AND m.estado_movimiento_id = sm.id
			AND m.detalle ->> 'Funcionario' = ?;
		`

	if _, err = o.Raw(query, terceroId).QueryRows(&ids); err != nil {
		return err
	}

	if len(ids) == 0 {
		return
	}

	if l, _, err := GetAllMovimiento(
		map[string]string{"Id__in": strings.Trim(strings.Replace(fmt.Sprint(ids), " ", "|", -1), "[]")}, []string{}, nil, nil, 0, -1); err != nil {
		return err
	} else {
		*bajas = l
	}

	return
}

// GetBodegaByTerceroId Retorna la lista de solicitudes de bodega de consumo por un funcionario determinado
func GetBodegaByTerceroId(terceroId int, solicitudes *[]interface{}) (err error) {

	o := orm.NewOrm()

	var ids []int

	query :=
		`
		SELECT	m.id
		FROM ` + esquema + `.movimiento m,` +
			esquema + `.formato_tipo_movimiento fm
		WHERE
			fm.codigo_abreviacion = 'SOL_BOD'
			AND m.formato_tipo_movimiento_id = fm.id
			AND m.detalle ->> 'Funcionario' = ?;
		`

	if _, err = o.Raw(query, terceroId).QueryRows(&ids); err != nil {
		return err
	}

	if len(ids) == 0 {
		return
	}

	if l, _, err := GetAllMovimiento(
		map[string]string{"Id__in": strings.Trim(strings.Replace(fmt.Sprint(ids), " ", "|", -1), "[]")}, []string{}, nil, nil, 0, -1); err != nil {
		return err
	} else {
		*solicitudes = l
	}

	return
}
