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
	Elemento  *ElementosMovimiento
	Entradas  []*Movimiento
	Salida    *Movimiento
	Traslados []*Movimiento
	Novedades []NovedadElemento
	Baja      *Movimiento
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

// GetElementosFuncionario Consulta los elementos en poder del funcionario sin bajas o traslados pendientes
func GetElementosFuncionario(funcionarioId int, elementos *[]int) (err error) {

	// Los elementos se determinan de la siguiente manera
	// + Elementos asignados en una salida
	// - Elementos trasladados a otro funcionario
	// + Elementos trasladados desde otro funcionario
	// - Elementos dados de baja
	// - Elementos solicitados para traslado

	o := orm.NewOrm()
	query := `
	WITH
		asignados AS (
			SELECT
				DISTINCT em.id as elemento_id
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
		),	trasladados AS (
			SELECT DISTINCT ON (1)
				em.id as elemento_id,
				func_o as origen,
				func_d as destino
			FROM
				movimientos_arka.movimiento m,
				movimientos_arka.estado_movimiento sm,
				movimientos_arka.elementos_movimiento em,
				jsonb_array_elements(m.detalle -> 'Elementos') AS elem,
				jsonb(m.detalle -> 'FuncionarioOrigen') AS func_o,
				jsonb(m.detalle -> 'FuncionarioDestino') AS func_d
			WHERE
				sm.nombre = 'Traslado Aprobado'
				AND CAST(elem as INTEGER) = em.id
				AND m.estado_movimiento_id = sm.id
				AND (CAST(func_o as INTEGER) = ?
				OR CAST(func_d as INTEGER) = ?)
			ORDER BY em.id, m.id DESC
		), recibidos AS (
			SELECT elemento_id
			FROM asignados
			UNION
			SELECT elemento_id
			FROM trasladados
			WHERE destino = ?
		), bajas AS (
			SELECT DISTINCT ON (1)
				em.id as elemento_id
			FROM
				movimientos_arka.movimiento m,
				movimientos_arka.estado_movimiento sm,
				movimientos_arka.elementos_movimiento em,
				jsonb_array_elements(m.detalle -> 'Elementos') AS elem,
				recibidos
			WHERE
				(
					sm.nombre LIKE 'Baja%' OR
					sm.nombre IN ('Traslado Por Confirmar','Traslado Rechazado','Traslado Confirmado')
				)
				AND m.estado_movimiento_id = sm.id
				AND CAST(elem as INTEGER) = em.id
				AND em.id = recibidos.elemento_id
		), entregados AS (
			SELECT elemento_id
			FROM bajas
			UNION
			SELECT elemento_id
			FROM trasladados
			WHERE origen = ?
		)

	SELECT elemento_id
	FROM recibidos
	EXCEPT
	SELECT elemento_id
	FROM entregados;`

	if _, err = o.Raw(query, funcionarioId, funcionarioId, funcionarioId, funcionarioId, funcionarioId).QueryRows(elementos); err != nil {
		return err
	}

	return
}

// Retorna los movimientos que han involucrado un elemento
func GetHistorialElemento(elementoId int, acta, entradas, novedades, final bool, historial *Historial) (err error) {

	var ids []int

	o := orm.NewOrm()
	query := ""
	if acta {
		query = "ElementoActaId"
	} else {
		query = "Id"
	}

	var elemento ElementosMovimiento
	_, err = o.QueryTable(new(ElementosMovimiento)).RelatedSel().Filter(query, strconv.Itoa(elementoId)).All(&elemento)
	if err != nil || elemento.Id == 0 {
		return
	}

	elementoId = elemento.Id
	historial.Elemento = &elemento
	historial.Salida = elemento.MovimientoId

	query =
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

	_, err = o.Raw(query, elementoId).QueryRows(&ids)
	if err != nil {
		return
	} else if ids != nil && len(ids) > 0 {
		l, err := GetAllMovimiento(
			map[string]string{"Id__in": ArrayToString(ids, "|")}, []string{}, nil, nil, 0, -1)
		if err != nil {
			return err
		}

		err = formatdata.FillStruct(l, &historial.Traslados)
		if err != nil {
			return err
		}
	}

	ids = []int{}
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

	_, err = o.Raw(query, elementoId).QueryRows(&ids)
	if err != nil {
		return err
	} else if ids != nil && len(ids) > 0 {
		if l, err := GetAllMovimiento(
			map[string]string{"Id__in": ArrayToString(ids, "|")}, []string{}, nil, nil, 0, -1); err != nil {
			return err
		} else {
			var baja []*Movimiento
			if err := formatdata.FillStruct(l, &baja); err != nil {
				return err
			}
			historial.Baja = baja[0]
		}
	}

	if entradas {
		ids = []int{}
		query = `
		SELECT
			m.id
		FROM
			movimientos_arka.movimiento m,
			movimientos_arka.estado_movimiento sm,
			jsonb_array_elements(m.detalle #> '{elementos}') elementos
		WHERE
				sm.nombre LIKE 'Entrada%'
			AND m.estado_movimiento_id = sm.id
			AND (elementos ->> 'Id')::int = ?
		ORDER BY m.fecha_corte DESC;`

		_, err = o.Raw(query, elementoId).QueryRows(&ids)
		if err != nil {
			return err
		} else if ids != nil && len(ids) > 0 {
			l, err := GetAllMovimiento(map[string]string{"Id__in": ArrayToString(ids, "|")}, []string{}, nil, nil, 0, -1)
			if err != nil {
				return err
			}

			err = formatdata.FillStruct(l, &historial.Entradas)
			if err != nil {
				return err
			}
		}
	}

	if novedades {
		_, err = o.QueryTable(new(NovedadElemento)).RelatedSel().
			Filter("ElementoMovimientoId__Id", elementoId).OrderBy("-MovimientoId__FechaCorte").Limit(1).All(&historial.Novedades)
	}

	return
}

func ArrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
