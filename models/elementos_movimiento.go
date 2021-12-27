package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/formatdata"
)

type ElementosMovimiento struct {
	Id                 int         `orm:"column(id);pk;auto"`
	ElementoActaId     int         `orm:"column(elemento_acta_id)"`
	ElementoCatalogoId int         `orm:"column(elemento_catalogo_id)"`
	Unidad             float64     `orm:"column(unidad)"`
	ValorUnitario      float64     `orm:"column(valor_unitario)"`
	ValorTotal         float64     `orm:"column(valor_total)"`
	SaldoCantidad      float64     `orm:"column(saldo_cantidad)"`
	VidaUtil           float64     `orm:"column(vida_util)"`
	ValorResidual      float64     `orm:"column(valor_residual)"`
	SaldoValor         float64     `orm:"column(saldo_valor)"`
	Activo             bool        `orm:"column(activo)"`
	FechaCreacion      time.Time   `orm:"auto_now_add;column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion  time.Time   `orm:"auto_now;column(fecha_modificacion);type(timestamp without time zone)"`
	MovimientoId       *Movimiento `orm:"column(movimiento_id);rel(fk)"`
}

type Historial struct {
	Salida       *Movimiento
	Traslados    []*Movimiento
	Baja         *Movimiento
	Depreciacion *Movimiento
}

func (t *ElementosMovimiento) TableName() string {
	return "elementos_movimiento"
}

func init() {
	orm.RegisterModel(new(ElementosMovimiento))
}

// AddElementosMovimiento insert a new ElementosMovimiento into database and returns
// last inserted Id on success.
func AddElementosMovimiento(m *ElementosMovimiento) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetElementosMovimientoById retrieves ElementosMovimiento by Id. Returns error if
// Id doesn't exist
func GetElementosMovimientoById(id int) (v *ElementosMovimiento, err error) {
	o := orm.NewOrm()
	v = &ElementosMovimiento{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllElementosMovimiento retrieves all ElementosMovimiento matches certain condition. Returns empty list if
// no records exist
func GetAllElementosMovimiento(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ElementosMovimiento)).RelatedSel()
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
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
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
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []ElementosMovimiento
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
		return ml, nil
	}
	return nil, err
}

// UpdateElementosMovimiento updates ElementosMovimiento by Id and returns error if
// the record to be updated doesn't exist
func UpdateElementosMovimientoById(m *ElementosMovimiento) (err error) {
	o := orm.NewOrm()
	v := ElementosMovimiento{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteElementosMovimiento deletes ElementosMovimiento by Id and returns error if
// the record to be deleted doesn't exist
func DeleteElementosMovimiento(id int) (err error) {
	o := orm.NewOrm()
	v := ElementosMovimiento{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ElementosMovimiento{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// GetElementosFuncionario retrieves all Movimiento matches an specific acta_recibido_id. Returns empty list if
func GetElementosFuncionario(funcionarioId int) (entrada []int, err error) {

	// Los elementos se determinan de la siguiente manera
	// + Elementos asignados en una salida
	// - Elementos trasladados a otro funcionario
	// + Elementos trasladados desde otro funcionario
	// - Elementos dados de baja
	// - Elementos solicitados para traslado

	var (
		query       string
		elementos   []int
		trasladados []int
	)

	o := orm.NewOrm()
	queryC :=
		`bajas AS (
			SELECT
				DISTINCT em.id as bajas
			FROM
				movimientos_arka.movimiento m,
				movimientos_arka.estado_movimiento sm,
				movimientos_arka.elementos_movimiento em,
				jsonb_array_elements(m.detalle -> 'Elementos') AS elem
			WHERE
				sm.nombre LIKE 'Baja%'
				AND CAST(elem as INTEGER) = em.id
				AND m.estado_movimiento_id = sm.id
			),
		pendientes AS (
				SELECT
					DISTINCT em.id as pendientes
				FROM
					movimientos_arka.movimiento m,
					movimientos_arka.estado_movimiento sm,
					movimientos_arka.elementos_movimiento em,
					jsonb_array_elements(m.detalle -> 'Elementos') AS elem
				WHERE
					sm.nombre IN ('Traslado En Trámite','Traslado Rechazado','Traslado Aprobado')
					AND CAST(elem as INTEGER) = em.id
					AND m.estado_movimiento_id = sm.id
				),
		excepto as (
			SELECT bajas
			FROM bajas
			UNION
			SELECT pendientes
			FROM pendientes
				)

		SELECT *
		FROM elementos
		EXCEPT
		SELECT *
		FROM excepto;`
	query =
		`WITH
			elementos AS (
				SELECT
					DISTINCT em.id as elementos
				FROM
					movimientos_arka.movimiento m,
					movimientos_arka.estado_movimiento sm,
					movimientos_arka.elementos_movimiento em,
					jsonb(m.detalle -> 'funcionario') AS func
				WHERE
					sm.nombre = 'Salida Aprobada'
					AND m.estado_movimiento_id = sm.id
					AND em.movimiento_id = m.id
					AND CAST(func as INTEGER) = ?
				),`

	if _, err = o.Raw(query+queryC, funcionarioId).QueryRows(&elementos); err != nil {
		return nil, err
	}

	query =
		`WITH
			elementos AS (
				SELECT
					DISTINCT em.id as elementos
				FROM
					movimientos_arka.movimiento m,
					movimientos_arka.estado_movimiento sm,
					movimientos_arka.elementos_movimiento em,
					jsonb_array_elements(m.detalle -> 'Elementos') AS elem,
					jsonb(m.detalle -> 'FuncionarioOrigen') AS func_o,
					jsonb(m.detalle -> 'FuncionarioDestino') AS func_d
				WHERE
					sm.nombre = 'Traslado Confirmado'
					AND CAST(elem as INTEGER) = em.id
					AND m.estado_movimiento_id = sm.id
					AND (CAST(func_o as INTEGER) = ?
					OR CAST(func_d as INTEGER) = ?)
				),`

	if _, err = o.Raw(query+queryC, funcionarioId, funcionarioId).QueryRows(&trasladados); err != nil {
		return nil, err
	}

	var entregados []int
	var recibidos []int
	for _, el := range trasladados {
		var funcionario []int
		query =
			`SELECT
				func_o
			FROM
				movimientos_arka.movimiento m,
				movimientos_arka.estado_movimiento sm,
				jsonb_array_elements(m.detalle -> 'Elementos') AS elem,
				jsonb(m.detalle -> 'FuncionarioOrigen') AS func_o,
				jsonb(m.detalle -> 'FuncionarioDestino') AS func_d
			WHERE
				sm.nombre = 'Traslado Confirmado'
				AND m.estado_movimiento_id = sm.id
				AND CAST(elem as integer) = ?
				AND (CAST(func_o as INTEGER) = ?
				OR CAST(func_d as INTEGER) = ?)
			ORDER BY
				m.id DESC
			LIMIT
				1;`

		if _, err = o.Raw(query, el, funcionarioId, funcionarioId).QueryRows(&funcionario); err != nil {
			return nil, err
		}

		if funcionario[0] == funcionarioId {
			entregados = append(entregados, el)
		} else {
			recibidos = append(recibidos, el)
		}
	}

	for _, ent := range entregados {
		for i, rec := range elementos {
			if ent == rec {
				elementos = append(elementos[:i], elementos[i+1:]...)
			}
		}
	}

	elementos = append(elementos, recibidos...)

	return elementos, nil
}

// Retorna los movimientos que han involucrado un elemento
func GetHistorialElemento(elementoId int, final bool) (historial *Historial, err error) {

	var (
		baja      []int
		traslados []int
	)

	historial = new(Historial)

	o := orm.NewOrm()
	if l, err := GetAllElementosMovimiento(
		map[string]string{"Id": strconv.Itoa(elementoId)}, []string{}, nil, nil, 0, -1); err != nil {
		return nil, err
	} else {
		var salida_ []*ElementosMovimiento
		if err := formatdata.FillStruct(l, &salida_); err != nil {
			return nil, err
		}
		historial.Salida = salida_[0].MovimientoId
	}

	query :=
		`SELECT
			m.id
		FROM
			movimientos_arka.movimiento m,
			movimientos_arka.estado_movimiento sm
		WHERE
			sm.nombre LIKE 'Traslado%'
			AND m.estado_movimiento_id = sm.id
			AND m.detalle->'Elementos' @> ?
		ORDER BY
			m.id DESC`

	if final {
		query += " LIMIT 1"
	}

	if _, err = o.Raw(query, elementoId).QueryRows(&traslados); err != nil {
		return nil, err
	} else if traslados != nil {
		if l, err := GetAllMovimiento(
			map[string]string{"Id__in": ArrayToString(traslados, "|")}, []string{}, nil, nil, 0, -1); err != nil {
			return nil, err
		} else {
			var traslados_ []*Movimiento
			if err := formatdata.FillStruct(l, &traslados_); err != nil {
				return nil, err
			}
			historial.Traslados = traslados_
		}
	}

	query =
		`SELECT
			m.id
		FROM
			movimientos_arka.movimiento m,
			movimientos_arka.estado_movimiento sm
		WHERE
			sm.nombre LIKE 'Baja%'
			AND m.estado_movimiento_id = sm.id
			AND m.detalle->'Elementos' @> ?
		ORDER BY m.id DESC
		LIMIT 1`

	if _, err = o.Raw(query, elementoId).QueryRows(&baja); err != nil {
		return nil, err
	} else if baja != nil {
		if l, err := GetAllMovimiento(
			map[string]string{"Id__in": ArrayToString(baja, "|")}, []string{}, nil, nil, 0, -1); err != nil {
			return nil, err
		} else {
			var baja []*Movimiento
			if err := formatdata.FillStruct(l, &baja); err != nil {
				return nil, err
			}
			historial.Baja = baja[0]
		}
	}

	return historial, nil
}

func ArrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
