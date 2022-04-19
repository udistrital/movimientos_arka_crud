package models

import (
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type detalle struct {
	ValorPresente        float64
	ElementoMovimientoId int
	VidaUtil             float64
	ElementoActaId       int
	ValorResidual        float64
	NovedadElementoId    int
	FechaRef             time.Time
}

// GetCorteDepreciacion Retorna los valores y detalles necesarios para generar
// la transacción contable correspondiente a la depreciación dada una fecha de corte
func GetCorteDepreciacion(fechaCorte string) (entrada []*detalle, err error) {

	// Los elementos se determinan de la siguiente manera
	// + Elementos sin novedad y vida útil > 0
	// + Novedades con vida útil > 0
	// - Elementos solicitados para baja antes de la fecha de corte

	var (
		query     string
		elementos []*detalle
		novedades []*detalle
	)

	o := orm.NewOrm()
	query =
		`SELECT
		em.id as elemento_movimiento_id,
		em.valor_residual,
		em.vida_util,
		em.elemento_acta_id,
		em.valor_total as valor_presente,
		m.fecha_modificacion as fecha_ref
	FROM
		movimientos_arka.elementos_movimiento em,
		movimientos_arka.movimiento m,
		movimientos_arka.estado_movimiento sm,
		movimientos_arka.formato_tipo_movimiento fm
	WHERE
		fm.nombre = 'Salida'
		AND sm.nombre = 'Salida Aprobada'
		AND m.formato_tipo_movimiento_id  = fm.id
		AND m.fecha_modificacion < ?
		AND m.estado_movimiento_id = sm.id
		AND em.movimiento_id = m.id
		AND em.valor_total > 0
		AND em.vida_util > 0
		AND NOT EXISTS (
			SELECT
			FROM
				movimientos_arka.novedad_elemento ne
			WHERE
				ne.elemento_movimiento_id = em.id)
		AND em.id NOT IN (
			SELECT
				em.id
			FROM
				 movimientos_arka.movimiento m,
				 movimientos_arka.elementos_movimiento em,
				 movimientos_arka.formato_tipo_movimiento fm,
				 jsonb_array_elements(m.detalle -> 'Elementos') AS elem
			WHERE
				fm.codigo_abreviacion LIKE 'BJ%'
				AND m.formato_tipo_movimiento_id = fm.id
				AND m.fecha_creacion < ?
				AND em.id = CAST(elem as INTEGER));`

	if _, err = o.Raw(query, fechaCorte, fechaCorte).QueryRows(&elementos); err != nil {
		return nil, err
	}

	query =
		`SELECT
		DISTINCT ON (ne.elemento_movimiento_id) ne.elemento_movimiento_id,
		ne.id as novedad_elemento_id,
		ne.vida_util,
		ne.valor_residual,
		ne.valor_libros as valor_presente,
		m.detalle ->> 'FechaCorte' as fecha_ref,
		em.elemento_acta_id
	FROM
		movimientos_arka.novedad_elemento ne,
		movimientos_arka.elementos_movimiento em,
		movimientos_arka.movimiento m
	WHERE
		em.fecha_creacion < ?
		AND ne.movimiento_id = m.id
		AND ne.elemento_movimiento_id = em.id
		AND ne.activo = true
		AND ne.valor_libros > 0
		AND ne.vida_util > 0
		AND ne.elemento_movimiento_id NOT IN (
			SELECT
				ne.elemento_movimiento_id
			FROM
				movimientos_arka.movimiento m,
				movimientos_arka.novedad_elemento ne,
				movimientos_arka.elementos_movimiento em,
				movimientos_arka.formato_tipo_movimiento fm,
				jsonb_array_elements(m.detalle -> 'Elementos') AS elem
			WHERE
				fm.codigo_abreviacion LIKE 'BJ%'
				AND m.formato_tipo_movimiento_id = fm.id
				AND m.fecha_creacion < ?
				AND ne.elemento_movimiento_id = CAST(elem as INTEGER))
	ORDER BY
		ne.elemento_movimiento_id,
		ne.id DESC;`

	if _, err = o.Raw(query, fechaCorte, fechaCorte).QueryRows(&novedades); err != nil {
		return nil, err
	}
	elementos = append(elementos, novedades...)

	return elementos, nil
}

// AddNovedadElemento insert a new NovedadElemento into database and returns
// last inserted Id on success.
func AddTrNovedadElemento(m *NovedadElemento) (id int64, err error) {
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

	var novedades []*NovedadElemento

	if _, err := o.QueryTable(new(NovedadElemento)).RelatedSel().Filter("ElementoMovimientoId__Id", m.ElementoMovimientoId).Filter("Activo", true).All(&novedades, "Id"); err == nil {
		for _, nv := range novedades {
			nv.Activo = false
			if _, err = o.Update(nv, "Activo"); err != nil {
				panic(err.Error())
			}
		}
		if id, err = o.Insert(m); err != nil {
			panic(err.Error())
		}
	} else {
		panic(err.Error())
	}

	return
}
