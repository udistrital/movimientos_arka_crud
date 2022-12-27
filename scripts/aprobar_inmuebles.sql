WITH cierre AS (
    SELECT (TO_DATE(?, 'YYYY-MM-DD') + INTERVAL '1 day')::date fecha_corte
), con_novedad AS (
    SELECT DISTINCT ON (1)
        ne.elemento_movimiento_id,
        fecha,
        ne.valor_residual,
        ne.vida_util,
        ne.valor_libros,
        CASE
            WHEN
                delta_dias > 1 AND (
                    EXTRACT(day FROM (DATE_TRUNC('month', fecha) + interval '1 month - 1 day')) = 31 AND (
                        EXTRACT(month FROM cierre.fecha_corte - 1) != EXTRACT(month FROM fecha) OR
                        EXTRACT(year FROM cierre.fecha_corte - 1) != EXTRACT(year FROM fecha)
                    ) OR (
                        EXTRACT(day FROM cierre.fecha_corte - 1) = 31 AND delta_meses = 0 AND delta_year = 0
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
        cierre,
        EXTRACT(year FROM AGE(cierre.fecha_corte, fecha + interval '1 day')) delta_year,
        EXTRACT(month FROM AGE(cierre.fecha_corte, fecha + interval '1 day')) delta_meses,
        EXTRACT(day FROM AGE(cierre.fecha_corte, fecha + interval '1 day')) delta_dias
    WHERE
			fecha < cierre.fecha_corte
        AND	ne.elemento_movimiento_id = em.id
        AND ne.movimiento_id = m.id
		AND em.id IN (
			SELECT
				em.id
			FROM 
		        movimientos_arka.elementos_movimiento em,
		        movimientos_arka.movimiento m,
		        movimientos_arka.formato_tipo_movimiento fm
			WHERE
					fm.codigo_abreviacion = 'INM_REG'
				AND m.formato_tipo_movimiento_id = fm.id
				AND em.movimiento_id = m.id
		)
    ORDER BY 1 DESC, fecha DESC
), sin_novedad AS (
    SELECT
        em.id elemento_movimiento_id,
        em.valor_residual,
        em.vida_util,
        em.valor_total valor_libros,
        CASE
            WHEN
                delta_dias > 1 AND (
                    (EXTRACT(day FROM (DATE_TRUNC('month', fecha) + interval '1 month - 1 day')) = 31 AND (delta_meses > 0 OR delta_year > 0)) OR
                    (EXTRACT(day FROM cierre.fecha_corte - 1) = 31 AND delta_meses = 0 AND delta_year = 0)
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
        movimientos_arka.formato_tipo_movimiento fm,
        to_date(m.detalle->>'FechaCorte', 'YYYY-MM-DD') AS fecha,
        cierre,
        EXTRACT(day FROM AGE(cierre.fecha_corte, fecha)) delta_dias,
        EXTRACT(month FROM AGE(cierre.fecha_corte, fecha)) delta_meses,
        EXTRACT(year FROM AGE(cierre.fecha_corte, fecha)) delta_year
    WHERE
            fm.codigo_abreviacion = 'INM_REG'
        AND m.formato_tipo_movimiento_id  = fm.id
        AND fecha < cierre.fecha_corte
        AND em.movimiento_id = m.id
        AND em.valor_total > 0
        AND em.vida_util > 0
        AND em.id NOT IN (
            SELECT elemento_movimiento_id
            FROM con_novedad
        )
), referencia AS (
	SELECT
		elemento_movimiento_id,
		valor_residual,
		vida_util,
		valor_libros,
		delta_dias,
		delta_dias_
	FROM
		con_novedad
	WHERE
			vida_util > 0
		AND valor_libros > 0
	UNION
	SELECT *
	FROM
		sin_novedad
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
    delta.elemento_movimiento_id,
    ?,
    true,
    now(),
    now()
FROM
    delta;
