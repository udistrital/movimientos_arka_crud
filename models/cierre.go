package models

import (
	"strings"

	"github.com/astaxie/beego"
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
func GetCorteDepreciacion(fechaCorte string, elementos interface{}) (err error) {

	query_ := map[string]string{
		"FormatoTipoMovimientoId__CodigoAbreviacion": "CRR",
		"EstadoMovimientoId__Nombre":                 "Cierre Aprobado",
		"FechaCorte__gte":                            fechaCorte,
	}

	cierres, _, err := GetAllMovimiento(query_, nil, nil, nil, 0, 1)
	if err != nil || len(cierres) > 0 {
		return err
	}

	o := orm.NewOrm()
	// Los elementos se determinan de la siguiente manera
	// + Elementos sin novedad y vida útil > 0
	// + Novedades con vida útil > 0
	// - Elementos solicitados para baja antes de la fecha de corte

	_, err = o.Raw(getScriptCorte(), fechaCorte).QueryRows(elementos)

	return
}

func getScriptCorte() string {

	const script = `
	WITH fecha_corte_ AS (
		SELECT (TO_DATE(?, 'YYYY-MM-DD'))::date fecha_corte_
	), bajas AS (
		SELECT
			em.id
		FROM
			ESQUEMA.movimiento m,
			ESQUEMA.elementos_movimiento em,
			ESQUEMA.formato_tipo_movimiento fm,
			jsonb_array_elements(m.detalle -> 'Elementos') AS elem,
			fecha_corte_
		WHERE
			fm.codigo_abreviacion LIKE 'BJ%'
			AND m.formato_tipo_movimiento_id = fm.id
			AND em.id = CAST(elem as INTEGER)
	), con_novedad AS (
		SELECT DISTINCT ON (1)
			ne.elemento_movimiento_id,
			m.fecha_corte,
			ne.valor_residual,
			ne.vida_util,
			em.elemento_acta_id,
			ne.valor_libros valor_presente
		FROM
			ESQUEMA.novedad_elemento ne,
			ESQUEMA.elementos_movimiento em,
			ESQUEMA.movimiento m,
			fecha_corte_
		WHERE
				m.fecha_corte < fecha_corte_ + 1
			AND	ne.elemento_movimiento_id = em.id
			AND ne.movimiento_id = m.id
			AND em.id NOT IN (
				SELECT
					em.id
				FROM
					bajas
				WHERE
					bajas.id = em.id
			)
			AND em.id IN (
				SELECT
					em.id
				FROM
					ESQUEMA.movimiento m,
					ESQUEMA.elementos_movimiento em,
					ESQUEMA.formato_tipo_movimiento fm
				WHERE
					fm.codigo_abreviacion = 'SAL'
					AND m.formato_tipo_movimiento_id = fm.id
					AND em.movimiento_id = m.id
			)
		ORDER BY 1 DESC, m.fecha_corte DESC
	), sin_novedad AS (
		SELECT
			em.id elemento_movimiento_id,
			m.fecha_corte - 1 fecha_corte,
			em.valor_residual,
			em.vida_util,
			em.elemento_acta_id,
			em.valor_total valor_presente
		FROM
			ESQUEMA.elementos_movimiento em,
			ESQUEMA.movimiento m,
			ESQUEMA.estado_movimiento sm,
			ESQUEMA.formato_tipo_movimiento fm,
			fecha_corte_
		WHERE
				fm.codigo_abreviacion = 'SAL'
			AND sm.nombre = 'Salida Aprobada'
			AND m.formato_tipo_movimiento_id  = fm.id
			AND m.fecha_corte < fecha_corte_ + 1
			AND m.estado_movimiento_id = sm.id
			AND em.movimiento_id = m.id
			AND em.valor_total > 0
			AND em.vida_util > 0
			AND em.id NOT IN (
				SELECT elemento_movimiento_id
				FROM con_novedad
			)
			AND em.id NOT IN (
				SELECT
					em.id
				FROM
					bajas
				WHERE
					bajas.id = em.id
			)
	), referencia AS (
		SELECT *
		FROM sin_novedad
		UNION
		SELECT *
		FROM con_novedad
	), referencia_completo AS (
		SELECT
			referencia.*,
			CASE
				WHEN
					delta_dias > 1 AND (
						EXTRACT(day FROM (DATE_TRUNC('month', referencia.fecha_corte) + interval '1 month - 1 day')) = 31 AND (
							delta_meses > 0 OR delta_year > 0
						) OR (
							EXTRACT(day FROM fecha_corte_) = 31 AND delta_meses = 0 AND delta_year = 0
						)
					)
				THEN
					delta_dias - 1
				ELSE
					delta_dias
			END delta_dias,
			(delta_meses * 30) + (delta_year * 360) AS delta_dias_
		FROM
			referencia,
			fecha_corte_,
			EXTRACT(year FROM AGE(fecha_corte_, referencia.fecha_corte)) delta_year,
			EXTRACT(month FROM AGE(fecha_corte_, referencia.fecha_corte)) delta_meses,
			EXTRACT(day FROM AGE(fecha_corte_, referencia.fecha_corte)) delta_dias
	), delta_valor AS (
		SELECT
			elemento_movimiento_id,
			elemento_acta_id,
			CASE
				WHEN
					360 * vida_util - delta_dias - delta_dias_ > 1
				THEN
					(valor_presente - valor_residual) * (delta_dias + delta_dias_) / (vida_util * 360)
				ELSE valor_presente - valor_residual
			END delta_valor
		FROM referencia_completo
		WHERE
			vida_util > 0
			AND valor_presente > valor_residual
	)

	SELECT * from delta_valor;
	`

	return replaceSquema(script)
}

// SubmitCierre Actualiza el cierre y genera las novedades correspondientes
func SubmitCierre(cierre *Movimiento) (err error) {

	o := orm.NewOrm()
	err = o.Begin()

	defer func() {
		if r := recover(); r != nil {
			o.Rollback()
			logs.Error(r)
		} else {
			o.Commit()
		}
	}()

	err = o.QueryTable(new(Movimiento)).RelatedSel().Filter("Id", cierre.Id).One(cierre)
	if err != nil || cierre.EstadoMovimientoId.Nombre != "Cierre En Curso" || cierre.FechaCorte == nil {
		return
	}

	err = o.QueryTable(new(EstadoMovimiento)).Filter("Nombre", "Cierre Aprobado").One(cierre.EstadoMovimientoId)
	if err != nil {
		return
	}

	_, err = o.Raw(getScriptAprobacion(), cierre.Id).Exec()
	if err != nil {
		return
	}

	cierre.Detalle = "{}"
	_, err = o.Update(cierre, "Detalle", "EstadoMovimientoId")

	return
}

func getScriptAprobacion() string {

	const script = `
	WITH cierre AS (
		SELECT
			id,
			fecha_corte AS fecha_corte_
		FROM ESQUEMA.movimiento
		WHERE id = ?
	), bajas AS (
		SELECT
			em.id
		FROM
			ESQUEMA.movimiento m,
			ESQUEMA.elementos_movimiento em,
			ESQUEMA.formato_tipo_movimiento fm,
			jsonb_array_elements(m.detalle -> 'Elementos') AS elem,
			cierre
		WHERE
			fm.codigo_abreviacion LIKE 'BJ%'
			AND m.formato_tipo_movimiento_id = fm.id
			AND m.fecha_creacion < cierre.fecha_corte_ + 1
			AND em.id = CAST(elem as INTEGER)
	), con_novedad AS (
		SELECT DISTINCT ON (1)
			ne.elemento_movimiento_id,
			m.fecha_corte,
			ne.valor_residual,
			ne.vida_util,
			ne.valor_libros
		FROM
			ESQUEMA.novedad_elemento ne,
			ESQUEMA.elementos_movimiento em,
			ESQUEMA.movimiento m,
			cierre
		WHERE
				m.fecha_corte < cierre.fecha_corte_ + 1
			AND	ne.elemento_movimiento_id = em.id
			AND ne.movimiento_id = m.id
			AND em.id IN (
				SELECT
					em.id
				FROM
					ESQUEMA.elementos_movimiento em,
					ESQUEMA.movimiento m,
					ESQUEMA.formato_tipo_movimiento fm
				WHERE
					fm.codigo_abreviacion IN ('SAL', 'INM_REG')
					AND m.formato_tipo_movimiento_id = fm.id
					AND em.movimiento_id = m.id
			)
			AND em.id NOT IN (
				SELECT em.id
				FROM bajas
				WHERE bajas.id = em.id
			)
		ORDER BY 1 DESC, m.fecha_corte DESC
	), sin_novedad AS (
		SELECT
			em.id elemento_movimiento_id,
			m.fecha_corte - 1 fecha_corte,
			em.valor_residual,
			em.vida_util,
			em.valor_total valor_libros
		FROM
			ESQUEMA.elementos_movimiento em,
			ESQUEMA.movimiento m,
			ESQUEMA.estado_movimiento sm,
			ESQUEMA.formato_tipo_movimiento fm,
			cierre
		WHERE
				fm.codigo_abreviacion IN ('SAL', 'INM_REG')
			AND sm.nombre IN ('Salida Aprobada', 'Bienes inmuebles registrados')
			AND m.formato_tipo_movimiento_id  = fm.id
			AND m.estado_movimiento_id = sm.id
			AND m.fecha_corte < cierre.fecha_corte_ + 1
			AND em.movimiento_id = m.id
			AND em.valor_total > 0
			AND em.vida_util > 0
			AND em.id NOT IN (
				SELECT elemento_movimiento_id
				FROM con_novedad
			)
			AND em.id NOT IN (
				SELECT em.id
				FROM bajas
				WHERE bajas.id = em.id
			)
	), referencia AS (
		SELECT *
		FROM con_novedad
		WHERE
				vida_util > 0
			AND valor_libros > 0
		UNION
		SELECT *
		FROM sin_novedad
	), referencia_completo AS (
		SELECT
			referencia.*,
			CASE
				WHEN
					delta_dias > 1 AND (
						EXTRACT(day FROM (DATE_TRUNC('month', referencia.fecha_corte) + interval '1 month - 1 day')) = 31 AND (
							delta_meses > 0 OR delta_year > 0
						) OR (
							EXTRACT(day FROM cierre.fecha_corte_) = 31 AND delta_meses = 0 AND delta_year = 0
						)
					)
				THEN
					delta_dias - 1
				ELSE
					delta_dias
			END delta_dias,
			(delta_meses * 30) + (delta_year * 360) AS delta_dias_
		FROM
			cierre,
			referencia,
			EXTRACT(year FROM AGE(cierre.fecha_corte_, referencia.fecha_corte)) delta_year,
			EXTRACT(month FROM AGE(cierre.fecha_corte_, referencia.fecha_corte)) delta_meses,
			EXTRACT(day FROM AGE(cierre.fecha_corte_, referencia.fecha_corte)) delta_dias
	), delta AS (
		SELECT
			elemento_movimiento_id,
			valor_residual,
			CASE
				WHEN
					360 * vida_util - delta_dias - delta_dias_ > 1
				THEN
					vida_util - (delta_dias + delta_dias_) / 360
				ELSE 0
			END vida_util,
			CASE
				WHEN
					360 * vida_util - delta_dias - delta_dias_ > 1
				THEN
					valor_libros - (valor_libros - valor_residual) * (delta_dias + delta_dias_) / (vida_util * 360)
				ELSE valor_residual
			END valor_libros
		FROM referencia_completo
	)

	INSERT INTO ESQUEMA.novedad_elemento (
			vida_util,
			valor_libros,
			valor_residual,
			elemento_movimiento_id,
			movimiento_id,
			activo,
			fecha_modificacion,
			fecha_creacion)
	SELECT
		delta.vida_util,
		delta.valor_libros,
		delta.valor_residual,
		delta.elemento_movimiento_id,
		cierre.id,
		true,
		now(),
		now()
	FROM
		delta,
		cierre;
	`
	return replaceSquema(script)
}

func replaceSquema(script string) string {
	Esquema = beego.AppConfig.String("PGschemas")
	return strings.ReplaceAll(script, "ESQUEMA", Esquema)
}
