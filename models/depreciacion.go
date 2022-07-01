package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type DepreciacionElemento struct {
	DeltaValor           float64
	ElementoMovimientoId int
	ElementoActaId       int
}

// GetCorteDepreciacion Retorna los valores y detalles necesarios para generar
// la transacción contable correspondiente a la depreciación dada una fecha de corte
func GetCorteDepreciacion(fechaCorte string, elementos *[]*DepreciacionElemento) (err error) {

	o := orm.NewOrm()

	// Se confirma que la fecha de corte no es anterior a la de un cierre existente
	query :=
		`SELECT m.id
		FROM
			movimientos_arka.movimiento m,
			movimientos_arka.formato_tipo_movimiento fm,
			movimientos_arka.estado_movimiento sm,
			TO_DATE(?, 'YYYY-MM-DD') AS fecha_corte,
			TO_DATE(m.detalle->>'FechaCorte', 'YYYY-MM-DD') AS fecha
		WHERE
			fm.codigo_abreviacion = 'DEP'
			AND m.formato_tipo_movimiento_id = fm.id
			AND sm.nombre = 'Cierre Aprobado'
			AND fecha >= fecha_corte
		LIMIT 1;`

	cierres := make([]*Movimiento, 0)
	if _, err = o.Raw(query, fechaCorte).QueryRows(&cierres); err != nil {
		return err
	} else if len(cierres) > 0 {
		return
	}

	// Los elementos se determinan de la siguiente manera
	// + Elementos sin novedad y vida útil > 0
	// + Novedades con vida útil > 0
	// - Elementos solicitados para baja antes de la fecha de corte
	query =
		`WITH fecha_corte AS (
			SELECT (TO_DATE(?, 'YYYY-MM-DD') + INTERVAL '1 day')::date fecha_corte
		), bajas AS (
			SELECT
				em.id
			FROM
				movimientos_arka.movimiento m,
				movimientos_arka.elementos_movimiento em,
				movimientos_arka.formato_tipo_movimiento fm,
				jsonb_array_elements(m.detalle -> 'Elementos') AS elem,
				fecha_corte
			WHERE
				fm.codigo_abreviacion LIKE 'BJ%'
				AND m.formato_tipo_movimiento_id = fm.id
				AND m.fecha_creacion < fecha_corte
				AND em.id = CAST(elem as INTEGER)
		), sin_novedad AS (
			SELECT
				em.id elemento_movimiento_id,
				em.valor_residual,
				em.vida_util,
				em.elemento_acta_id,
				em.valor_total valor_presente,
				date_part('day', fecha_corte - (m.fecha_modificacion::date)::timestamp) delta_tiempo
			FROM
				movimientos_arka.elementos_movimiento em,
				movimientos_arka.movimiento m,
				movimientos_arka.estado_movimiento sm,
				movimientos_arka.formato_tipo_movimiento fm,
				fecha_corte
			WHERE
				fm.codigo_abreviacion = 'SAL'
				AND sm.nombre = 'Salida Aprobada'
				AND m.formato_tipo_movimiento_id  = fm.id
				AND m.fecha_modificacion < fecha_corte
				AND m.estado_movimiento_id = sm.id
				AND em.movimiento_id = m.id
				AND em.valor_total > 0
				AND em.vida_util > 0
				AND em.id NOT IN (
					SELECT 
						em.id
					FROM
						movimientos_arka.novedad_elemento ne
					WHERE
						ne.elemento_movimiento_id = em.id
				)
				AND em.id NOT IN (
					SELECT
						em.id
					FROM 
						bajas
					WHERE
						bajas.id = em.id
				)
		), con_novedad AS (
			SELECT DISTINCT ON (1)
				ne.elemento_movimiento_id,
				fecha,
				ne.valor_residual,
				ne.vida_util,
				em.elemento_acta_id,
				ne.valor_libros valor_presente,
				date_part('day', fecha_corte - (fecha::date)::timestamp) as delta_tiempo
			FROM
				movimientos_arka.novedad_elemento ne,
				movimientos_arka.elementos_movimiento em,
				movimientos_arka.movimiento m,
				to_date(m.detalle->>'FechaCorte', 'YYYY-MM-DD') AS fecha,
				fecha_corte
			WHERE
				fecha < fecha_corte
				AND	ne.elemento_movimiento_id = em.id
				AND ne.movimiento_id = m.id
				AND ne.vida_util > 0
				AND ne.valor_libros > ne.valor_residual
				AND em.id NOT IN (
					SELECT
						bajas.id
					FROM 
						bajas
					WHERE
						bajas.id = em.id
				)
			ORDER BY 1 DESC, 2 DESC
		), referencia AS (
			SELECT *
			FROM
				sin_novedad ndp
			UNION
			SELECT
				dp.elemento_movimiento_id,
				dp.valor_residual,
				dp.vida_util,
				dp.elemento_acta_id,
				dp.valor_presente,
				dp.delta_tiempo - 1 delta_tiempo
			FROM
				con_novedad dp
		), delta_valor AS (
			SELECT
				sn.elemento_movimiento_id,
				sn.elemento_acta_id,
				sn.valor_presente,
				sn.delta_tiempo,
				(sn.valor_presente - sn.valor_residual) * (sn.delta_tiempo / (sn.vida_util * 365)) delta_valor
			FROM
				referencia sn
		), no_depreciados AS (
			SELECT
				dv.elemento_movimiento_id,
				dv.elemento_acta_id,
				dv.delta_valor
			FROM
				delta_valor dv
			WHERE
				dv.valor_presente - dv.delta_valor >= 0
		), depreciados AS (
			SELECT
				dv.elemento_movimiento_id,
				dv.elemento_acta_id,
				dv.valor_presente as delta_valor
			FROM
				delta_valor dv
			WHERE
				dv.valor_presente - dv.delta_valor < 0
		), calculados AS (
			SELECT ndp.*
			FROM no_depreciados ndp
			UNION
			SELECT dp.*
			FROM depreciados dp
		)
		
		SELECT * from calculados;`

	if _, err = o.Raw(query, fechaCorte).QueryRows(elementos); err != nil {
		return err
	}

	return
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

	if _, err = o.QueryTable(new(NovedadElemento)).RelatedSel().Filter("ElementoMovimientoId__Id", m.ElementoMovimientoId).Filter("Activo", true).All(&novedades, "Id"); err == nil {
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
