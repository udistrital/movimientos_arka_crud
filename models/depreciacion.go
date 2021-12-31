package models

import (
	"time"

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
	// + Novedad con vida útil > 0

	var (
		query     string
		elementos []*detalle
		novedad   []*detalle
	)
	o := orm.NewOrm()
	query =
		`SELECT
		em.id as elemento_movimiento_id,
		em.valor_residual,
		em.vida_util,
		em.elemento_acta_id,
		em.valor_total as valor_presente,
		m.fecha_creacion as fecha_ref
	FROM
		movimientos_arka.elementos_movimiento em,
		movimientos_arka.movimiento m,
		movimientos_arka.formato_tipo_movimiento fm
	WHERE	fm.nombre = 'Salida'
		AND m.formato_tipo_movimiento_id  = fm.id
		AND m.fecha_creacion < ?
		AND em.movimiento_id = m.id
		AND em.vida_util > 0
		AND NOT EXISTS (
			SELECT
			FROM
				movimientos_arka.novedad_elemento ne
			WHERE
				ne.elemento_movimiento_id = em.id);`

	if _, err = o.Raw(query, fechaCorte).QueryRows(&elementos); err != nil {
		return nil, err
	}

	query =
		`SELECT
		DISTINCT ON (ne.elemento_movimiento_id) ne.elemento_movimiento_id,
		ne.id as novedad_elemento_id,
		ne.vida_util,
		ne.valor_residual,
		ne.valor_libros as valor_presente,
		ne.fecha_creacion as fecha_ref,
		em.elemento_acta_id
	FROM
		movimientos_arka.novedad_elemento ne,
		movimientos_arka.elementos_movimiento em,
		movimientos_arka.movimiento m
	WHERE
		m.fecha_creacion < ?
		AND em.movimiento_id = m.id
		AND ne.elemento_movimiento_id = em.id
		AND ne.activo = true
		AND ne.vida_util > 0
	ORDER BY
		ne.elemento_movimiento_id,
		ne.id DESC;`

	if _, err = o.Raw(query, fechaCorte).QueryRows(&novedad); err != nil {
		return nil, err
	}
	elementos = append(elementos, novedad...)

	return elementos, nil
}
