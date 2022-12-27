package models

import (
	"encoding/json"
	"io/ioutil"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type FormatoCierre struct {
	ConsecutivoId int
	Consecutivo   string
	FechaCorte    string
	RazonRechazo  string
}

type DepreciacionElemento struct {
	DeltaValor           float64
	ElementoMovimientoId int
	ElementoActaId       int
}

type TransaccionCierre struct {
	MovimientoId         int
	ElementoMovimientoId []int
}

// GetCorteDepreciacion Retorna los valores y detalles necesarios para generar
// la transacción contable correspondiente a la depreciación dada una fecha de corte
func GetCorteDepreciacion(fechaCorte string, elementos interface{}) (err error) {

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
			fm.codigo_abreviacion = 'CRR'
			AND sm.nombre = 'Cierre Aprobado'
			AND m.formato_tipo_movimiento_id = fm.id
			AND m.estado_movimiento_id = sm.id
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
				CASE
					WHEN
						delta_dias > 1 AND (
							EXTRACT(day FROM (DATE_TRUNC('month', m.fecha_modificacion) + interval '1 month - 1 day')) = 31 AND (
								EXTRACT(month FROM fecha_corte - 1) != EXTRACT(month FROM m.fecha_modificacion) OR
								EXTRACT(year FROM fecha_corte - 1) != EXTRACT(year FROM m.fecha_modificacion)
							) OR (
								EXTRACT(day FROM fecha_corte - 1) = 31 AND delta_meses = 0 AND delta_year = 0
							)
						)
					THEN
						delta_dias - 1
					ELSE
						delta_dias
				END delta_dias,
				(delta_meses * 30) + (delta_year * 360) AS delta_dias_
			FROM
				movimientos_arka.elementos_movimiento em,
				movimientos_arka.movimiento m,
				movimientos_arka.estado_movimiento sm,
				movimientos_arka.formato_tipo_movimiento fm,
				fecha_corte,
				EXTRACT(year FROM AGE(fecha_corte, m.fecha_modificacion - interval '1 day')) delta_year,
				EXTRACT(month FROM AGE(fecha_corte, m.fecha_modificacion - interval '1 day')) delta_meses,
				EXTRACT(day FROM AGE(fecha_corte, m.fecha_modificacion - interval '1 day')) delta_dias
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
				ne.valor_residual,
				ne.vida_util,
				em.elemento_acta_id,
				ne.valor_libros valor_presente,
				CASE
					WHEN
						delta_dias > 1 AND (
							EXTRACT(day FROM (DATE_TRUNC('month', fecha) + interval '1 month - 1 day')) = 31 AND (
								EXTRACT(month FROM fecha_corte - 1) != EXTRACT(month FROM fecha) OR
								EXTRACT(year FROM fecha_corte - 1) != EXTRACT(year FROM fecha)
							) OR (
								EXTRACT(day FROM fecha_corte - 1) = 31 AND delta_meses = 0 AND delta_year = 0
							)
						)
					THEN
						delta_dias - 1
					ELSE
						delta_dias
				END delta_dias,
				(delta_meses * 30) + (delta_year * 360) AS delta_dias_
			FROM
				movimientos_arka.novedad_elemento ne,
				movimientos_arka.elementos_movimiento em,
				movimientos_arka.movimiento m,
				to_date(m.detalle->>'FechaCorte', 'YYYY-MM-DD') AS fecha,
				fecha_corte,
				EXTRACT(year FROM AGE(fecha_corte, fecha + interval '1 day')) delta_year,
				EXTRACT(month FROM AGE(fecha_corte, fecha + interval '1 day')) delta_meses,
				EXTRACT(day FROM AGE(fecha_corte, fecha + interval '1 day')) delta_dias
			WHERE
				fecha < fecha_corte
				AND	ne.elemento_movimiento_id = em.id
				AND ne.movimiento_id = m.id
				AND em.id NOT IN (
					SELECT
						bajas.id
					FROM 
						bajas
					WHERE
						bajas.id = em.id
				)
				AND em.id IN (
					SELECT
						em.id
					FROM
						movimientos_arka.movimiento m,
						movimientos_arka.elementos_movimiento em,
						movimientos_arka.formato_tipo_movimiento fm
					WHERE
						fm.codigo_abreviacion = 'SAL'
						AND m.formato_tipo_movimiento_id = fm.id
						AND em.movimiento_id = m.id
				)
			ORDER BY 1 DESC, fecha DESC
		), referencia AS (
			SELECT *
			FROM
				sin_novedad ndp
			UNION
			SELECT *
			FROM
				con_novedad dp
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
			FROM
				referencia
			WHERE
				vida_util > 0
				AND valor_presente > valor_residual
		)
		
		SELECT * from delta_valor;`

	if _, err = o.Raw(query, fechaCorte).QueryRows(elementos); err != nil {
		return err
	}

	return
}

// SubmitCierre Actualiza el cierre y genera las novedades correspondientes
func SubmitCierre(m *TransaccionCierre, cierre *Movimiento) (err error) {

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

	if m.MovimientoId <= 0 || len(m.ElementoMovimientoId) == 0 {
		return
	}

	if _, err = o.QueryTable(new(Movimiento)).RelatedSel().Filter("Id", m.MovimientoId).All(cierre); err != nil {
		return
	} else if cierre.EstadoMovimientoId.Nombre != "Cierre En Curso" {
		return
	}

	if _, err = o.QueryTable(new(EstadoMovimiento)).RelatedSel().Filter("Nombre", "Cierre Aprobado").All(cierre.EstadoMovimientoId); err != nil {
		return
	}

	var detalle FormatoCierre
	err = json.Unmarshal([]byte(cierre.Detalle), &detalle)
	if err != nil {
		return err
	}

	if detalle.FechaCorte == "" {
		return
	}

	query := `
	WITH fecha_corte AS (
		SELECT (TO_DATE(?, 'YYYY-MM-DD') + INTERVAL '1 day')::date fecha_corte
	), elemento AS (
		SELECT CAST(? as INTEGER) id
	), referencia AS (
		(SELECT
			ne.valor_residual,
			ne.vida_util,
			ne.valor_libros valor_presente,
			CASE
				WHEN
					delta_dias > 1 AND	(
						EXTRACT(day FROM (DATE_TRUNC('month', fecha) + interval '1 month - 1 day')) = 31 AND (
							EXTRACT(month FROM fecha_corte) != EXTRACT(month FROM fecha) OR
							EXTRACT(year FROM fecha_corte) != EXTRACT(year FROM fecha)
						)
					) OR (
						EXTRACT(day FROM fecha_corte) = 31 AND delta_meses = 0 AND delta_year = 0
					)
				THEN
					delta_dias - 1
				ELSE
					delta_dias
			END delta_dias,
			(delta_meses * 30) + (delta_year * 360) AS delta_dias_
		FROM
			movimientos_arka.novedad_elemento ne,
			movimientos_arka.elementos_movimiento em,
			movimientos_arka.movimiento m,
			to_date(m.detalle->>'FechaCorte', 'YYYY-MM-DD') AS fecha,
			fecha_corte,
			elemento,
			EXTRACT(year FROM AGE(fecha_corte, fecha + interval '1 day')) delta_year,
			EXTRACT(month FROM AGE(fecha_corte, fecha + interval '1 day')) delta_meses,
			EXTRACT(day FROM AGE(fecha_corte, fecha + interval '1 day')) delta_dias
		WHERE
			fecha < fecha_corte
			AND em.id = elemento.id
			AND	ne.elemento_movimiento_id = em.id
			AND ne.movimiento_id = m.id
		ORDER BY fecha DESC
		LIMIT 1)
		UNION
		SELECT
			em.valor_residual,
			em.vida_util,
			em.valor_total valor_presente,
			CASE
				WHEN
					delta_dias > 1 AND	(
						EXTRACT(day FROM (DATE_TRUNC('month', m.fecha_modificacion) + interval '1 month - 1 day')) = 31 AND (
							EXTRACT(month FROM fecha_corte) != EXTRACT(month FROM m.fecha_modificacion) OR
							EXTRACT(year FROM fecha_corte) != EXTRACT(year FROM m.fecha_modificacion)
						)
					) OR (
						EXTRACT(day FROM fecha_corte) = 31 AND delta_meses = 0 AND delta_year = 0
					)
				THEN
					delta_dias - 1
				ELSE
					delta_dias
			END delta_dias,
			(delta_meses * 30) + (delta_year * 360) AS delta_dias_
		FROM
			movimientos_arka.elementos_movimiento em,
			movimientos_arka.movimiento m,
			movimientos_arka.estado_movimiento sm,
			movimientos_arka.formato_tipo_movimiento fm,
			fecha_corte,
			elemento,
			EXTRACT(year FROM AGE(fecha_corte, m.fecha_modificacion - interval '1 day')) delta_year,
			EXTRACT(month FROM AGE(fecha_corte, m.fecha_modificacion - interval '1 day')) delta_meses,
			EXTRACT(day FROM AGE(fecha_corte, m.fecha_modificacion - interval '1 day')) delta_dias
		WHERE
			fm.codigo_abreviacion = 'SAL'
			AND sm.nombre = 'Salida Aprobada'
			AND m.formato_tipo_movimiento_id  = fm.id
			AND m.estado_movimiento_id = sm.id
			AND em.movimiento_id = m.id
			AND em.id = elemento.id
			AND em.id NOT IN (
				SELECT 
					em.id
				FROM
					movimientos_arka.novedad_elemento ne
				WHERE
					ne.elemento_movimiento_id = em.id
				)
	), delta AS (
		SELECT
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
					valor_presente - (valor_presente - valor_residual) * (delta_dias + delta_dias_) / (vida_util * 360)
				ELSE valor_residual
			END valor_libros
		FROM
			referencia
	)

	INSERT INTO movimientos_arka.novedad_elemento (
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
		elemento.id,
		?,
		true,
		now(),
		now()
	FROM
		delta,
		elemento;`

	p, err := o.Raw(query).Prepare()
	if err != nil {
		return err
	}

	for _, el := range m.ElementoMovimientoId {
		_, err = p.Exec(detalle.FechaCorte, el, m.MovimientoId)
		if err != nil {
			return err
		}
	}

	if err = p.Close(); err != nil {
		return
	}

	var script []byte
	script, err = ioutil.ReadFile("models/aprobar_inmuebles.sql")
	if err != nil {
		return
	}

	_, err = o.Raw(string(script), detalle.FechaCorte, m.MovimientoId).Exec()
	if err != nil {
		return
	}

	detalle.RazonRechazo = ""
	if dt, err := json.Marshal(detalle); err != nil {
		return err
	} else {
		cierre.Detalle = string(dt)
	}

	err = UpdateMovimientoById(cierre)

	return
}
