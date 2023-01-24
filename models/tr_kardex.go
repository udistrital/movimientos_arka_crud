package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type TrKardex struct {
	Kardex    *Movimiento
	Elementos []*ElementosMovimiento
}
type KardexGeneral struct {
	Movimiento *[]*TrKardex
}

type Apertura struct {
	CantidadMinima     int
	CantidadMaxima     int
	ElementoCatalogoId int
	FechaCreacion      time.Time
	MetodoValoracion   int
	SaldoCantidad      float64
	SaldoValor         float64
	Unidad             float64
	ValorUnitario      float64
	ValorTotal         float64
}

// AddTransaccionKardex Nuevo registro a una ficha kárdex
func AddTransaccionKardex(n *KardexGeneral) (err error) {
	o := orm.NewOrm()

	err = o.Begin()
	if err != nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			o.Rollback()
			logs.Error(r)
		} else {
			o.Commit()
		}
	}()

	for _, m := range *n.Movimiento {

		idSalida, err_ := o.Insert(m.Kardex)
		if err_ != nil {
			return err_
		}

		mov := Movimiento{Id: int(idSalida)}
		for _, elemento := range m.Elementos {
			elemento.MovimientoId = &mov
			var elemento__ ElementosMovimiento
			id__ := elemento.Id
			elemento.Id = 0

			_, err = o.Insert(elemento)
			if err != nil {
				return
			}

			_, err = o.QueryTable(new(ElementosMovimiento)).RelatedSel().Filter("Id", id__).All(&elemento__)
			if err != nil {
				return
			}

			elemento__.Activo = false
			_, err = o.Update(&elemento__, "Activo")
			if err != nil {
				return
			}

		}

	}

	return
}

func ResponderSolicitud(Solicitud *KardexGeneral) (err error) {
	o := orm.NewOrm()

	err = o.Begin()
	if err != nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			o.Rollback()
			logs.Error(r)
		} else {
			o.Commit()
		}
	}()

	err = AddTransaccionKardex(Solicitud)
	if err != nil {
		return
	}

	solicitud := *Solicitud.Movimiento
	var elemento__ Movimiento
	_, err = o.QueryTable(new(Movimiento)).RelatedSel().Filter("Id", solicitud[0].Kardex.MovimientoPadreId.Id).All(&elemento__)
	if err != nil {
		return
	}

	elemento__.EstadoMovimientoId = solicitud[0].Kardex.MovimientoPadreId.EstadoMovimientoId
	elemento__.Detalle = solicitud[0].Kardex.MovimientoPadreId.Detalle
	_, err = o.Update(&elemento__, "EstadoMovimientoId", "Detalle")
	return

}

func RechazarSolicitud(id *Movimiento) (err error) {
	o := orm.NewOrm()

	err = o.Begin()
	if err != nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			o.Rollback()
			logs.Error(r)
		} else {
			o.Commit()
		}
	}()

	var solicitud Movimiento
	err = o.QueryTable(new(Movimiento)).RelatedSel().Filter("Id", id.Id).One(&solicitud)
	if err != nil || solicitud.EstadoMovimientoId.Nombre != "Solicitud Pendiente" {
		if err == nil {
			err = errors.New(`solicitud.EstadoMovimientoId.Nombre != "Solicitud Pendiente"`)
		}
		return
	}

	solicitud.EstadoMovimientoId = id.EstadoMovimientoId
	solicitud.Detalle = id.Detalle
	_, err = o.Update(&solicitud, "EstadoMovimientoId", "Detalle")
	return
}

func GetAllAperturas(conSaldo bool) (aperturas []Apertura, err error) {

	o := orm.NewOrm()
	err = o.Begin()

	if err != nil {
		return
	}

	query :=
		`
	WITH aperturas AS (
		SELECT DISTINCT ON (1)
			em.elemento_catalogo_id,
			m.detalle
		FROM
			movimientos_arka.elementos_movimiento em,
			movimientos_arka.movimiento m,
			movimientos_arka.formato_tipo_movimiento tm
		WHERE
				tm.codigo_abreviacion = 'AP_KDX'
			AND m.formato_tipo_movimiento_id = tm.id
			AND em.movimiento_id = m.id
	), saldo AS (
		SELECT DISTINCT ON (1)
			em.elemento_catalogo_id,
			m.fecha_creacion,
			em.saldo_cantidad,
			em.saldo_valor,
			em.unidad,
			em.valor_total,
			em.valor_unitario,
			aperturas.detalle
		FROM
			movimientos_arka.elementos_movimiento em,
			movimientos_arka.movimiento m,
			movimientos_arka.formato_tipo_movimiento tm,
			aperturas
		WHERE
				tm.codigo_abreviacion IN ('AP_KDX', 'ENT_KDX', 'SAL_KDX')
			AND m.formato_tipo_movimiento_id = tm.id
			AND em.movimiento_id = m.id
			AND aperturas.elemento_catalogo_id = em.elemento_catalogo_id
		ORDER BY 1, 2 DESC
	)

	SELECT *,
		detalle ->> 'Metodo_Valoracion' metodo_valoracion,
		detalle ->> 'Cantidad_Minima' cantidad_minima,
		detalle ->> 'Cantidad_Maxima' cantidad_maxima
	FROM saldo
	`

	if conSaldo {
		query +=
			`
		WHERE saldo_cantidad > 0;
		`
	}

	_, err = o.Raw(query).QueryRows(&aperturas)
	return
}
