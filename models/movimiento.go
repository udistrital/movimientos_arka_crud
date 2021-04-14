package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Movimiento struct {
	Id                      int                    `orm:"column(id);pk;auto"`
	Observacion             string                 `orm:"column(observacion);null"`
	Detalle                 string                 `orm:"column(detalle);type(json)"`
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
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMovimientoById retrieves Movimiento by Id. Returns error if
// Id doesn't exist
func GetMovimientoById(id int) (v *Movimiento, err error) {
	o := orm.NewOrm()
	v = &Movimiento{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMovimiento retrieves all Movimiento matches certain condition. Returns empty list if
// no records exist
func GetAllMovimiento(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Movimiento)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
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
		return ml, nil
	}
	return nil, err
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

// GetEntradaByActa retrieves all Movimiento matches an specific acta_recibido_id. Returns empty list if
// no records exist
func GetEntradaByActa(acta_recibido_id int) (entrada []Movimiento, err error) {
	var estadoMovimiento []int
	var movimientos []Movimiento

	estados := []string{"Entrada Aceptada", "Entrada Con Salida"}
	query_estado := "SELECT e.id FROM movimientos_arka.estado_movimiento e WHERE e.nombre IN (?, ?)"
	query_movimiento := "SELECT * FROM movimientos_arka.movimiento m  WHERE CAST(m.detalle ->>'acta_recibido_id' as INTEGER) = ? AND m.estado_movimiento_id IN (?, ?)"

	o := orm.NewOrm()
	_, err = o.Raw(query_estado, estados).QueryRows(&estadoMovimiento)
	_, err = o.Raw(query_movimiento, acta_recibido_id, estadoMovimiento).QueryRows(&movimientos)

	return movimientos, err
}
